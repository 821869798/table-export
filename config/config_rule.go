package config

type RawMetaRule struct {
	RuleName  string                `toml:"rule_name"`
	ConfigDir string                `toml:"config_dir"`
	Json      *RawMetaRuleUnitJson  `toml:"json"`
	Lua       *RawMetaRuleUnitLua   `toml:"lua"`
	CSBin     *RawMetaRuleUnitCSBin `toml:"cs_bin"`
	RuleUnits []MetaRuleUnit
}

func (r *RawMetaRule) initMetaRule() {
	if r.Json != nil {
		r.RuleUnits = append(r.RuleUnits, r.Json)
	}
	if r.Lua != nil {
		r.RuleUnits = append(r.RuleUnits, r.Lua)
	}
	if r.CSBin != nil {
		r.RuleUnits = append(r.RuleUnits, r.CSBin)
	}
}

type RawMetaRuleUnitJson struct {
	JsonOutputDir string `toml:"json_out"`
}

func (r *RawMetaRuleUnitJson) RuleExportType() ExportType {
	return ExportType_Json
}

type RawMetaRuleUnitLua struct {
	LuaOutputDir   string `toml:"lua_out"`
	EnableProcess  bool   `toml:"post_process"` //允许后处理
	TempDir        string `toml:"temp_dir"`
	PostProcessLua string `toml:"post_process_lua"`
	PostWorkDir    string `toml:"post_work_dir"`
	LuaWinDir      string `toml:"lua_win_dir"`
	LuaMacDir      string `toml:"lua_mac_dir"`
}

func (r *RawMetaRuleUnitLua) RuleExportType() ExportType {
	return ExportType_Lua
}
