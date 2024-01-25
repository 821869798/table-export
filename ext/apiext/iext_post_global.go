package apiext

import "github.com/821869798/table-export/data/model"

type IExtPostGlobal interface {
	PostGlobal(tableMap map[string]*model.TableModel) error
}
