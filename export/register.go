package export

import (
	"table-export/define"
	"table-export/export/api"
	"table-export/export/cs_proto"
	"table-export/export/json"
	"table-export/export/lua"
	"table-export/meta"
)

var exportCreators map[define.ExportType]func([]*meta.RawTableMeta, map[string]string) api.IExport

func init() {
	exportCreators = make(map[define.ExportType]func([]*meta.RawTableMeta, map[string]string) api.IExport)
	exportCreators[define.ExportType_Lua] = lua.NewExportLua
	exportCreators[define.ExportType_Json] = json.NewExportJson
	exportCreators[define.ExportType_CS_Proto] = cs_proto.NewExportCsProto
}
