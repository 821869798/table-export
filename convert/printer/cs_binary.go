package printer

import (
	"fmt"
	"github.com/821869798/table-export/config"
	"github.com/821869798/table-export/convert/apiconvert"
	"github.com/821869798/table-export/convert/wrap"
	"github.com/821869798/table-export/meta"
	"github.com/gookit/slog"
	"strings"
)

type CSBinaryPrint struct {
	//数据集是否声明为只读的
	CollectionReadonly bool
}

func NewCSBinaryPrint(collectionReadonly bool) apiconvert.ICodePrinter {
	p := &CSBinaryPrint{
		CollectionReadonly: collectionReadonly,
	}
	return p
}

func (c *CSBinaryPrint) AcceptField(fieldType *meta.TableFieldType, fieldName string, reader string) string {
	return wrap.GetCodePrintValue(c, fieldType, fieldName, reader, 0)
}

func (c *CSBinaryPrint) AcceptOptimizeAssignment(fieldName string, reader string, commonDataName string) string {
	return fmt.Sprintf("{ int dataIndex = %s.ReadInt() - 1; %s = %s[dataIndex]; }", reader, fieldName, commonDataName)
}

func (c *CSBinaryPrint) AcceptInt(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return fmt.Sprintf("%s = %s.ReadInt();", fieldName, reader)
}

func (c *CSBinaryPrint) AcceptUInt(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return fmt.Sprintf("%s = %s.ReadUint();", fieldName, reader)
}

func (c *CSBinaryPrint) AcceptLong(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return fmt.Sprintf("%s = %s.ReadLong();", fieldName, reader)
}

func (c *CSBinaryPrint) AcceptULong(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return fmt.Sprintf("%s = %s.ReadUlong();", fieldName, reader)
}

func (c *CSBinaryPrint) AcceptFloat(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return fmt.Sprintf("%s = %s.ReadFloat();", fieldName, reader)
}

func (c *CSBinaryPrint) AcceptDouble(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return fmt.Sprintf("%s = %s.ReadDouble();", fieldName, reader)
}

func (c *CSBinaryPrint) AcceptBool(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return fmt.Sprintf("%s = %s.ReadBool();", fieldName, reader)
}

func (c *CSBinaryPrint) AcceptString(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	return fmt.Sprintf("%s = %s.ReadString();", fieldName, reader)
}

func (c *CSBinaryPrint) AcceptArray(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	collectionReadonly := c.CollectionReadonly
	valueDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, fieldType, false)
	if err != nil {
		slog.Fatal(err)
	}
	_n := fmt.Sprintf("__n%d", depth)
	_i := fmt.Sprintf("__i%d", depth)
	_v := fmt.Sprintf("__v%d", depth)
	_vDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, fieldType.Value, false)
	if err != nil {
		slog.Fatal(err)
	}

	// 添加数组大小
	index := strings.Index(valueDef, "[") // 获取第一个 '[' 的索引位置
	valueDefInit := ""
	if index != -1 {
		valueDefInit = valueDef[:index+1] + _n + valueDef[index+1:]
	} else {
		slog.Fatal("Accept Array type error:" + valueDef)
	}

	assignment := wrap.GetCodePrintValue(c, fieldType.Value, _v, reader, depth+1)

	if collectionReadonly {
		_f := fmt.Sprintf("__f%d", depth)
		return fmt.Sprintf("{int %s = %s.ReadSize(); var %s = new %s; %s = %s; for(var %s = 0 ; %s < %s ; %s++ ){ %s %s; %s %s[%s] = %s; } }", _n, reader, _f, valueDefInit, fieldName, _f, _i, _i, _n, _i, _vDef, _v, assignment, _f, _i, _v)
	} else {
		return fmt.Sprintf("{int %s = %s.ReadSize(); %s = new %s; for(var %s = 0 ; %s < %s ; %s++ ){ %s %s; %s %s[%s] = %s; } }", _n, reader, fieldName, valueDefInit, _i, _i, _n, _i, _vDef, _v, assignment, fieldName, _i, _v)
	}
}

func (c *CSBinaryPrint) AcceptMap(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string {
	collectionReadonly := c.CollectionReadonly

	keyDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, fieldType.Key, collectionReadonly)
	if err != nil {
		slog.Fatal(err)
	}
	valueDef, err := wrap.GetOutputDefTypeValue(config.ExportType_CS_Bin, fieldType.Value, collectionReadonly)
	if err != nil {
		slog.Fatal(err)
	}

	_n := fmt.Sprintf("__n%d", depth)
	_i := fmt.Sprintf("__i%d", depth)
	_k := fmt.Sprintf("__k%d", depth)
	_v := fmt.Sprintf("__v%d", depth)
	_f := fmt.Sprintf("__f%d", depth)

	mapDef := fmt.Sprintf("Dictionary<%s, %s>", keyDef, valueDef)

	keyAssignment := wrap.GetCodePrintValue(c, fieldType.Key, _k, reader, depth+1)
	valueAssignment := wrap.GetCodePrintValue(c, fieldType.Value, _v, reader, depth+1)

	if collectionReadonly {
		return fmt.Sprintf("{ int %s = %s.ReadSize(); var %s = new %s (%s * 3 / 2) ; %s = %s; for(var %s = 0 ; %s < %s ; %s++ ) {%s %s; %s %s %s; %s %s.Add(%s, %s); } }", _n, reader, _f, mapDef, _n, fieldName, _f, _i, _i, _n, _i, keyDef, _k, keyAssignment, valueDef, _v, valueAssignment, _f, _k, _v)
	} else {
		return fmt.Sprintf("{ int %s = %s.ReadSize(); %s = new %s (%s * 3 / 2); for(var %s = 0 ; %s < %s ; %s++ ) {%s %s; %s %s %s; %s %s.Add(%s, %s); } }", _n, reader, fieldName, mapDef, _n, _i, _i, _n, _i, keyDef, _k, keyAssignment, valueDef, _v, valueAssignment, fieldName, _k, _v)
	}
}
