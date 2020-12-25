package meta

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	log "github.com/sirupsen/logrus"
	"path/filepath"
	"strconv"
	"strings"
	"table-export/config"
	"table-export/constant"
	"table-export/util"
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

	if len(sourceSlice) != 3 {
		log.WithFields(log.Fields{
			"GenSrouce": g.genSource,
		}).Fatal("generator source arg error!")
	}

	targetName := sourceSlice[0]
	srcFileName := sourceSlice[1]
	sheetName := sourceSlice[2]

	fileName := srcFileName + constant.ExcelFileSuffix
	filePath := filepath.Join(config.GlobalConfig.Table.SrcDir, fileName)

	if !util.ExistFile(filePath) {
		log.WithFields(log.Fields{
			"FilePath": filePath,
		}).Fatal("generator source source file path not exist!")
	}

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	if len(rows) < config.GlobalConfig.Table.DataStart {
		log.Fatal("excel source row count must >= " + strconv.Itoa(config.GlobalConfig.Table.DataStart))
	}

	rtm := NewRawTableMeta()
	rtm.Target = targetName
	rtm.Mode = ""
	rtm.ExcelSources = []*RawTableSource{
		&RawTableSource{
			Table: fileName,
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
			log.WithFields(log.Fields{
				"Name:": cellStr,
			}).Fatal("excel source field name repeated!")
		}
		fieldSet[cellStr] = true
		descCellStr := descCols[index]
		rtf := NewRawTableField(cellStr, descCellStr)
		rtm.Fields = append(rtm.Fields, rtf)
	}

	genFilePath := filepath.Join(config.GlobalConfig.Meta.GenDir, targetName+constant.MetaFileSuffix)
	err = rtm.SaveTableMetaTemplateByDir(genFilePath)
	if err != nil {
		log.Fatal(err)
	}
}
