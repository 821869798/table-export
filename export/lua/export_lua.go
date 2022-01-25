package lua

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"table-export/config"
	"table-export/data/model"
	"table-export/export/api"
	"table-export/export/common"
	"table-export/meta"
	"table-export/util"
	"text/template"
)

type ExportLua struct {
	tableMetas []*meta.RawTableMeta
}

func NewExportLua(tableMetas []*meta.RawTableMeta, extraArg map[string]string) api.IExport {
	e := &ExportLua{
		tableMetas: tableMetas,
	}
	return e
}

func (e *ExportLua) TableMetas() []*meta.RawTableMeta {
	return e.tableMetas
}

func (e *ExportLua) Export() {

	luaRule := config.GlobalConfig.Meta.RuleLua

	//清空目录
	if err := util.ClearDirAndCreateNew(luaRule.TempDir); err != nil {
		log.Fatal(err)
	}

	//实际开始转换
	common.CommonMutilExport(e.tableMetas, func(dataModel *model.TableModel) {
		exportLuaFile(dataModel, luaRule.TempDir)
	})

}

type luaWriteContent struct {
	Name    string
	Content interface{}
}

func exportLuaFile(dataModel *model.TableModel, outputPath string) {
	writeContent := common.CommonGetWriteContentMap(dataModel)

	//创建,绑定全局函数,并且解析
	tmpl, err := template.New("export_lua").Funcs(template.FuncMap{
		"IsString": func(v interface{}) bool {
			_, ok := v.(string)
			return ok
		},
		"ToString": func(v interface{}) string {
			str, ok := v.(string)
			if !ok {
				return ""
			}
			return str
		},
	}).Parse(templateLua)

	filePath := filepath.Join(outputPath, dataModel.Meta.Target+".lua")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	//填充数据
	result := &luaWriteContent{
		Name:    dataModel.Meta.Target,
		Content: writeContent,
	}

	//渲染输出
	err = tmpl.Execute(file, result)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

}
