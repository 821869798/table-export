package export

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/export/api"
	"github.com/821869798/table-export/export/cs_bin"
	"github.com/821869798/table-export/export/json"
	"github.com/821869798/table-export/export/lua"
	"github.com/821869798/table-export/meta"
)

var exportCreators map[config.ExportType]func([]*meta.RawTableMeta, map[string]string) api.IExport

func init() {
	exportCreators = make(map[config.ExportType]func([]*meta.RawTableMeta, map[string]string) api.IExport)
	exportCreators[config.ExportType_Lua] = lua.NewExportLua
	exportCreators[config.ExportType_Json] = json.NewExportJson
	exportCreators[config.ExportType_CS_Bin] = cs_bin.NewExportCSBin
}
