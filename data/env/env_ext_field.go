package env

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/field_type"
	"strings"
)

func AddExtFieldType(extFieldType field_type.IExtFieldType) error {
	name := strings.ToLower(extFieldType.Name())
	_, ok := environment.extFieldMapping[name]
	if ok {
		return errors.New(fmt.Sprintf("ext field_type type name repeated: name[%v] ", name))
	}
	environment.extFieldMapping[name] = extFieldType
	fieldType := extFieldType.TableFieldType()
	if fieldType.Type == field_type.EFieldType_Class {
		environment.extFieldClassFiles[extFieldType.DefineFile()] = append(environment.extFieldClassFiles[extFieldType.DefineFile()], fieldType)
	}
	return nil
}

func GetExtFieldType(name string) (field_type.IExtFieldType, bool) {
	v, ok := environment.extFieldMapping[strings.ToLower(name)]
	return v, ok
}

func GetExtFieldClassFiles() map[string][]*field_type.TableFieldType {
	return environment.extFieldClassFiles
}
