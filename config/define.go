package config

type ExportType int

const (
	ExportType_None ExportType = iota
	ExportType_Lua
	ExportType_Json
	ExportType_CS_Bin
)

// MetaRuleUnit 普通版本的规则，freestyle
type MetaRuleUnit interface {
	RuleExportType() ExportType
}

// MetaRuleUnitPlus 高级版的规则，支持的特性更多
type MetaRuleUnitPlus interface {
	MetaRuleUnit
	ActiveOptimizeData() bool
	GetEnumFiles() []string
}
