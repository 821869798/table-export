package cs_proto

import (
	"table-export/export/api"
	"table-export/meta"
)

type ExportCsProto struct {
	tableMetas []*meta.RawTableMeta
}

func NewExportCsProto(tableMetas []*meta.RawTableMeta) api.IExport {
	e := &ExportCsProto{
		tableMetas: tableMetas,
	}
	return e
}

func (e *ExportCsProto) TableMetas() []*meta.RawTableMeta {
	return e.tableMetas
}

func (e *ExportCsProto) Export() {

}
