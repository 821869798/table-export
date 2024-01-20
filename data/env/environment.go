package env

import (
	"errors"
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/enum"
)

// 全局环境变量
var (
	environment *Environment
)

// Environment 当前的环境
type Environment struct {
	enumFiles   []*enum.DefineEnumFile
	enumMapping map[string]*enum.DefineEnum
}

func InitEnv() {
	environment = &Environment{
		enumMapping: make(map[string]*enum.DefineEnum),
	}
}

func AddEnumDefines(enumConfigs []*config.RawMetaEnumConfig) error {
	fileMapping := make(map[string]*enum.DefineEnumFile)
	for _, enumConfig := range enumConfigs {
		enumFile := enum.NewDefineEnumFile(enumConfig)
		for _, enumDefine := range enumFile.Enums {
			_, ok := environment.enumMapping[enumDefine.Name]
			if ok {
				return errors.New(fmt.Sprintf("enum name repeated: name[%v] ", enumDefine.Name))
			}
			if err := enumDefine.InitEnum(); err != nil {
				return err
			}
			environment.enumMapping[enumDefine.Name] = enumDefine
		}
		prevEnumFile, ok := fileMapping[enumFile.FileName]
		if ok {
			// 合并文件名定义相同的
			prevEnumFile.Enums = append(prevEnumFile.Enums, enumFile.Enums...)
		} else {
			environment.enumFiles = append(environment.enumFiles, enumFile)
			fileMapping[enumFile.FileName] = enumFile
		}
	}
	return nil
}

func GetEnumDefine(name string) (*enum.DefineEnum, bool) {
	v, ok := environment.enumMapping[name]
	return v, ok
}

func EnumFiles() []*enum.DefineEnumFile {
	return environment.enumFiles
}
