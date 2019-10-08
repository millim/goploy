package main

import (
	"github.com/millim/goploy/cmd"
	"github.com/millim/goploy/cmd/deploy"
	"github.com/millim/goploy/goflag"
	"github.com/millim/goploy/lib/util"
	"log"
)

func main(){

	args := goflag.Get()

	if util.ArrayExists(&cmd.FirstArg, args[0]) {
		cm := args[0]
		switch cm {
		case "init":
			cmd.Init()
			if len(args) == 2 {
				cmd.Create(args[1])
			}
			break
		case "install":
			if len(args) == 2 {
				deploy.Install()
			}else{
				log.Fatal("install need dir name")
			}
			break
		case "deploy":
			if len(args) == 2 {
				deploy.Deploy()
			}else{
				log.Fatal("install need dir name")
			}
			break
		case "start":
			if len(args) == 2 {
				deploy.Start()
			}else{
				log.Fatal("install need dir name")
			}
			break
		case "stop":
			if len(args) == 2 {
				deploy.Stop()
			}else{
				log.Fatal("install need dir name")
			}
			break
		case "restart":
			if len(args) == 2 {
				deploy.Restart()
			}else{
				log.Fatal("install need dir name")
			}
			break
		}
	}

}