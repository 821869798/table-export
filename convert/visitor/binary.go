package visitor

import (
	log "github.com/sirupsen/logrus"
	"table-export/config"
	"table-export/convert/adapter"
	"table-export/convert/api"
	"table-export/convert/wrap"
	"table-export/data/model"
	"table-export/serialization"
)

type BinaryVisitor struct {
	byteBuff *serialization.ByteBuf
}

func NewBinary(buff *serialization.ByteBuf) api.IDataVisitor {
	b := &BinaryVisitor{
		byteBuff: buff,
	}
	return b
}

func (b BinaryVisitor) AcceptTable(dataModel *model.TableModel) {
	optimize := dataModel.Optimize
	if optimize != nil && len(optimize.OptimizeFields) > 0 {
		b.byteBuff.WriteSize(len(optimize.OptimizeFields))
		for _, tableOptimizeField := range optimize.OptimizeFields {
			b.byteBuff.WriteSize(len(tableOptimizeField.OriginDatas))
			for _, origin := range tableOptimizeField.OriginDatas {
				err := wrap.GetDataVisitorValue(b, tableOptimizeField.Field.Type, origin)
				if err != nil {
					log.Fatalf("export binary target file[%v] optimize error:%v", dataModel.Meta.Target, err.Error())
				}
			}
		}
	} else {
		b.byteBuff.WriteSize(0)
	}

	b.byteBuff.WriteSize(len(dataModel.RawData))
	rowDataOffset := config.GlobalConfig.Table.DataStart + 1
	for rowIndex, rowData := range dataModel.RawData {
		for _, tf := range dataModel.Meta.Fields {
			rawIndex := dataModel.NameIndexMapping[tf.Target]
			var rawStr string
			if rawIndex < len(rowData) {
				rawStr = rowData[rawIndex]
			}
			if optimize != nil {
				tableOptimizeField, _ := optimize.GetOptimizeField(tf)
				if tableOptimizeField != nil {
					var dIndex = tableOptimizeField.DataIndexs[rowIndex]
					// 索引+1,从1开始给之后有空类型的数据0考虑
					b.byteBuff.WriteInt(int32(dIndex) + 1)
					continue
				}
			}
			err := wrap.GetDataVisitorValue(b, tf.Type, rawStr)
			if err != nil {
				log.Fatalf("export binary target file[%v] RowCount[%v] filedName[%v] error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err.Error())
			}
		}
	}
}

func (b BinaryVisitor) AcceptInt(r int32) {
	b.byteBuff.WriteInt(r)
}

func (b BinaryVisitor) AcceptUInt(r uint32) {
	b.byteBuff.WriteUint(r)
}

func (b BinaryVisitor) AcceptLong(r int64) {
	b.byteBuff.WriteLong(r)
}

func (b BinaryVisitor) AcceptULong(r uint64) {
	b.byteBuff.WriteUlong(r)
}

func (b BinaryVisitor) AcceptBool(r bool) {
	b.byteBuff.WriteBool(r)
}

func (b BinaryVisitor) AcceptFloat(r float32) {
	b.byteBuff.WriteFloat(r)
}

func (b BinaryVisitor) AcceptDouble(r float64) {
	b.byteBuff.WriteDouble(r)
}

func (b BinaryVisitor) AcceptString(r string) {
	b.byteBuff.WriteString(r)
}

func (b BinaryVisitor) AcceptByte(r byte) {
	b.byteBuff.WriteByte(r)
}

func (b BinaryVisitor) AcceptArray(r *adapter.Array) {
	b.byteBuff.WriteSize(len(r.Datas))
	for _, origin := range r.Datas {
		err := wrap.GetDataVisitorValue(b, r.ValueType, origin)
		if err != nil {
			log.Fatalf("export binary AcceptArray failed: %s", err.Error())
		}
	}
}

func (b BinaryVisitor) AcceptMap(r *adapter.Map) {
	b.byteBuff.WriteSize(len(r.Datas))
	for key, value := range r.Datas {
		err := wrap.GetDataVisitorValue(b, r.KeyType, key)
		if err != nil {
			log.Fatalf("export binary AcceptMap failed: %s", err.Error())
		}
		err = wrap.GetDataVisitorValue(b, r.ValueType, value)
		if err != nil {
			log.Fatalf("export binary AcceptMap failed: %s", err.Error())
		}
	}
}
