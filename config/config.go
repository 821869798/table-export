package config

import (
	"github.com/821869798/table-export/util"
	"github.com/BurntSushi/toml"
	"github.com/gookit/slog"
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
	GenDir  string         `toml:"gen_dir"`
	Rules   []*RawMetaRule `toml:"rules"`
	RuleMap map[string]*RawMetaRule
}

func GetMetaRuleConfigByName(name string) *RawMetaRule {
	rule, ok := GlobalConfig.Meta.RuleMap[name]
	if ok {
		return rule
	}
	return nil
}

var GlobalConfig *RawGlobalConfig

func ParseConfig(configFile string) {
	configFile = util.AbsOrRelExecutePath(configFile)
	GlobalConfig = new(RawGlobalConfig)
	if _, err := toml.DecodeFile(configFile, GlobalConfig); err != nil {
		slog.Fatalf("load global config error:%v", err)
	}

	//初始化
	GlobalConfig.Meta.RuleMap = make(map[string]*RawMetaRule, 4)
	for _, rule := range GlobalConfig.Meta.Rules {
		rule.initMetaRule()
		GlobalConfig.Meta.RuleMap[rule.RuleName] = rule
	}

	slog.Debug("load global config success!")

}
