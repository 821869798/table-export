package json

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"table-export/config"
	"table-export/data/model"
	"table-export/export/api"
	"table-export/export/common"
	"table-export/meta"
	"table-export/util"
)

type ExportJson struct {
	tableMetas []*meta.RawTableMeta
}

func NewExportJson(tableMetas []*meta.RawTableMeta, extraArg map[string]string) api.IExport {
	e := &ExportJson{
		tableMetas: tableMetas,
	}
	return e
}

func (e *ExportJson) TableMetas() []*meta.RawTableMeta {
	return e.tableMetas
}

func (e *ExportJson) Export() {
	jsonRule := config.GlobalConfig.Meta.RuleJson

	//清空目录
	if err := util.ClearDirAndCreateNew(jsonRule.OutputDir); err != nil {
		log.Fatal(err)
	}

	//实际开始转换
	common.CommonMutilExport(e.tableMetas, func(dataModel *model.TableModel) {
		exportJsonFile(dataModel, jsonRule.OutputDir)
	})

}

func exportJsonFile(dataModel *model.TableModel, outputPath string) {
	writeContent := common.CommonGetWriteContentMap(dataModel)

	b, err := json.MarshalIndent(writeContent, "", "    ")
	if err != nil {
		log.Fatalf("export json error:%v", err)
	}

	filePath := filepath.Join(outputPath, dataModel.Meta.Target+".json")

	f, err := os.Create(filePath) //os.OpenFile(destFile, os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("export json error:%v", err)
	}

	_, err = f.Write(b)
	if err != nil {
		log.Fatalf("export json error:%v", err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalf("export json error:%v", err)
	}
}
