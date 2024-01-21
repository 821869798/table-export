package field_type

type IExtFieldType interface {
	// Name 该类型的名字
	Name() string
	// DefineFile 定义的文件
	DefineFile() string
	// TableFieldType 字段的类型
	TableFieldType() *TableFieldType
	// ParseDataOne 通过单个字段string解析数据
	ParseDataOne(origin string) (interface{}, error)
	// ParseDataMultiple 通过多个字段string解析数据
	ParseDataMultiple(originArray []string) (interface{}, error)
}
