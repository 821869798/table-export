package export

import (
	"table-export/config"
	"table-export/export/api"
	"table-export/export/cs_bin"
	"table-export/export/cs_proto"
	"table-export/export/json"
	"table-export/export/lua"
	"table-export/meta"
)

var exportCreators map[config.ExportType]func([]*meta.RawTableMeta, map[string]string) api.IExport

func init() {
	exportCreators = make(map[config.ExportType]func([]*meta.RawTableMeta, map[string]string) api.IExport)
	exportCreators[config.ExportType_Lua] = lua.NewExportLua
	exportCreators[config.ExportType_Json] = json.NewExportJson
	exportCreators[config.ExportType_CS_Proto] = cs_proto.NewExportCsProto
	exportCreators[config.ExportType_CS_Bin] = cs_bin.NewExportCSBin
}
