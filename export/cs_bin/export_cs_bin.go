package cs_bin

import (
	"github.com/821869798/fankit/fanpath"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/visitor"
	"github.com/821869798/table-export/data/enum"
	"github.com/821869798/table-export/data/env"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/export/api"
	"github.com/821869798/table-export/export/common"
	"github.com/821869798/table-export/field_type"
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
	if err := fanpath.InitDirAndClearFile(fanpath.RelExecuteDir(csBinRule.DataTempDir), `^.*?\.bytes$`); err != nil {
		slog.Fatal(err)
	}

	if err := fanpath.InitDirAndClearFile(fanpath.RelExecuteDir(csBinRule.CodeTempDir), `^.*?\.cs`); err != nil {
		slog.Fatal(err)
	}

	defer util.TimeCost(time.Now(), "export c# bin time cost = %v\n")

	// 生成导出数据
	allDataModel := common.ExportPlusCommon(e.tableMetas, csBinRule)

	//生成代码和二进制数据
	wg := sync.WaitGroup{}
	wg.Add(len(allDataModel))
	for _, tableModel := range allDataModel {
		go func(tm *model.TableModel) {
			GenCSBinCodeTable(tm, csBinRule, csBinRule.CodeTempDir)
			GenCSBinData(tm, csBinRule.DataTempDir)
			wg.Done()
		}(tableModel)
	}
	// 生成枚举代码文件
	enumFiles := env.EnumFiles()
	wg.Add(len(enumFiles))
	for _, enumFile := range enumFiles {
		go func(ef *enum.DefineEnumFile) {
			GenCSBinCodeEnum(ef, csBinRule, csBinRule.CodeTempDir)
			wg.Done()
		}(enumFile)
	}

	// 生成自定义Class定义文件
	extFieldClassFiles := env.GetExtFieldClassFiles()
	wg.Add(len(extFieldClassFiles))
	for n, e := range extFieldClassFiles {
		go func(fileName string, extFieldClasses []*field_type.TableFieldType) {
			GenCSBinCodeExtClass(fileName, extFieldClasses, csBinRule, csBinRule.CodeTempDir)
			wg.Done()
		}(n, e)
	}
	// 等待完成
	wg.Wait()

	//清空目录
	if err := fanpath.InitDirAndClearFile(fanpath.RelExecuteDir(csBinRule.DataBinDir), `^.*?\.bytes$`); err != nil {
		slog.Fatal(err)
	}

	if err := fanpath.InitDirAndClearFile(fanpath.RelExecuteDir(csBinRule.GenCodeDir), `^.*?\.cs$`); err != nil {
		slog.Fatal(err)
	}

	// 拷贝到最终的目录去
	err := fanpath.CopyDir(csBinRule.DataTempDir, csBinRule.DataBinDir)
	if err != nil {
		slog.Fatal("Copy files error:" + err.Error())
	}
	err = fanpath.CopyDir(csBinRule.CodeTempDir, csBinRule.GenCodeDir)
	if err != nil {
		slog.Fatal("Copy files error:" + err.Error())
	}
}

func GenCSBinData(dataModel *model.TableModel, outputPath string) {
	byteBuff := serialization.NewByteBuf(1024)
	binary := visitor.NewBinary(byteBuff)
	binary.AcceptTable(dataModel)
	bytes := byteBuff.GetBytes()
	filePath := fanpath.RelExecuteDir(outputPath, dataModel.Meta.Target+".bytes")
	file, err := os.Create(filePath)
	if err != nil {
		slog.Fatalf("export cs_bin bytes file create file error:%v", err)
	}
	_, err = file.Write(bytes)
	if err != nil {
		slog.Fatalf("export cs_bin bytes file write error:%v", err)
	}
}
