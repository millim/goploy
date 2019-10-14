package main

import (
	"github.com/millim/goploy/cmd"
	"github.com/millim/goploy/cmd/deploy"
	"github.com/millim/goploy/goflag"
	"github.com/millim/goploy/lib/util"
	"github.com/millim/goploy/monit"
	"log"
)

func main(){

	args := goflag.Get()

	if len(args) == 1 && args[0] == "init"{
		cmd.Init()
		return
	}
	if len(args) < 2 {
		log.Fatal("install need dir name")
		return
	}

	if util.ArrayExists(&cmd.FirstArg, args[1]) {
		cm := args[1]
		switch cm {
		case "init":
			cmd.Init()
			cmd.Create(args[0])
			break
		case "install":
			deploy.Install()
			break
		case "deploy":
			deploy.Deploy()
			break
		case "start":
			deploy.Start()
			break
		case "stop":
			deploy.Stop()
			break
		case "restart":
			deploy.Restart()
			break
		case "monit-install":
			monit.InstallMonitTo()
			break
		case "monit-config":
			monit.SetConfig()
			break
		}
	}

}