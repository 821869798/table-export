package meta

import (
	"errors"
	"fmt"
)

type TableMeta struct {
	Target     string
	Mode       string
	SourceType string
	Sources    []*TableSource
	Fields     []*TableField
	Keys       []*TableField   //关键key
	SourceMap  map[string]bool //需要的source字段名
}

type TableSource struct {
	Table string
	Sheet string
}

func NewTableMeta(rtm *RawTableMeta) (*TableMeta, error) {
	t := &TableMeta{
		Target:     rtm.Target,
		Mode:       rtm.Mode,
		SourceType: rtm.SourceType,
		SourceMap:  make(map[string]bool),
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
		return nil, errors.New(fmt.Sprintf("table meta config[%v] not one active field", rtm.Target))
	}

	//检车key是否合法
	keyCount := len(keysMap)
	if keyCount == 0 {
		return nil, errors.New(fmt.Sprintf("table meta config[%v] key count must >= 0", rtm.Target))
	}
	keys := make([]*TableField, keyCount, keyCount)
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

	return t, nil
}
