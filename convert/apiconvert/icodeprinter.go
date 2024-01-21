package apiconvert

import (
	"github.com/821869798/table-export/field_type"
)

type ICodePrinter interface {
	AcceptField(fieldType *field_type.TableFieldType, fieldName string, reader string) string
	AcceptOptimizeAssignment(fieldName string, reader string, commonDataName string) string
	AcceptInt(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptUInt(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptLong(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptULong(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptFloat(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptDouble(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptBool(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptString(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptEnum(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptArray(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptMap(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
	AcceptClass(fieldType *field_type.TableFieldType, fieldName string, reader string, depth int32) string
}
