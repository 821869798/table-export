package cs_bin

import (
	"fmt"
	"github.com/gookit/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"table-export/config"
	"table-export/convert/printer"
	"table-export/convert/wrap"
	"table-export/data/model"
	"table-export/meta"
	"table-export/util"
	"text/template"
)

func GenCSBinCode(dataModel *model.TableModel, csBinRule *config.RawMetaRuleUnitCSBin, outputPath string) {
	//集合是否用只读的类型
	collectionReadonly := false

	templateRoot := &CSCodeWriteContent{
		NameSpace:          csBinRule.GenCodeNamespace,
		CodeHead:           csBinRule.GenCodeHead,
		TableName:          dataModel.Meta.Target,
		RecordClassName:    "Cfg" + dataModel.Meta.Target,
		TableClassName:     "Tbl" + dataModel.Meta.Target,
		TableCommonName:    "_TbCommon" + dataModel.Meta.Target,
		CollectionReadonly: collectionReadonly,
		CSCodeWriteFields:  make([]*CSCodeWriteField, 0, len(dataModel.Meta.Fields)),
		TableOptimize:      dataModel.Optimize,
		TableModel:         dataModel,
	}

	templateRoot.KeyDefTypeMap = getKeyDefTypeMap(dataModel, templateRoot.RecordClassName, 0)

	for _, field := range dataModel.Meta.Fields {
		typeDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, field.Type, collectionReadonly)
		if err != nil {
			slog.Fatalf("gen cs code error:%v", typeDef)
		}

		writeField := &CSCodeWriteField{
			TypeDef:            typeDef,
			Name:               field.Target,
			Desc:               field.Desc,
			Field:              field,
			OptimizeFieldIndex: -1,
		}
		if dataModel.Optimize != nil {
			optimizeField, optimizeIndex := dataModel.Optimize.GetOptimizeField(field)
			if optimizeField != nil {
				writeField.OptimizeFieldIndex = optimizeIndex
			}
		}
		templateRoot.CSCodeWriteFields = append(templateRoot.CSCodeWriteFields, writeField)
	}

	//创建代理打印器
	codePrinter := printer.NewCSBinaryPrint(collectionReadonly)

	//目标路径
	filePath := filepath.Join(outputPath, templateRoot.TableClassName+".cs")
	file, err := os.Create(filePath)
	if err != nil {
		slog.Fatal(err)
	}

	//创建模板,绑定全局函数,并且解析
	tmpl, err := template.New("export_cs_code").Funcs(template.FuncMap{
		"CSbinFieldReader": func(fieldType *meta.TableFieldType, fieldName string, reader string) string {
			return codePrinter.AcceptField(fieldType, fieldName, reader)
		},
		"CSbinFieldReaderEx": func(writeField *CSCodeWriteField, reader string, commonDataName string) string {
			if writeField.OptimizeFieldIndex >= 0 {
				return codePrinter.AcceptOptimizeAssignment(writeField.Name, reader, commonDataName+strconv.Itoa(writeField.OptimizeFieldIndex))
			} else {
				return codePrinter.AcceptField(writeField.Field.Type, writeField.Name, reader)
			}
		},
		"GetOutputDefTypeValue": func(filedType *meta.TableFieldType, collectionReadonly bool) string {
			typeDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, filedType, collectionReadonly)
			if err != nil {
				slog.Fatalf("gen cs code error:%v", typeDef)
			}
			return typeDef
		},
		"UniqueMapAssignment": func(writeContent *CSCodeWriteContent, originMapName string, recordName string, space int) string {
			var sb strings.Builder
			tableModel := writeContent.TableModel
			keyLen := len(tableModel.Meta.Keys)

			mapName := originMapName

			spaces := strings.Repeat("\t", space)

			if keyLen > 1 {

				for i := 0; i < keyLen-1; i++ {
					field := tableModel.Meta.Keys[i]
					lastMapName := mapName
					mapName = "_tmpuk" + strconv.Itoa(i)
					sb.WriteString(spaces)
					sb.WriteString(fmt.Sprintf("if (!%s.TryGetValue(%s.%s, out var %s))\n", lastMapName, recordName, field.Target, mapName))
					sb.WriteString(spaces)
					sb.WriteString("{\n")
					nextMapType := getKeyDefTypeMap(tableModel, writeContent.RecordClassName, i+1)
					sb.WriteString(spaces)
					sb.WriteString(fmt.Sprintf("\t%s = new %s();\n", mapName, nextMapType))
					sb.WriteString(spaces)
					sb.WriteString(fmt.Sprintf("\t%s[%s.%s] = %s;\n", lastMapName, recordName, field.Target, mapName))
					sb.WriteString(spaces)
					sb.WriteString("}\n")
				}
			}
			sb.WriteString(spaces)
			sb.WriteString(fmt.Sprintf("%s[%s.%s] = %s;", mapName, recordName, tableModel.Meta.Keys[keyLen-1].Target, recordName))
			return sb.String()
		},
		"UniqueMapGetFunc": func(writeContent *CSCodeWriteContent, originMapName string, space int) string {
			tableModel := writeContent.TableModel
			keyLen := len(tableModel.Meta.Keys)

			var sb strings.Builder
			spaces := strings.Repeat("\t", space)

			sb.WriteString(spaces)
			sb.WriteString(fmt.Sprintf("public %s GetDataById(", writeContent.RecordClassName))
			for i, field := range tableModel.Meta.Keys {
				keyName := "__k" + strconv.Itoa(i)
				if i > 0 {
					sb.WriteString(", ")
				}
				defType, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, field.Type, false)
				if err != nil {
					slog.Fatal(err)
				}
				sb.WriteString(fmt.Sprintf("%s %s", defType, keyName))
			}
			sb.WriteString(") { if (")

			mapName := originMapName
			for i := 0; i < keyLen; i++ {
				lastMapName := mapName
				mapName = "__tmpv" + strconv.Itoa(i)
				keyName := "__k" + strconv.Itoa(i)
				if i > 0 {
					sb.WriteString(" && ")
				}
				sb.WriteString(fmt.Sprintf("%s.TryGetValue(%s, out var %s)", lastMapName, keyName, mapName))
				if i == keyLen-1 {
					sb.WriteString(fmt.Sprintf(") { return %s; } return null; ", mapName))
				}
			}
			sb.WriteString("}")
			return sb.String()
		},
	}).Parse(templateCSCode)

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

func getKeyDefTypeMap(dataModel *model.TableModel, recordName string, offset int) string {
	keyDefType := dataModel.Meta.GetKeyDefTypeOffset(meta.FieldType_Int, offset)
	keyDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, keyDefType, false)
	if err != nil {
		slog.Fatal(err)
	}
	return util.ReplaceLast(keyDef, "int", recordName)
}

type CSCodeWriteContent struct {
	NameSpace          string
	CodeHead           string
	TableName          string
	RecordClassName    string
	TableClassName     string
	TableCommonName    string
	KeyDefTypeMap      string
	CollectionReadonly bool
	CSCodeWriteFields  []*CSCodeWriteField
	TableOptimize      *model.TableOptimize
	TableModel         *model.TableModel
}

type CSCodeWriteField struct {
	TypeDef            string
	Name               string
	Desc               string
	Field              *meta.TableField
	OptimizeFieldIndex int
}
