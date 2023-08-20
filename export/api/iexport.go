package api

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/meta"
)

type IExport interface {
	Export(config.MetaRuleUnit)
	TableMetas() []*meta.RawTableMeta
}
