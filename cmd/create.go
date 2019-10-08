package cmd

import (
	"fmt"
	"github.com/millim/goploy/goflag"
	"github.com/millim/goploy/lib/util"
	"log"
	"os"
	"text/template"
)

//Create <----
func Create(dirName string) {
	createConfigDir(dirName)
	createConfigFile(dirName)
}

func createConfigDir(dirName string) {
	fullPath := fmt.Sprintf("%s/%s", goflag.DefaultDir, dirName)
	if !util.FileExists(fullPath) {
		fmt.Fprintf(os.Stdout, "%s: %s\n", util.FontColor("Create dir"), fullPath)
		os.MkdirAll(fullPath, os.ModePerm)
	}
}

func createConfigFile(dirName string) {
	fullPath := fmt.Sprintf("%s/%s/config.json", goflag.DefaultDir, dirName)
	if !util.FileExists(fullPath) {
		fmt.Fprintf(os.Stdout, "%s: %s\n", util.FontColor("Create file"), fullPath)
		lf, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0644)
		defer lf.Close()
		if err != nil {
			log.Println(err)
			return
		}
		t := template.New("config.json")
		t, _ = t.Parse(configTemplate())
		t.Execute(lf, NilTemplate{})
	} else {
		fmt.Println("Exist`s file")
	}
}

func configTemplate() string {
	return `{
	"user":"",
	"sshHost": "",
	"sshPort": "",
	"serverDir": "",
	"serverFile": "",
	"needSudo": false,
	"pidFile": "",
	"env": ["GIN_MODE=release"],
	"command": ["--port=6061"]
}`
}
