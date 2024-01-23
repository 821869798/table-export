package cs_bin

import (
	"github.com/821869798/fankit/fanstr"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/printer"
	"github.com/821869798/table-export/convert/wrap"
	"github.com/821869798/table-export/field_type"
	"github.com/gookit/slog"
	"os"
	"path/filepath"
	"text/template"
)

func GenCSBinCodeExtClass(fileName string, extFieldClasses []*field_type.TableFieldType, csBinRule *config.RawMetaRuleUnitCSBin, outputPath string) {

	// 生成模板数据
	templateRoot := &CSCodeWriteExtClassFile{
		CodeHead:          fanstr.ReplaceWindowsLineEnd(csBinRule.GenCodeHead),
		NameSpace:         fanstr.ReplaceWindowsLineEnd(csBinRule.GenCodeNamespace),
		ClassDefinePrefix: csBinRule.GetClassDefinePrefix(),
	}

	for _, extFieldClass := range extFieldClasses {
		templateExtClass := &CSCodeWriteExtClass{
			Name: extFieldClass.Name,
		}
		for _, extField := range extFieldClass.Class.AllFields() {
			typeDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, extField.Type, false)
			if err != nil {
				slog.Fatalf("gen cs code error:%v", typeDef)
			}
			templateExtClass.Fields = append(templateExtClass.Fields, &CSCodeWriteExtClassField{
				TypeDef:   typeDef,
				Name:      extField.Name,
				FieldType: extField.Type,
			})
		}
		templateRoot.ClassTypes = append(templateRoot.ClassTypes, templateExtClass)
	}

	//创建代理打印器
	codePrinter := printer.NewCSBinaryPrint(false)

	//目标路径
	filePath := filepath.Join(outputPath, fileName+".cs")
	file, err := os.Create(filePath)
	if err != nil {
		slog.Fatal(err)
	}

	//创建模板,绑定全局函数,并且解析
	tmpl, err := template.New("cs_bin_class").Funcs(template.FuncMap{
		"CSbinFieldReader": func(fieldType *field_type.TableFieldType, fieldName string, reader string) string {
			return codePrinter.AcceptField(fieldType, fieldName, reader)
		},
	}).Parse(templateCSCodeClass)

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

type CSCodeWriteExtClassFile struct {
	CodeHead          string
	NameSpace         string
	ClassDefinePrefix string
	ClassTypes        []*CSCodeWriteExtClass
}

type CSCodeWriteExtClass struct {
	Name   string
	Fields []*CSCodeWriteExtClassField
}

type CSCodeWriteExtClassField struct {
	TypeDef   string
	Name      string
	FieldType *field_type.TableFieldType
}
