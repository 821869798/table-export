package visitor

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/convert/wrap"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/field_type"
	"github.com/821869798/table-export/serialization"
	"github.com/gookit/slog"
)

type BinaryVisitor struct {
	byteBuff *serialization.ByteBuf
}

func NewBinary(buff *serialization.ByteBuf) apiconvert.IDataVisitor {
	b := &BinaryVisitor{
		byteBuff: buff,
	}
	return b
}

func (b *BinaryVisitor) AcceptTable(dataModel *model.TableModel) {
	optimize := dataModel.Optimize
	if optimize != nil && len(optimize.OptimizeFields) > 0 {
		b.byteBuff.WriteSize(len(optimize.OptimizeFields))
		for _, tableOptimizeField := range optimize.OptimizeFields {
			b.byteBuff.WriteSize(len(tableOptimizeField.OptimizeDataInTableRow))
			for _, originRowIndex := range tableOptimizeField.OptimizeDataInTableRow {
				recordMap, err := dataModel.MemTable.GetRecordByRow(originRowIndex)
				//err := wrap.RunDataVisitorString(b, tableOptimizeField.Field.Type, origin)
				if err != nil {
					slog.Fatalf("export binary target file[%v] optimize get record error:%v", dataModel.Meta.Target, err.Error())
				}
				value := recordMap[tableOptimizeField.Field.Target]
				err = wrap.RunDataVisitorValue(b, tableOptimizeField.Field.Type, value)
				if err != nil {
					slog.Fatalf("export binary target file[%v] optimize visitor error:%v", dataModel.Meta.Target, err.Error())
				}
			}
		}
	} else {
		b.byteBuff.WriteSize(0)
	}

	b.byteBuff.WriteSize(len(dataModel.RawData))
	rowDataOffset := config.GlobalConfig.Table.DataStart + 1
	for rowIndex, recordIndex := range dataModel.MemTable.RowIndexList() {
		for _, tf := range dataModel.Meta.Fields {
			if optimize != nil {
				tableOptimizeField, _ := optimize.GetOptimizeField(tf)
				if tableOptimizeField != nil {
					var dIndex = tableOptimizeField.DataUseIndex[rowIndex]
					// 索引+1,从1开始给之后有空类型的数据0考虑
					b.byteBuff.WriteInt(int32(dIndex) + 1)
					continue
				}
			}

			recordMap := dataModel.MemTable.GetRecordRecordMap(recordIndex)
			value, ok := recordMap[tf.Target]
			if !ok {
				slog.Fatalf("export binary target file[%v] RowCount[%v] filedName[%v] not found", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source)
			}
			err := wrap.RunDataVisitorValue(b, tf.Type, value)
			if err != nil {
				slog.Fatalf("export binary target file[%v] RowCount[%v] filedName[%v] error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err.Error())
			}
		}
	}

}

func (b *BinaryVisitor) AcceptInt(r int32) {
	b.byteBuff.WriteInt(r)
}

func (b *BinaryVisitor) AcceptUInt(r uint32) {
	b.byteBuff.WriteUint(r)
}

func (b *BinaryVisitor) AcceptLong(r int64) {
	b.byteBuff.WriteLong(r)
}

func (b *BinaryVisitor) AcceptULong(r uint64) {
	b.byteBuff.WriteUlong(r)
}

func (b *BinaryVisitor) AcceptBool(r bool) {
	b.byteBuff.WriteBool(r)
}

func (b *BinaryVisitor) AcceptFloat(r float32) {
	b.byteBuff.WriteFloat(r)
}

func (b *BinaryVisitor) AcceptDouble(r float64) {
	b.byteBuff.WriteDouble(r)
}

func (b *BinaryVisitor) AcceptString(r string) {
	b.byteBuff.WriteString(r)
}

func (b *BinaryVisitor) AcceptByte(r byte) {
	b.byteBuff.WriteByte(r)
}

func (b *BinaryVisitor) AcceptArray(r []interface{}, ValueType *field_type.TableFieldType) {
	b.byteBuff.WriteSize(len(r))
	for _, origin := range r {
		err := wrap.RunDataVisitorValue(b, ValueType, origin)
		if err != nil {
			slog.Fatalf("export binary AcceptArray failed: %v", err)
		}
	}
}

func (b *BinaryVisitor) AcceptStringArray(r []string, ValueType *field_type.TableFieldType) {
	b.byteBuff.WriteSize(len(r))
	for _, origin := range r {
		err := wrap.RunDataVisitorString(b, ValueType, origin)
		if err != nil {
			slog.Fatalf("export binary AcceptArray failed: %v", err)
		}
	}
}

func (b *BinaryVisitor) AcceptStringMap(r map[string]string, KeyType *field_type.TableFieldType, ValueType *field_type.TableFieldType) {
	b.byteBuff.WriteSize(len(r))
	for _, key := range GetMapSortedKeys(r) {
		value := r[key]
		err := wrap.RunDataVisitorString(b, KeyType, key)
		if err != nil {
			slog.Fatalf("export binary AcceptStringMap failed: %v", err)
		}
		err = wrap.RunDataVisitorString(b, ValueType, value)
		if err != nil {
			slog.Fatalf("export binary AcceptStringMap failed: %v", err)
		}
	}
}

func (b *BinaryVisitor) AcceptMap(r map[string]interface{}, KeyType *field_type.TableFieldType, ValueType *field_type.TableFieldType) {
	b.byteBuff.WriteSize(len(r))
	for _, key := range GetMapSortedKeys(r) {
		value := r[key]
		err := wrap.RunDataVisitorString(b, KeyType, key)
		if err != nil {
			slog.Fatalf("export binary AcceptMap failed: %v", err)
		}
		err = wrap.RunDataVisitorValue(b, ValueType, value)
		if err != nil {
			slog.Fatalf("export binary AcceptMap failed: %v", err)
		}
	}
}
func (b *BinaryVisitor) AcceptCommonMap(r map[interface{}]interface{}, KeyType *field_type.TableFieldType, ValueType *field_type.TableFieldType) {
	b.byteBuff.WriteSize(len(r))
	for _, key := range GetMapSortedKeysInterface(r, KeyType) {
		value := r[key]
		err := wrap.RunDataVisitorValue(b, KeyType, key)
		if err != nil {
			slog.Fatalf("export binary AcceptCommonMap failed: %v", err)
		}
		err = wrap.RunDataVisitorValue(b, ValueType, value)
		if err != nil {
			slog.Fatalf("export binary AcceptCommonMap failed: %v", err)
		}
	}
}

func (b *BinaryVisitor) AcceptClass(r map[string]interface{}, class *field_type.TableFieldClass) {
	for _, field := range class.AllFields() {
		value, ok := r[field.Name]
		if !ok {
			// 写入默认值
			err := wrap.RunDataVisitorString(b, field.Type, "")
			if err != nil {
				slog.Fatalf("export binary AcceptClass failed: %v", err)
			}
		}
		err := wrap.RunDataVisitorValue(b, field.Type, value)
		if err != nil {
			slog.Fatalf("export binary AcceptClass failed: %v", err)
		}
	}
}

func (b *BinaryVisitor) AcceptClassString(r map[string]string, class *field_type.TableFieldClass) {
	for _, field := range class.AllFields() {
		value, ok := r[field.Name]
		if !ok {
			// 使用默认值解析
			value = ""
		}
		err := wrap.RunDataVisitorString(b, field.Type, value)
		if err != nil {
			slog.Fatalf("export binary AcceptClass failed: %v", err)
		}
	}
}
