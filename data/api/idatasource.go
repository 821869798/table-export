package api

import (
	"table-export/data/model"
	"table-export/meta"
)

type DataSource interface {
	LoadDataModel(tableMetal *meta.RawTableMeta) (*model.DataModel, error)
}
