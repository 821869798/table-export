package field_type

type IExtFieldType interface {
	// Name 该类型的名字
	Name() string
	// DefineFile 定义的文件
	DefineFile() string
	// TableFieldType 字段的类型
	TableFieldType() *TableFieldType
	// ParseOriginData 通过单个字段string解析数据
	ParseOriginData(origin string) (interface{}, error)
}
