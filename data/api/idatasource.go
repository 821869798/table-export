package api

import (
	"table-export/data/model"
	"table-export/meta"
)

type DataSource interface {
	LoadDataModel(tableMetal *meta.TableMeta) (*model.TableModel, error)
}
