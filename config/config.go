package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"table-export/define"
)

type RawGlobalConfig struct {
	Table *RawExcelTableConfig `toml:"table"`
	Meta  *RawMetaConfig       `toml:"meta"`
}

type RawExcelTableConfig struct {
	SrcDir     string `toml:"src_dir"`
	Name       int    `toml:"name"`
	Desc       int    `toml:"desc"`
	DataStart  int    `toml:"data_start"`
	ArraySplit string `toml:"array_split"`
	MapSplit1  string `toml:"map_split1"`
	MapSplit2  string `toml:"map_split2"`
}

type RawMetaConfig struct {
	GenDir      string              `toml:"gen_dir"`
	RuleLua     *RawMetaRuleLua     `toml:"rule_lua"`
	RuleJson    *RawMetaRuleJson    `toml:"rule_json"`
	RuleCSProto *RawMetaRuleCSProto `toml:"rule_cs_proto"`
}

func (m *RawMetaConfig) GetRawMetaBaseConfig(exportType define.ExportType) *RawMetaRuleBase {
	switch exportType {
	case define.ExportType_Lua:
		return m.RuleLua.RawMetaRuleBase
	case define.ExportType_Json:
		return m.RuleJson.RawMetaRuleBase
	case define.ExportType_CS_Proto:
		return m.RuleCSProto.RawMetaRuleBase
	}
	return nil
}

type RawMetaRuleBase struct {
	ConfigDir string `toml:"config_dir"`
}

type RawMetaRuleLua struct {
	*RawMetaRuleBase
	TempDir string `toml:"temp_dir"`
}

type RawMetaRuleJson struct {
	*RawMetaRuleBase
	OutputDir string `toml:"output_dir"`
}

type RawMetaRuleCSProto struct {
	*RawMetaRuleBase
	ProtoDir string `toml:"proto_dir"`
	BytesDir string `toml:"bytes_dir"`
}

var GlobalConfig *RawGlobalConfig

func ParseConfig(configFile string) {
	initPath(configFile)
	GlobalConfig = new(RawGlobalConfig)
	if _, err := toml.DecodeFile(GPath.Config(), GlobalConfig); err != nil {
		log.Fatalf("load global config error:%v", err)
	}
	log.Debug("load global config success!")

	//m3, err3 := json.Marshal(GlobalConfig)
	//if err3 != nil {
	//	panic(err3)
	//}
	//log.Println(string(m3))
}
