package config

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type GlobalPath struct {
	exeDir       string
	configFile   string
	templatePath string
}

func (g *GlobalPath) ExeDir() string {
	return g.exeDir
}

func (g *GlobalPath) AbsExeDir(paths ...string) string {
	paths = append([]string{g.ExeDir()}, paths...)
	return filepath.Join(paths...)
}

func (g *GlobalPath) Config() string {
	return g.configFile
}

func (g *GlobalPath) TemplatePath() string {
	return g.templatePath
}

var GPath *GlobalPath

func initPath(configFile string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal("init const path error:", err)
	}
	if configFile == "" {
		//如果参数为空，使用默认路径
		configFile = filepath.Join(dir, "conf/config.toml")
	}
	templatePath := filepath.Join(dir, "tpl")
	GPath = &GlobalPath{
		exeDir:       dir,
		configFile:   configFile,
		templatePath: templatePath,
	}
}
