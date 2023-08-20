package meta

import (
	"encoding/csv"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/consts"
	"github.com/821869798/table-export/util"
	"github.com/gookit/slog"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"strings"
)

type GenMeta struct {
	genSource string
}

func NewGenMeta(genSource string) *GenMeta {
	g := &GenMeta{
		genSource: genSource,
	}
	return g
}

func (g *GenMeta) Run() {

	sourceSlice := strings.Split(g.genSource, ",")

	inputOk := false
	isCsv := false
	if len(sourceSlice) == 2 && strings.HasSuffix(sourceSlice[1], ".csv") {
		inputOk = true
		isCsv = true
	}
	if len(sourceSlice) == 3 && !strings.HasSuffix(sourceSlice[1], ".csv") {
		inputOk = true
	}

	if !inputOk {
		slog.Fatalf("generator source arg error! GenSource:%s", g.genSource)
	}

	targetName := sourceSlice[0]
	srcFileName := sourceSlice[1]

	filePath := util.RelExecuteDir(config.GlobalConfig.Table.SrcDir, srcFileName)

	if !util.ExistFile(filePath) {
		slog.Fatalf("generator source source file path not exist! FilePath:%s", filePath)
	}

	sheetName := ""
	var rows [][]string
	var err error
	if !isCsv {
		sheetName = sourceSlice[2]
		rows, err = readExcelFile(filePath, sheetName)
	} else {
		rows, err = readCsvFile(filePath)
	}

	if err != nil {
		slog.Fatal(err)
	}

	if len(rows) < config.GlobalConfig.Table.DataStart {
		slog.Fatal("excel source row count must >= " + strconv.Itoa(config.GlobalConfig.Table.DataStart))
	}

	rtm := NewRawTableMeta()
	rtm.Target = targetName
	rtm.Mode = ""
	if isCsv {
		rtm.SourceType = "csv"
	} else {
		rtm.SourceType = "excel"
	}

	rtm.Sources = []*RawTableSource{
		&RawTableSource{
			Table: srcFileName,
			Sheet: sheetName,
		},
	}
	rtm.Fields = make([]*RawTableField, 0)
	fieldSet := make(map[string]bool)
	nameCols := rows[config.GlobalConfig.Table.Name]
	descCols := rows[config.GlobalConfig.Table.Desc]
	for index, cellStr := range nameCols {
		if cellStr == "" {
			continue
		}
		//判断是否重复
		_, ok := fieldSet[cellStr]
		if ok {
			slog.Fatalf("excel source field name[%s] repeated!", cellStr)
		}
		fieldSet[cellStr] = true
		descCellStr := ""
		if index < len(descCols) {
			descCellStr = descCols[index]
		}
		rtf := NewRawTableField(cellStr, descCellStr)
		rtm.Fields = append(rtm.Fields, rtf)
	}

	genFilePath := util.RelExecuteDir(config.GlobalConfig.Meta.GenDir, targetName+consts.MetaFileSuffix)
	err = rtm.SaveTableMetaTemplateByDir(genFilePath)
	if err != nil {
		slog.Fatal(err)
	}
}

func readExcelFile(filePath string, sheetName string) ([][]string, error) {
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, err
	}

	return f.GetRows(sheetName)
}

func readCsvFile(filePath string) ([][]string, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(csvFile) //创建一个新的写入文件流
	csvReader.LazyQuotes = true
	rowData, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	//去除utf-8的bom
	if len(rowData) > 0 && len(rowData[0]) > 0 {
		firstData := rowData[0][0]
		rowData[0][0] = strings.TrimPrefix(firstData, "\uFEFF")
	}
	return rowData, nil
}
