package export

import (
	"table-export/export/api"
	"table-export/export/common"
	"table-export/export/cs_proto"
	"table-export/export/json"
	"table-export/export/lua"
	"table-export/meta"
)

var exportCreators map[string]func([]*meta.RawTableMeta) api.IExport

func init() {
	exportCreators = make(map[string]func([]*meta.RawTableMeta) api.IExport)
	exportCreators[common.ExportType_Lua.ExportTypeToString()] = lua.NewExportLua
	exportCreators[common.ExportType_Json.ExportTypeToString()] = json.NewExportJson
	exportCreators[common.ExportType_CS_Proto.ExportTypeToString()] = cs_proto.NewExportCsProto
}
