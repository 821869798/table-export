package api

import (
	"table-export/config"
	"table-export/meta"
)

type IExport interface {
	Export(config.MetaRuleUnit)
	TableMetas() []*meta.RawTableMeta
}
