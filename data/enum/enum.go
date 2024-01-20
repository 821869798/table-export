package enum

import (
	"cmp"
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"slices"
)

type DefineEnumFile struct {
	FileName string
	Enums    []*DefineEnum
}

type DefineEnum struct {
	Name            string
	Parse4String    bool
	Values          []*config.RawMetaEnumValue
	ValueMapping    map[int32]*config.RawMetaEnumValue
	ValueStrMapping map[string]*config.RawMetaEnumValue
}

func NewNewDefineEnum(enumDefine *config.RawMetaEnumDefine) *DefineEnum {
	e := &DefineEnum{
		Name:         enumDefine.Name,
		Parse4String: enumDefine.Parse4String,
		Values:       enumDefine.Values,
		ValueMapping: make(map[int32]*config.RawMetaEnumValue, len(enumDefine.Values)),
	}
	if e.Parse4String {
		e.ValueStrMapping = make(map[string]*config.RawMetaEnumValue, len(enumDefine.Values))
	}
	return e
}

func (d *DefineEnum) InitEnum() error {
	for _, enumValue := range d.Values {
		_, ok := d.ValueMapping[enumValue.Index]
		if ok {
			return errors.New(fmt.Sprintf("enum value repeated: name[%s] index[%d]", d.Name, enumValue.Index))
		}
		d.ValueMapping[enumValue.Index] = enumValue
		if d.Parse4String {
			_, ok = d.ValueStrMapping[enumValue.ValueString]
			if ok {
				return errors.New(fmt.Sprintf("enum value repeated: name[%s] valueString[%s]", d.Name, enumValue.ValueString))
			}
			d.ValueStrMapping[enumValue.ValueString] = enumValue
		}
	}

	_, ok := d.ValueMapping[0]
	if !ok {
		return errors.New(fmt.Sprintf("enum value must hava index 0: name[%s]", d.Name))
	}

	// 排序
	slices.SortFunc(d.Values, func(a, b *config.RawMetaEnumValue) int {
		return cmp.Compare(a.Index, b.Index)
	})
	return nil
}

func NewDefineEnumFile(enumConfig *config.RawMetaEnumConfig) *DefineEnumFile {
	enumFile := &DefineEnumFile{
		Enums:    make([]*DefineEnum, 0, len(enumConfig.EnumDefines)),
		FileName: enumConfig.EnumFileName,
	}
	for _, enumDefine := range enumConfig.EnumDefines {
		enumFile.Enums = append(enumFile.Enums, NewNewDefineEnum(enumDefine))
	}
	return enumFile
}
