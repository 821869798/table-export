package source

import (
	"errors"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/meta"
)

func GetDataModelByType(tableMetal *meta.TableMeta) (*model.TableModel, error) {
	loader, ok := dataSourceLoads[tableMetal.SourceType]
	if !ok {
		return nil, errors.New("can't support source model:" + tableMetal.SourceType)
	}

	dataModel, err := loader.LoadDataModel(tableMetal)

	return dataModel, err
}
