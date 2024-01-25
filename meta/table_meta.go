package meta

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/field_type"
	"github.com/gookit/slog"
)

type TableMeta struct {
	Target        string
	Mode          string
	SourceType    string
	Sources       []*TableSource
	Fields        []*TableField
	Keys          []*TableField   //关键key
	SourceMap     map[string]bool //需要的source字段名
	RecordChecks  []*TableCheck   // 针对一行数据的检查
	GlobalChecks  []*TableCheck   // 只调一次的Check
	PostScript    string
	ExtraFields   []*TableField // 额外字段
	extraFieldMap map[string]bool
}

type TableSource struct {
	Table string
	Sheet string
}

func NewTableMeta(rtm *RawTableMeta) (*TableMeta, error) {
	t := &TableMeta{
		Target:        rtm.Target,
		Mode:          rtm.Mode,
		SourceType:    rtm.SourceType,
		SourceMap:     make(map[string]bool),
		PostScript:    rtm.PostScript,
		extraFieldMap: make(map[string]bool),
	}

	//source
	sourceLen := len(rtm.Sources)
	if sourceLen == 0 {
		return nil, errors.New(fmt.Sprintf("table meta config[%v] sources count must >= 0", rtm.Target))
	}
	t.Sources = make([]*TableSource, sourceLen)
	for index, source := range rtm.Sources {
		t.Sources[index] = &TableSource{
			Table: source.Table,
			Sheet: source.Sheet,
		}
	}

	//校验
	fields := make([]*TableField, 0)
	keysMap := make(map[int]*TableField)
	for _, rtf := range rtm.Fields {
		if !rtf.Active {
			continue
		}

		tft, err := newTableField(rtf)

		if err != nil {
			return nil, err
		}
		fields = append(fields, tft)

		if tft.Key > 0 {
			//同一个key不能重复
			if _, ok := keysMap[tft.Key]; ok {
				return nil, errors.New(fmt.Sprintf("table meta config[%v] key value repeat[%v]", rtm.Target, tft.Key))
			}
			//key只支持基础类型
			if !tft.Type.IsBaseType() {
				return nil, errors.New(fmt.Sprintf("table meta config[%v] key only support base type[%v]", rtm.Target, tft.Target))
			}
			keysMap[tft.Key] = tft

		}
	}

	if len(fields) == 0 {
		return nil, errors.New(fmt.Sprintf("table meta config[%v] not one active field_type", rtm.Target))
	}

	//检车key是否合法
	keyCount := len(keysMap)
	if keyCount == 0 {
		return nil, errors.New(fmt.Sprintf("table meta config[%v] key count must >= 0", rtm.Target))
	}
	keys := make([]*TableField, keyCount)
	for keyIndex, tft := range keysMap {
		if keyIndex > keyCount {
			return nil, errors.New(fmt.Sprintf("table meta config[%v] key value error[%v]", rtm.Target, keyIndex))
		}
		keys[keyIndex-1] = tft
	}

	t.Fields = fields
	t.Keys = keys

	//判断Source是否重复，并且生成需要原始字段的map
	filedMap := make(map[string]*TableField)
	for _, tf := range fields {
		if _, ok := filedMap[tf.Target]; ok {
			return nil, errors.New(fmt.Sprintf("table meta config[%v] target name repeated[%v]", rtm.Target, tf.Target))
		}
		filedMap[tf.Target] = tf
		t.SourceMap[tf.Source] = true
	}

	// 生成Check
	for _, rck := range rtm.Checks {
		if !rck.Active {
			continue
		}
		tck := newTableCheck(rck)
		if rck.Global {
			t.GlobalChecks = append(t.GlobalChecks, tck)
		} else {
			t.RecordChecks = append(t.RecordChecks, tck)
		}
	}

	return t, nil
}

func (tm *TableMeta) NotKeyFieldCount() int {
	return len(tm.Fields) - len(tm.Keys)
}

func (tm *TableMeta) GetKeyDefType(finalType field_type.EFieldType) *field_type.TableFieldType {
	return tm.GetKeyDefTypeOffset(finalType, 0)
}

func (tm *TableMeta) GetKeyDefTypeOffset(finalType field_type.EFieldType, offset int) *field_type.TableFieldType {
	value := field_type.NewTableFieldType(finalType)
	for i := len(tm.Keys) - 1; i >= offset; i-- {
		key := tm.Keys[i]
		value = field_type.NewTableFieldMapType(key.Type, value)
	}
	return value
}

func (tm *TableMeta) AddExtraField(name string, value *field_type.TableFieldType, desc string) {
	if value == nil {
		slog.Fatalf("table meta config[%v] extra field type is nil", tm.Target)
	}
	_, ok := tm.extraFieldMap[name]
	if ok {
		slog.Fatalf("table meta config[%v] extra field name repeated[%v]", tm.Target, name)
		return
	}

	tm.extraFieldMap[name] = true
	tm.ExtraFields = append(tm.ExtraFields, &TableField{
		Target: name,
		Type:   value,
		Desc:   desc,
	})
}
