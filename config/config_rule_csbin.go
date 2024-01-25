package config

type RawMetaRuleUnitCSBin struct {
	CodeTempDir       string   `toml:"code_temp_dir"`
	DataTempDir       string   `toml:"data_temp_dir"`
	GenCodeDir        string   `toml:"gen_code_dir"`
	DataBinDir        string   `toml:"data_bin_dir"`
	GenCodeNamespace  string   `toml:"code_namespace"`
	GenCodeHead       string   `toml:"gen_code_head"`
	CodeNotFoundKey   string   `toml:"code_not_found_key"`
	GenOptimizeData   bool     `toml:"gen_optimize"`
	EnumFiles         []string `toml:"enum_files"`
	EnumDefinePrefix  string   `toml:"enum_define_prefix"`
	ClassDefinePrefix string   `toml:"class_define_prefix"`
	BuiltinFieldTypes []string `toml:"builtin_field_types"`
	ScriptFieldTypes  []string `toml:"script_field_types"`
	PostGlobalScript  string   `toml:"post_global_script"`
}

func (r *RawMetaRuleUnitCSBin) RuleExportType() ExportType {
	return ExportType_CS_Bin
}

func (r *RawMetaRuleUnitCSBin) ActiveOptimizeData() bool {
	return r.GenOptimizeData
}

func (r *RawMetaRuleUnitCSBin) GetEnumFiles() []string {
	return r.EnumFiles
}

func (r *RawMetaRuleUnitCSBin) GetEnumDefinePrefix() string {
	return r.EnumDefinePrefix
}

func (r *RawMetaRuleUnitCSBin) GetClassDefinePrefix() string {
	return r.ClassDefinePrefix
}

func (r *RawMetaRuleUnitCSBin) GetBuiltinFieldTypes() []string {
	return r.BuiltinFieldTypes
}

func (r *RawMetaRuleUnitCSBin) GetExtFieldTypeScriptPath() []string {
	return r.ScriptFieldTypes
}

func (r *RawMetaRuleUnitCSBin) GetPostGlobalScriptPath() string {
	return r.PostGlobalScript
}
