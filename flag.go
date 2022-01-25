package main

import (
	"flag"
	"fmt"
	"os"
)

type parameters struct {
	//生成模式
	mode string
	//额外参数
	extra string
	//source模板生成
	genSource string
	//自定义配置路径
	confFile string
	//使用帮助
	help bool
	//版本信息
	version bool
}

var params *parameters

func init() {

	params = new(parameters)
	flag.StringVar(&params.mode, "m", "", `export data by source config <lua,json,cs_proto>, example [ -m=lua ] or [ -m=lua,json ] `)
	flag.StringVar(&params.extra, "e", "", `export data extra arg example <arg1=value1|arg2=value2>`)
	flag.StringVar(&params.genSource, "g", "", `generator source template <target_name,excel_name,sheet_name>, example [ -g=base_test,base_test,Sheet1 ]`)
	flag.BoolVar(&params.help, "help", false, "this help")
	flag.BoolVar(&params.version, "version", false, "this version")
	flag.StringVar(&params.confFile, "conf", "", "use custom config.toml filepath,default is conf/config.toml")
}

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: table-export [-m=<mode>] [-f=<target_name,excel_name,sheet_name>]
`)
	flag.PrintDefaults()
}
