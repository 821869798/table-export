package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type RawGlobalConfig struct {
	Table  *RawTableConfig  `toml:"table"`
	Meta   *RawMetaConfig   `toml:"meta"`
	Plugin *RawPluginConfig `toml:"plugin"`
}

type RawMetaConfig struct {
	GenDir string                        `toml:"gen_dir"`
	Rule   map[string]*RawMetaRuleConfig `toml:"rule"`
}

type RawMetaRuleConfig struct {
	ExportType string `toml:"type"`
	ConfigDir  string `toml:"config_dir"`
}

type RawTableConfig struct {
	SrcDir     string `toml:"src_dir"`
	Name       int    `toml:"name"`
	Desc       int    `toml:"desc"`
	DataStart  int    `toml:"data_start"`
	ArraySplit string `toml:"array_split"`
	MapSplit1  string `toml:"map_split1"`
	MapSplit2  string `toml:"map_split2"`
}

type RawPluginConfig struct {
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
