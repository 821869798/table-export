package env

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/enum"
	"github.com/821869798/table-export/field_type"
)

// 全局环境变量
var (
	environment *Environment
)

// Environment 当前的环境
type Environment struct {
	metaRuleUnitPlus config.MetaRuleUnitPlus
	// enumFiles 枚举文件
	enumFiles []*enum.DefineEnumFile
	// enumMapping 枚举映射
	enumMapping map[string]*enum.DefineEnum
	// extFieldMapping 扩展字段类型映射
	extFieldMapping map[string]field_type.IExtFieldType
	// extFieldClassFiles 生成自定义Class文件
	extFieldClassFiles map[string][]*field_type.TableFieldType
}

func InitEnv() {
	environment = &Environment{
		enumMapping:        make(map[string]*enum.DefineEnum),
		extFieldMapping:    make(map[string]field_type.IExtFieldType),
		extFieldClassFiles: make(map[string][]*field_type.TableFieldType),
	}
}

func SetMetaRuleUnitPlus(metaRuleUnitPlus config.MetaRuleUnitPlus) {
	environment.metaRuleUnitPlus = metaRuleUnitPlus
}

func GetMetaRuleUnitPlus() config.MetaRuleUnitPlus {
	return environment.metaRuleUnitPlus
}
