package config

import (
	"encoding/json"
	"fmt"
	"github.com/millim/goploy/goflag"
	"github.com/millim/goploy/lib/util"
	"log"
	"os"
)


var serverConfig *ServerConfig
var localConfig *LocalConfig



//LoadConfigFile 读取配置
func LoadConfigFile() (*ServerConfig, *LocalConfig){
	args := goflag.Get()
	if len(args) < 2{
		return nil, nil
	}
	dirName := args[0]
	if serverConfig == nil || localConfig == nil {
		serverConfigPath := fmt.Sprintf("%s/%s/config.json", goflag.DefaultDir, dirName)
		localConfigPath := fmt.Sprintf("%s/local.json", goflag.DefaultDir)

		if !util.FileExists(serverConfigPath) || !util.FileExists(localConfigPath){
			log.Println("not find config dir")
			os.Exit(0)
			return nil, nil
		}
		serverConfig = &ServerConfig{}
		localConfig = &LocalConfig{}
		buildConfig(serverConfigPath, serverConfig)
		buildConfig(localConfigPath, localConfig)

	}
	return serverConfig, localConfig
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
