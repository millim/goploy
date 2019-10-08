package deploy

import (
	"encoding/json"
	"fmt"
	"github.com/millim/goploy/goflag"
	"github.com/millim/goploy/lib/util"
	"log"
	"os"
)

func loadConfigFile(dirName string) {
	serverConfigPath := fmt.Sprintf("%s/%s/config.json", goflag.DefaultDir, dirName)
	localConfigPath := fmt.Sprintf("%s/local.json", goflag.DefaultDir)

	if !util.FileExists(serverConfigPath) || !util.FileExists(localConfigPath){
		log.Println("not find config dir")
		os.Exit(0)
		return
	}

	buildConfig(serverConfigPath, &serverConfig)
	buildConfig(localConfigPath, &localConfig)
}

func buildConfig(path string, s interface{}) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		panic(fmt.Sprintf("config file open error --->%s", err))
	}

	err = json.NewDecoder(file).Decode(s)
	if err != nil {
		panic(fmt.Sprintf("config format error --->%s", err))
	}
}
