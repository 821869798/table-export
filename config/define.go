package config

type ExportType int

const (
	ExportType_None ExportType = iota
	ExportType_Lua
	ExportType_Json
	ExportType_CS_Proto
	ExportType_CS_Bin
)

type MetaRuleUnit interface {
	RuleExportType() ExportType
}
