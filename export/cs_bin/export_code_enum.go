package cs_bin

import (
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/data/enum"
	"github.com/821869798/table-export/util"
	"github.com/gookit/slog"
	"os"
	"path/filepath"
	"text/template"
)

func GenCSBinCodeEnum(enumFile *enum.DefineEnumFile, csBinRule *config.RawMetaRuleUnitCSBin, outputPath string) {

	templateRoot := &CSCodeWriteEnumFile{
		NameSpace: util.ReplaceWindowsLineEnd(csBinRule.GenCodeNamespace),
		Enums:     enumFile.Enums,
	}

	//目标路径
	filePath := filepath.Join(outputPath, enumFile.FileName+".cs")
	file, err := os.Create(filePath)
	if err != nil {
		slog.Fatal(err)
	}

	//创建模板,绑定全局函数,并且解析
	tmpl, err := template.New("cs_bin_enum").Funcs(template.FuncMap{}).Parse(template_CS_Enum)

	//渲染输出
	err = tmpl.Execute(file, templateRoot)
	if err != nil {
		slog.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		slog.Fatal(err)
	}
}

type CSCodeWriteEnumFile struct {
	NameSpace string
	Enums     []*enum.DefineEnum
}
