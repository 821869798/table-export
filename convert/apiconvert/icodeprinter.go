package apiconvert

import "github.com/821869798/table-export/meta"

type ICodePrinter interface {
	AcceptField(fieldType *meta.TableFieldType, fieldName string, reader string) string
	AcceptOptimizeAssignment(fieldName string, reader string, commonDataName string) string
	AcceptInt(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptUInt(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptLong(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptULong(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptFloat(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptDouble(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptBool(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptString(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptEnum(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptArray(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptMap(fieldType *meta.TableFieldType, fieldName string, reader string, depth int32) string
}
