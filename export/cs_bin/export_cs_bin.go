package cs_bin

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/visitor"
	"github.com/821869798/table-export/data/check"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/export/api"
	"github.com/821869798/table-export/export/common"
	"github.com/821869798/table-export/meta"
	"github.com/821869798/table-export/serialization"
	"github.com/821869798/table-export/util"
	"github.com/gookit/slog"
	"os"
	"sync"
	"time"
)

type ExportCSBin struct {
	tableMetas []*meta.RawTableMeta
	extraArg   map[string]string
}

func NewExportCSBin(tableMetas []*meta.RawTableMeta, extraArg map[string]string) api.IExport {
	e := &ExportCSBin{
		tableMetas: tableMetas,
		extraArg:   extraArg,
	}
	return e
}

func (e *ExportCSBin) TableMetas() []*meta.RawTableMeta {
	return e.tableMetas
}

func (e *ExportCSBin) Export(ru config.MetaRuleUnit) {

	csBinRule, ok := ru.(*config.RawMetaRuleUnitCSBin)
	if !ok {
		slog.Fatal("Export CSBin expect *RawMetaRuleUnitJson Rule Unit")
	}

	//清空目录
	if err := util.InitDirAndClearFile(util.RelExecuteDir(csBinRule.DataTempDir), `^.*?\.bytes$`); err != nil {
		slog.Fatal(err)
	}

	if err := util.InitDirAndClearFile(util.RelExecuteDir(csBinRule.CodeTempDir), `^.*?\.cs`); err != nil {
		slog.Fatal(err)
	}

	defer util.TimeCost(time.Now(), "export c# bin time cost = %v\n")

	//实际开始转换
	allDataModel := common.ExportPlusParallel(e.tableMetas, csBinRule, func(dataModel *model.TableModel) {

	})

	// TODO 表的数据后处理

	// 表的数据检查
	global := make(map[string]map[interface{}]interface{}, len(allDataModel))
	for _, m := range allDataModel {
		global[m.Meta.Target] = m.MemTable.RawDataMapping()
	}
	wgCheck := sync.WaitGroup{}
	wgCheck.Add(len(allDataModel))
	for _, m := range allDataModel {
		go func(m *model.TableModel) {
			check.Run(m, global)
			wgCheck.Done()
		}(m)
	}
	wgCheck.Wait()

	//生成代码和二进制数据
	wg := sync.WaitGroup{}
	wg.Add(len(allDataModel))
	for _, tableModel := range allDataModel {
		go func(tm *model.TableModel) {
			GenCSBinCode(tm, csBinRule, csBinRule.CodeTempDir)
			GenCSBinData(tm, csBinRule.DataTempDir)
			wg.Done()
		}(tableModel)
	}
	wg.Wait()

	// 拷贝到最终的目录去
	err := util.CopyDir(csBinRule.DataTempDir, csBinRule.DataBinDir)
	if err != nil {
		slog.Fatal("Copy files error:" + err.Error())
	}
	err = util.CopyDir(csBinRule.CodeTempDir, csBinRule.GenCodeDir)
	if err != nil {
		slog.Fatal("Copy files error:" + err.Error())
	}
}

func GenCSBinData(dataModel *model.TableModel, outputPath string) {
	byteBuff := serialization.NewByteBuf(1024)
	binary := visitor.NewBinary(byteBuff)
	binary.AcceptTable(dataModel)
	bytes := byteBuff.GetBytes()
	filePath := util.RelExecuteDir(outputPath, dataModel.Meta.Target+".bytes")
	file, err := os.Create(filePath)
	if err != nil {
		slog.Fatalf("export cs_bin bytes file create file error:%v", err)
	}
	_, err = file.Write(bytes)
	if err != nil {
		slog.Fatalf("export cs_bin bytes file write error:%v", err)
	}
}
