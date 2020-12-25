package export

import (
	"table-export/export/api"
	"table-export/export/cs_proto"
	"table-export/export/json"
	"table-export/export/lua"
	"table-export/meta"
)

var exportCreators map[string]func([]*meta.RawTableMeta) api.IExport

func init() {
	exportCreators = make(map[string]func([]*meta.RawTableMeta) api.IExport)
	exportCreators["lua"] = lua.NewExportLua
	exportCreators["json"] = json.NewExportJson
	exportCreators["cs_proto"] = cs_proto.NewExportCsProto
}
