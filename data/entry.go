package data

import (
	"errors"
	"table-export/data/model"
	"table-export/meta"
)

func GetDataModelByType(tableMetal *meta.TableMeta) (*model.TableModel, error) {
	loader, ok := dataSourceLoads[tableMetal.SourceType]
	if !ok {
		return nil, errors.New("can't support source model:" + tableMetal.SourceType)
	}

	dataModel, err := loader.LoadDataModel(tableMetal)

	return dataModel, err
}
