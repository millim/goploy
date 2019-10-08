package deploy

import (
	"fmt"
	"github.com/millim/goploy/goflag"
	"log"
	"strings"
)

func setting() {
	args := goflag.Get()
	log.Println(args)
	if len(args) != 2{
		return
	}
	loadConfigFile(args[1])
	newSession()
}

func runCmd() string {
	//logParams := fmt.Sprintf(" --logDir=%s", serverConfig.ServerDir)
	return fmt.Sprintf("nohup %s %s >/dev/null 2>&1 &", mainFilePath(), strings.Join(serverConfig.Command, " "))
}

func mainFilePath() string {
	return fmt.Sprintf("%s/%s", serverConfig.ServerDir, serverConfig.ServerFile)
}
