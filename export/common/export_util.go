package common

import (
	log "github.com/sirupsen/logrus"
	"sync"
	"table-export/data"
	"table-export/data/model"
	"table-export/define"
	"table-export/meta"
)

//通用的并行执行的方法a
func CommonMutilExport(tableMetas []*meta.RawTableMeta, exportFunc func(*model.TableModel)) {
	wg := sync.WaitGroup{}
	wg.Add(len(tableMetas))
	for _, tableMeta := range tableMetas {
		if tableMeta.Mode == define.CommentSymbol {
			//是注释模式，不导出
			wg.Done()
			continue
		}
		go func(tableMeta *meta.RawTableMeta) {
			tm, err := meta.NewTableMeta(tableMeta)
			if err != nil {
				log.Fatal(err)
			}

			dataModel, err := data.GetDataModelByType(tm)
			if err != nil {
				log.Fatal(err)
			}

			//执行函数
			exportFunc(dataModel)

			wg.Done()
		}(tableMeta)
	}
	wg.Wait()
}
