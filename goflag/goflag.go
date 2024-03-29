package goflag

import (
	"flag"
	"fmt"
	"github.com/millim/goploy/version"
	"os"
)

//DefaultDir default local dir
const DefaultDir = "gop"

var args []string

//Get get args
func Get() []string {
	return args
}

func init() {
	flag.Usage = usage
	flag.Parse()
	args = flag.Args()
	if flag.NArg() == 0 {
		usage()
		os.Exit(0)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "当前版本 %s\n", version.VERSION)
	fmt.Fprintf(os.Stderr, " %s使用方式:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgoploy DIRNAME init #初始化指定目录的配置\n\n")
	fmt.Fprintf(os.Stderr, "使用目录中的配置文件对服务器进行操作:\n")
	fmt.Fprintf(os.Stderr, "\tgoploy DIRNAME install #创建目录\n")
	fmt.Fprintf(os.Stderr, "\tgoploy DIRNAME deploy #部署localMainFile指定的文件到服务器\n")
	fmt.Fprintf(os.Stderr, "\tgoploy DIRNAME start #执行go run xxx\n")
	fmt.Fprintf(os.Stderr, "\tgoploy DIRNAME stop #停止服务器（需要配置pidFile）\n")
	fmt.Fprintf(os.Stderr, "\tgoploy DIRNAME restart #重启服务器 （需要配置pidFile，且服务器启动为指定gracehttp）\n")
	flag.PrintDefaults()
}
