package api

import "table-export/meta"

type IExport interface {
	Export()
	TableMetas() []*meta.RawTableMeta
}
