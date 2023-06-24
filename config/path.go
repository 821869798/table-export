package config

import (
	"github.com/gookit/slog"
	"os"
	"path/filepath"
)

func ExeDir() string {
	return gpath.exeDir
}

func AbsExeDir(paths ...string) string {
	paths = append([]string{ExeDir()}, paths...)
	return filepath.Join(paths...)
}

func ConfigDir() string {
	return gpath.configFile
}

type GlobalPath struct {
	exeDir     string
	configFile string
}

var gpath *GlobalPath

func initPath(configFile string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		slog.Fatalf("init const path error:%v", err)
	}
	if configFile == "" {
		//如果参数为空，使用默认路径
		configFile = filepath.Join(dir, "conf/config.toml")
	}
	gpath = &GlobalPath{
		exeDir:     dir,
		configFile: configFile,
	}
}
