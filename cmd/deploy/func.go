package deploy

import (
	"fmt"
	"github.com/millim/goploy/config"
	"strings"
)

func setting() {
	serverConfig, localConfig = config.LoadConfigFile()
}

func runCmd() string {
	//logParams := fmt.Sprintf(" --logDir=%s", serverConfig.ServerDir)
	return fmt.Sprintf("nohup %s %s >/dev/null 2>&1 &", mainFilePath(), strings.Join(serverConfig.Command, " "))
}

func mainFilePath() string {
	return fmt.Sprintf("%s/%s", serverConfig.ServerDir, serverConfig.ServerFile)
}
