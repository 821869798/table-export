package apiext

import "github.com/821869798/table-export/data/model"

// IExtPostTable 对一家表的后处理
type IExtPostTable interface {
	PostTable(model *model.TableModel) error
}
