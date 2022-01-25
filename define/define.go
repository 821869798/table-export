package define

type ExportType int

const (
	ExportType_None ExportType = iota
	ExportType_Lua
	ExportType_Json
	ExportType_CS_Proto
)

var exportStrMapping = map[string]ExportType{
	"lua":      ExportType_Lua,
	"json":     ExportType_Json,
	"cs_proto": ExportType_CS_Proto,
}

func GetExportTypeFromString(value string) (ExportType, bool) {
	result, ok := exportStrMapping[value]
	return result, ok
}
