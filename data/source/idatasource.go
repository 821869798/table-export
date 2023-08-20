package source

import (
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/meta"
)

type IDataSource interface {
	LoadDataModel(tableMetal *meta.TableMeta) (*model.TableModel, error)
}
