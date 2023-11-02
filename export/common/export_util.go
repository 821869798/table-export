package common

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/constant"
	"github.com/821869798/table-export/data"
	"github.com/821869798/table-export/data/memory_table"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/data/optimize"
	"github.com/821869798/table-export/meta"
	"github.com/gookit/slog"
	"sync"
)

// ExportParallel 通用的并行执行的方法
func ExportParallel(tableMetas []*meta.RawTableMeta, exportFunc func(*model.TableModel)) []*model.TableModel {
	wg := sync.WaitGroup{}
	wg.Add(len(tableMetas))
	allDataModel := make([]*model.TableModel, 0, len(tableMetas))
	var mutex sync.Mutex

	for _, tableMeta := range tableMetas {
		if tableMeta.Mode == constant.CommentSymbol {
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

// ExportPlusParallel 增强版的并行导出，生成中间层数据，支持配置中预处理，自定义类型等功能
func ExportPlusParallel(tableMetas []*meta.RawTableMeta, ruleUnit config.MetaRuleUnitPlus, exportFunc func(*model.TableModel)) []*model.TableModel {
	wg := sync.WaitGroup{}
	wg.Add(len(tableMetas))
	allDataModel := make([]*model.TableModel, 0, len(tableMetas))
	var mutex sync.Mutex

	for _, tableMeta := range tableMetas {
		if tableMeta.Mode == constant.CommentSymbol {
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

			memoryTable, err := memory_table.NewMemTableCommon(dataModel, len(dataModel.RawData), len(dataModel.RawData))

			if err != nil {
				slog.Fatal(err)
			}
			dataModel.MemTable = memoryTable
			if ruleUnit.ActiveOptimizeData() {
				optimize.OptimizeTableDataRepeat(dataModel)
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
