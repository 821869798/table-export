package lua

import (
	"sync"
	"table-export/constant"
	"table-export/export/api"
	"table-export/meta"
)

type ExportLua struct {
	tableMetas []*meta.RawTableMeta
}

func NewExportLua(tableMetas []*meta.RawTableMeta) api.IExport {
	e := &ExportLua{
		tableMetas: tableMetas,
	}
	return e
}

func (e *ExportLua) TableMetas() []*meta.RawTableMeta {
	return e.tableMetas
}

func (e *ExportLua) Export() {
	wg := sync.WaitGroup{}
	wg.Add(len(e.tableMetas))
	for _, tableMeta := range e.tableMetas {
		if tableMeta.Mode == constant.CommentSymbol {
			//是注释模式，不导出
			wg.Done()
			continue
		}
		go func(tableMeta *meta.RawTableMeta) {

			wg.Done()
		}(tableMeta)
	}
	wg.Wait()
}
