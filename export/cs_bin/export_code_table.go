package cs_bin

import (
	"fmt"
	"github.com/821869798/fankit/fanstr"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/printer"
	"github.com/821869798/table-export/convert/wrap"
	"github.com/821869798/table-export/data/model"
	"github.com/821869798/table-export/field_type"
	"github.com/821869798/table-export/meta"
	"github.com/gookit/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

func GenCSBinCodeTable(dataModel *model.TableModel, csBinRule *config.RawMetaRuleUnitCSBin, outputPath string) {
	//集合是否用只读的类型
	const collectionReadonly = false

	templateRoot := &CSCodeWriteTableFile{
		NameSpace:          fanstr.ReplaceWindowsLineEnd(csBinRule.GenCodeNamespace),
		CodeHead:           fanstr.ReplaceWindowsLineEnd(csBinRule.GenCodeHead),
		CodeNotFoundKey:    fanstr.ReplaceWindowsLineEnd(csBinRule.CodeNotFoundKey),
		TableName:          dataModel.Meta.Target,
		RecordClassName:    "Cfg" + dataModel.Meta.Target,
		TableClassName:     "Tbl" + dataModel.Meta.Target,
		TableCommonName:    "_TbCommon" + dataModel.Meta.Target,
		CollectionReadonly: collectionReadonly,
		CSCodeWriteFields:  make([]*CSCodeWriteTableField, 0, len(dataModel.Meta.Fields)),
		CSCodeExtraFields:  make([]*CSCodeWriteTableField, 0, len(dataModel.Meta.ExtraFields)),
		TableOptimize:      dataModel.Optimize,
		TableModel:         dataModel,
	}

	templateRoot.KeyDefTypeMap = getKeyDefTypeMap(dataModel, templateRoot.RecordClassName, 0)

	// 写个每行数据的字段
	for _, field := range dataModel.Meta.Fields {
		typeDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, field.Type, collectionReadonly)
		if err != nil {
			slog.Fatalf("gen cs code error:%v", typeDef)
		}

		writeField := &CSCodeWriteTableField{
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

	// 写入table root下的额外字段
	for _, field := range dataModel.Meta.ExtraFields {
		typeDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, field.Type, collectionReadonly)
		if err != nil {
			slog.Fatalf("gen cs code error:%v", typeDef)
		}
		writeField := &CSCodeWriteTableField{
			TypeDef:            typeDef,
			Name:               field.Target,
			Desc:               field.Desc,
			Field:              field,
			OptimizeFieldIndex: -1,
		}
		templateRoot.CSCodeExtraFields = append(templateRoot.CSCodeExtraFields, writeField)
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
	tmpl, err := template.New("cs_bin_table").Funcs(template.FuncMap{
		"CSbinFieldReader": func(fieldType *field_type.TableFieldType, fieldName string, reader string) string {
			return codePrinter.AcceptField(fieldType, fieldName, reader)
		},
		"CSbinFieldReaderEx": func(writeField *CSCodeWriteTableField, reader string, commonDataName string) string {
			if writeField.OptimizeFieldIndex >= 0 {
				return codePrinter.AcceptOptimizeAssignment(writeField.Name, reader, commonDataName+strconv.Itoa(writeField.OptimizeFieldIndex))
			} else {
				return codePrinter.AcceptField(writeField.Field.Type, writeField.Name, reader)
			}
		},
		"GetOutputDefTypeValue": func(fieldType *field_type.TableFieldType, collectionReadonly bool) string {
			typeDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, fieldType, collectionReadonly)
			if err != nil {
				slog.Fatalf("gen cs code error:%v", typeDef)
			}
			return typeDef
		},
		"UniqueMapAssignment": func(writeContent *CSCodeWriteTableFile, originMapName string, recordName string, space int) string {
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
		"UniqueMapGetFunc": func(writeContent *CSCodeWriteTableFile, originMapName string, space int) string {
			return UniqueMapGetFunc(writeContent, originMapName, space, "Get", false)
		},

		"UniqueMapGetFuncWithoutError": func(writeContent *CSCodeWriteTableFile, originMapName string, space int) string {
			return UniqueMapGetFunc(writeContent, originMapName, space, "GetWithoutError", true)
		},
	}).Parse(templateCSCodeTable)

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
	keyDefType := dataModel.Meta.GetKeyDefTypeOffset(field_type.EFieldType_Int, offset)
	keyDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, keyDefType, false)
	if err != nil {
		slog.Fatal(err)
	}
	return fanstr.ReplaceLast(keyDef, "int", recordName)
}

func UniqueMapGetFunc(writeContent *CSCodeWriteTableFile, originMapName string, space int, funcName string, withoutError bool) string {
	tableModel := writeContent.TableModel
	keyLen := len(tableModel.Meta.Keys)

	var sb strings.Builder
	spaces := strings.Repeat("    ", space)

	sb.WriteString(spaces)
	sb.WriteString(fmt.Sprintf("public %s %s(", writeContent.RecordClassName, funcName))
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
	sb.WriteString(fmt.Sprintf(") \n%s{\n%s    if (", spaces, spaces))

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
			sb.WriteString(fmt.Sprintf(") { return %s; }", mapName))
			if !withoutError {
				// 添加找不到报错
				errorKeyString := ""
				for i := 0; i < keyLen; i++ {
					if i == keyLen-1 {
						errorKeyString += fmt.Sprintf("__k%d.ToString()", i)
					} else {
						errorKeyString += fmt.Sprintf("__k%d.ToString() + \" \" + ", i)
					}
				}
				sb.WriteString("\n" + spaces + "    ")
				errorString := fmt.Sprintf(writeContent.CodeNotFoundKey, writeContent.TableClassName, errorKeyString)
				sb.WriteString(strings.ReplaceAll(errorString, "\n", "\n"+spaces+"    "))
			}
			sb.WriteString(fmt.Sprintf("\n%s    return null; \n", spaces))
		}
	}
	sb.WriteString(spaces)
	sb.WriteString("}")
	return sb.String()
}

type CSCodeWriteTableFile struct {
	NameSpace          string
	CodeHead           string
	CodeNotFoundKey    string
	TableName          string
	RecordClassName    string
	TableClassName     string
	TableCommonName    string
	KeyDefTypeMap      string
	CollectionReadonly bool
	CSCodeWriteFields  []*CSCodeWriteTableField
	CSCodeExtraFields  []*CSCodeWriteTableField
	TableOptimize      *model.TableOptimize
	TableModel         *model.TableModel
}

type CSCodeWriteTableField struct {
	TypeDef            string
	Name               string
	Desc               string
	Field              *meta.TableField
	OptimizeFieldIndex int
}
