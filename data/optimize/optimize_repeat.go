package optimize

import (
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/meta"
)

// OptimizeTableDataRepeat 优化重复的数据
func OptimizeTableDataRepeat(dataModel *model.TableModel, threshold float32) {
	//除了Key以外没有其他字段，不需要优化
	if dataModel.Meta.NotKeyFieldCount() <= 0 {
		return
	}
	refTableFields := make([]*OptimizeProcess, 0, dataModel.Meta.NotKeyFieldCount())
	for _, f := range dataModel.Meta.Fields {
		if f.Key > 0 {
			//key不需要生成
			continue
		}
		if f.Type.IsReferenceType() {
			refTableFields = append(refTableFields, newOptimizeProcess(f))
		}
	}

	// 检测每个字段的重复是否超过阈值
	for _, rowData := range dataModel.RawData {
		for _, refField := range refTableFields {
			rawIndex := dataModel.NameIndexMapping[refField.Field.Target]
			var rawStr string
			if rawIndex < len(rowData) {
				rawStr = rowData[rawIndex]
			}
			count, ok := refField.CellCounts[rawStr]
			if ok {
				refField.CellCounts[rawStr] = count + 1
				refField.RepeatedCount++
			} else {
				refField.CellCounts[rawStr] = 1
			}
		}
	}

	result := genOptimizeRefTableData(dataModel, threshold, refTableFields)
	if result != nil {
		dataModel.Optimize = result
	}
}

func genOptimizeRefTableData(dataModel *model.TableModel, threshold float32, refTableFields []*OptimizeProcess) *model.TableOptimize {
	tableOptimize := model.NewTableOptimize()
	optimizeFields := make([]*OptimizeProcess, 0)
	for _, refField := range refTableFields {
		if float32(refField.RepeatedCount)/float32(len(dataModel.RawData)) >= threshold {
			refField.TableOptimizeField = model.NewTableOptimizeField(refField.Field, len(refField.CellCounts), len(dataModel.RawData))
			refField.CellCounts = make(map[string]int, 0)
			refField.DataIndex = 0
			optimizeFields = append(optimizeFields, refField)

			tableOptimize.AddOptimizeField(refField.TableOptimizeField)
		}
	}
	if len(optimizeFields) < 0 {
		return nil
	}
	for i, rowData := range dataModel.RawData {
		for _, refField := range optimizeFields {
			rawIndex := dataModel.NameIndexMapping[refField.Field.Target]
			var rawStr string
			if rawIndex < len(rowData) {
				rawStr = rowData[rawIndex]
			}
			lastIndex, ok := refField.CellCounts[rawStr]
			if ok {
				refField.TableOptimizeField.DataIndexs[i] = lastIndex
			} else {
				refField.TableOptimizeField.OriginDatas[refField.DataIndex] = rawStr
				refField.TableOptimizeField.DataIndexs[i] = refField.DataIndex
				refField.CellCounts[rawStr] = refField.DataIndex
				refField.DataIndex++
			}
		}
	}
	return tableOptimize
}

type OptimizeProcess struct {
	Field              *meta.TableField
	CellCounts         map[string]int
	RepeatedCount      int
	DataIndex          int
	TableOptimizeField *model.TableOptimizeField
}

func newOptimizeProcess(f *meta.TableField) *OptimizeProcess {
	op := &OptimizeProcess{
		Field:         f,
		CellCounts:    make(map[string]int, 0),
		RepeatedCount: 0,
	}
	return op
}
