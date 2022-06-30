package cs_proto

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoprint"
	"os"
	"table-export/config"
	"table-export/data/model"
	"table-export/meta"
	"table-export/util"
)

func buildProtoFile(dataModel *model.TableModel, ruleCSProto *config.RawMetaRuleCSProto) (*desc.FileDescriptor, error) {
	tableMeta := dataModel.Meta
	protoBuilder := builder.NewFile(tableMeta.Target).SetPackageName(ruleCSProto.ProtoPackage).SetProto3(true)
	complexMap := make(map[string]interface{})
	outputName := util.FormatPascalString(tableMeta.Target)
	//创建cfg
	recordMsg := builder.NewMessage(outputName + "Cfg")
	for _, tf := range tableMeta.Fields {
		fieldBuilder := getProtoFieldBuilder(tf.Target, tf.Type, complexMap)
		if fieldBuilder == nil {
			return nil, errors.New(fmt.Sprintf("export .proto define file field[%v] error", tf.Target))
		}
		recordMsg.AddField(fieldBuilder)
	}

	rootMsg := builder.NewMessage(outputName + "CfgTable")
	recordField := builder.NewField("data", builder.FieldTypeMessage(recordMsg))
	recordField.SetRepeated()
	rootMsg.AddField(recordField)

	protoBuilder.AddMessage(recordMsg)
	protoBuilder.AddMessage(rootMsg)

	pfd, err := protoBuilder.Build()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("export .proto define file build error:%v", err))
	}
	pr := &protoprint.Printer{}
	var buf bytes.Buffer
	err = pr.PrintProtoFile(pfd, &buf)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("export .proto define file print error:%v", err))
	}

	filePath := config.AbsExeDir(ruleCSProto.ProtoTempDir, getOutputProtoFileName(tableMeta.Target))
	file, err := os.Create(filePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("export .proto define file create file error:%v", err))
	}
	_, err = file.Write(buf.Bytes())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("export .proto define file write error:%v", err))
	}

	return pfd, nil
}

func getOutputProtoFileName(target string) string {
	outputName := util.FormatPascalString(target)
	return outputName + "CfgTable.proto"
}

func getProtoFieldBuilder(fieldName string, tft *meta.TableFiledType, complexMap map[string]interface{}) *builder.FieldBuilder {
	var field *builder.FieldBuilder = nil
	if tft.IsBaseType() {
		ft := getProtoFieldType(tft, complexMap)
		field = builder.NewField(fieldName, ft)
	} else {
		switch tft.Type {
		case meta.FieldType_Slice:
			ft := getProtoFieldType(tft.Value, complexMap)
			if ft == nil {
				return nil
			}
			field = builder.NewField(fieldName, ft)
			field.SetRepeated()
		case meta.FieldType_Map:
			keyType, ok := tft.GetKeyFieldType()
			if !ok {
				return nil
			}
			if !tft.Value.IsBaseType() {
				return nil
			}
			kft := getProtoFieldType(keyType, complexMap)
			vft := getProtoFieldType(tft.Value, complexMap)
			if kft == nil || vft == nil {
				return nil
			}
			field = builder.NewMapField(fieldName, kft, vft)
		default:
			//复合类型
			if complexType, ok := complexMap[fieldName]; ok {
				var ft *builder.FieldType
				switch complexType.(type) {
				case *builder.MessageBuilder:
					ft = builder.FieldTypeMessage(complexType.(*builder.MessageBuilder))
				case *builder.EnumBuilder:
					ft = builder.FieldTypeEnum(complexType.(*builder.EnumBuilder))
				}
				if ft != nil {
					field = builder.NewField(fieldName, ft)
				}
			}
		}
	}
	return field
}

func getProtoFieldType(tft *meta.TableFiledType, complexMap map[string]interface{}) *builder.FieldType {
	var ft *builder.FieldType = nil
	switch tft.Type {
	case meta.FieldType_Int:
		ft = builder.FieldTypeInt32()
	case meta.FieldType_UInt:
		ft = builder.FieldTypeUInt32()
	case meta.FieldType_Long:
		ft = builder.FieldTypeInt64()
	case meta.FieldType_ULong:
		ft = builder.FieldTypeUInt64()
	case meta.FieldType_Bool:
		ft = builder.FieldTypeBool()
	case meta.FiledType_Float:
		ft = builder.FieldTypeFloat()
	case meta.FiledType_Double:
		ft = builder.FieldTypeDouble()
	case meta.FieldType_String:
		ft = builder.FieldTypeString()
	default:
	}
	return ft
}
