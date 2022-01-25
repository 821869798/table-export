package common

type ExportType int

const (
	ExportType_None ExportType = iota
	ExportType_Lua
	ExportType_Json
	ExportType_CS_Proto
)

var exportStrMapping = map[ExportType]string{
	ExportType_Lua:      "lua",
	ExportType_Json:     "json",
	ExportType_CS_Proto: "cs_proto",
}

func (e ExportType) ExportTypeToString() string {
	return exportStrMapping[e]
}
