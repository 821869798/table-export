package config

type ExportType int

const (
	ExportType_None ExportType = iota
	ExportType_Lua
	ExportType_Json
	ExportType_CS_Bin
)

type MetaRuleUnit interface {
	RuleExportType() ExportType
}

type MetaRuleUnitPlus interface {
	MetaRuleUnit
	ActiveOptimizeData() bool
}
