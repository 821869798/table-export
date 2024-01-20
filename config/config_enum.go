package config

type RawMetaEnumConfig struct {
	EnumFileName string               `toml:"enum_file"`
	EnumDefines  []*RawMetaEnumDefine `toml:"enums"`
}

type RawMetaEnumDefine struct {
	Name         string              `toml:"name"`         // 枚举的名字
	Parse4String bool                `toml:"parse_string"` // 是否从string数据解析，否则为int
	Values       []*RawMetaEnumValue `toml:"values"`
}

type RawMetaEnumValue struct {
	Index       int32  `tom:"index"`
	Name        string `toml:"name"`
	Desc        string `toml:"desc"`
	ValueString string `toml:"string"`
}
