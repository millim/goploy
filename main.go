package main

import (
	"flag"
	"fmt"
	"github.com/millim/goploy/cmd"
	"github.com/millim/goploy/lib/util"
	"github.com/millim/goploy/version"
	"os"
)



func usage() {
	fmt.Fprintf(os.Stderr, "This version %s\n", version.VERSION)
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tgoploy init #current directory create local file\n")
	fmt.Fprintf(os.Stderr, "\tgoploy create DIRNAME #create deploy config file\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}



func main(){

	//template.CreateFile()
	//return

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		usage()
		return
	}
	var args []string
	args = flag.Args()

	if util.ArrayExists(&cmd.FirstArg, args[0]) {
		cm := args[0]
		switch cm {
		case "init":
			cmd.Init()
			break
		case "create":
			if len(args) == 2 {
				cmd.Create(args[1])
			}else{
				panic("create command error!")
			}
			break
		}
	}

	//if len(args) == 1 {
	//	logrus.Print("goploy version is ", version.VERSION)
	//	return
	//}
	//
	//m := args[1]
	//if !util.ArrayExists(&cmd.FirstArg, m) {
	//	logrus.Error("error command")
	//}

}