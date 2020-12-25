package json

import (
	"table-export/export/api"
	"table-export/meta"
)

type ExportJson struct {
	tableMetas []*meta.RawTableMeta
}

func NewExportJson(tableMetas []*meta.RawTableMeta) api.IExport {
	e := &ExportJson{
		tableMetas: tableMetas,
	}
	return e
}

func (e *ExportJson) TableMetas() []*meta.RawTableMeta {
	return e.tableMetas
}

func (e *ExportJson) Export() {

}
