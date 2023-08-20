package common

import (
	"github.com/821869798/table-export/consts"
	"github.com/821869798/table-export/data"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/meta"
	"github.com/gookit/slog"
	"sync"
)

// CommonMutilExport 通用的并行执行的方法
func CommonMutilExport(tableMetas []*meta.RawTableMeta, exportFunc func(*model.TableModel)) []*model.TableModel {
	wg := sync.WaitGroup{}
	wg.Add(len(tableMetas))
	allDataModel := make([]*model.TableModel, 0, len(tableMetas))
	var mutex sync.Mutex

	for _, tableMeta := range tableMetas {
		if tableMeta.Mode == consts.CommentSymbol {
			//是注释模式，不导出
			wg.Done()
			continue
		}
		go func(tableMeta *meta.RawTableMeta) {
			tm, err := meta.NewTableMeta(tableMeta)
			if err != nil {
				slog.Fatal(err)
			}

			dataModel, err := data.GetDataModelByType(tm)
			if err != nil {
				slog.Fatal(err)
			}

			//执行函数
			exportFunc(dataModel)

			mutex.Lock()
			allDataModel = append(allDataModel, dataModel)
			mutex.Unlock()

			wg.Done()
		}(tableMeta)
	}
	wg.Wait()

	return allDataModel
}

//func CommonMutilExportWithDTable(tableMetas []*meta.RawTableMeta, exportFunc func(*model.TableModel)) []*model.TableModel {
//
//}
