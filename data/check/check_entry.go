package check

import (
	"github.com/821869798/table-export/data/model"
	"github.com/gookit/slog"
	"os"
)

func Run(dataMode *model.TableModel, global map[string]map[interface{}]interface{}) {
	if dataMode.MemTable == nil {
		return
	}
	meta := dataMode.Meta
	for _, tck := range meta.RecordChecks {
		err := tck.CompileCode(dataMode.MemTable.RawDataMapping(), global)
		if err != nil {
			slog.Fatalf("table[%s] check compile code[%s] error:%v", meta.Target, tck.Code, err)
			os.Exit(1)
		}
	}

	// 检查每条数据
	for _, record := range dataMode.MemTable.RawDataList() {
		for _, tck := range meta.RecordChecks {
			checkResult, err := tck.Run(record)
			if err != nil {
				slog.Fatalf("table[%s] check run code[%s] error:%v", meta.Target, tck.Code, err)
				os.Exit(1)
			}
			if !checkResult {
				slog.Fatalf("table[%s] check run code[%s] check faield:%v", meta.Target, tck.Code, record)
				os.Exit(1)
			}
		}
	}

	// 检查全局数据
	for _, tck := range meta.GlobalChecks {
		err := tck.CompileCode(dataMode.MemTable.RawDataMapping(), global)
		if err != nil {
			slog.Fatalf("table[%s] check compile global code[%s] error:%v", meta.Target, tck.Code, err)
			os.Exit(1)
		}
		checkResult, err := tck.RunGlobal()
		if err != nil {
			slog.Fatalf("table[%s] check run global code[%s] error:%v", meta.Target, tck.Code, err)
			os.Exit(1)
		}
		if !checkResult {
			slog.Fatalf("table[%s] check run global code[%s] check faield", meta.Target, tck.Code)
			os.Exit(1)
		}
	}
}
