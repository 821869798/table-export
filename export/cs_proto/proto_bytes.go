package cs_proto

import (
	"fmt"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	log "github.com/sirupsen/logrus"
	"os"
	"table-export/config"
	"table-export/data/model"
	"table-export/define"
	"table-export/export/wrap"
	"table-export/util"
)

func buildProtoBytesFile(dataModel *model.TableModel, ruleCSProto *config.RawMetaRuleCSProto, pfd *desc.FileDescriptor) {
	outputName := util.FormatPascalString(dataModel.Meta.Target)
	//存储数据表的List结构的Message
	rootMsg := pfd.FindMessage(fmt.Sprintf("%s.%sCfgTable", ruleCSProto.ProtoPackage, outputName))
	rootDm := dynamic.NewMessage(rootMsg)
	//单条结构
	recordMsg := pfd.FindMessage(fmt.Sprintf("%s.%sCfg", ruleCSProto.ProtoPackage, outputName))

	rowDataOffset := config.GlobalConfig.Table.DataStart + 1
	for rowIndex, rowData := range dataModel.RawData {
		//数据表中一行数据的字符串
		recordDm := dynamic.NewMessage(recordMsg)

		for fid, tf := range dataModel.Meta.Fields {
			rawIndex := dataModel.NameIndexMapping[tf.Target]
			var rawStr string
			if rawIndex < len(rowData) {
				rawStr = rowData[rawIndex]
			}
			output, err := wrap.GetOutputValue(define.ExportType_CS_Proto, tf.Type, rawStr)
			if err != nil {
				log.Fatalf("export cs_proto target file[%v] RowCount[%v] filedName[%v] error:%v", dataModel.Meta.Target, rowIndex+rowDataOffset, tf.Source, err.Error())
			}
			recordDm.SetFieldByNumber(fid+1, output)
		}

		rootDm.AddRepeatedFieldByNumber(1, recordDm)
	}

	buf, err := rootDm.Marshal()

	if err != nil {
		log.Fatalf("export cs_proto target file[%v] error:%v", dataModel.Meta.Target, err)
	}

	filePath := config.AbsExeDir(ruleCSProto.BytesDir, dataModel.Meta.Target+".bytes")
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("export .proto bytes file create file error:%v", err)
	}
	_, err = file.Write(buf)
	if err != nil {
		log.Fatalf("export .proto bytes file write error:%v", err)
	}

}
