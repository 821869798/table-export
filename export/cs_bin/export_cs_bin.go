package cs_bin

import (
	"github.com/gookit/slog"
	"os"
	"sync"
	"table-export/config"
	"table-export/convert/visitor"
	"table-export/data/model"
	"table-export/data/optimize"
	"table-export/export/api"
	"table-export/export/common"
	"table-export/meta"
	"table-export/serialization"
	"table-export/util"
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
		slog.Fatal("Export Json expect *RawMetaRuleUnitJson Rule Unit")
	}

	//清空目录
	if err := util.InitDirAndClearFile(config.AbsExeDir(csBinRule.DataTempDir), `^.*?\.bytes$`); err != nil {
		slog.Fatal(err)
	}

	if err := util.InitDirAndClearFile(config.AbsExeDir(csBinRule.CodeTempDir), `^.*?\.cs`); err != nil {
		slog.Fatal(err)
	}

	defer util.TimeCost(time.Now(), "export c# bin time cost = %v\n")

	//实际开始转换
	allDataModel := common.CommonMutilExport(e.tableMetas, func(dataModel *model.TableModel) {

	})

	//生成代码和二进制数据
	wg := sync.WaitGroup{}
	wg.Add(len(allDataModel))
	for _, tableModel := range allDataModel {
		go func(tm *model.TableModel) {
			if csBinRule.GenOptimizeData {
				//表数据优化
				optimize.OptimizeTableDataRepeat(tm, csBinRule.OptimizeThreshold)
			}
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

func GenCSBinData(dataModel *model.TableModel, ouputPath string) {
	byteBuff := serialization.NewByteBuf(1024)
	binary := visitor.NewBinary(byteBuff)
	binary.AcceptTable(dataModel)
	bytes := byteBuff.GetBytes()
	filePath := config.AbsExeDir(ouputPath, dataModel.Meta.Target+".bytes")
	file, err := os.Create(filePath)
	if err != nil {
		slog.Fatalf("export cs_bin bytes file create file error:%v", err)
	}
	_, err = file.Write(bytes)
	if err != nil {
		slog.Fatalf("export cs_bin bytes file write error:%v", err)
	}
}
