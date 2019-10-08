package cmd

import (
	"fmt"
	"github.com/millim/goploy/goflag"
	"github.com/millim/goploy/lib/util"
	"log"
	"os"
	"text/template"
)

//NilTemplate 无用内容
type NilTemplate struct {}

//Init 初始化基本数据
func Init() {
	createDir()
	createLocalFile()
}

func createDir(){
	if !util.FileExists(goflag.DefaultDir){
		fmt.Fprintf(os.Stdout, "%s: %s\n", util.FontColor("Create dir"), goflag.DefaultDir)
		os.MkdirAll(goflag.DefaultDir, os.ModePerm)
	}
}

func createLocalFile(){
	localFile := fmt.Sprintf("%s/local.json", goflag.DefaultDir)
	if !util.FileExists(localFile){
		fmt.Fprintf(os.Stdout, "%s: %s\n", util.FontColor("Create file"), localFile)
		lf, err := os.OpenFile(localFile, os.O_RDWR|os.O_CREATE, 0644)
		defer lf.Close()
		if err != nil {
			log.Println(err)
			return
		}

		t := template.New("local.json")
		t, _ = t.Parse(localTemplate())
		t.Execute(lf, NilTemplate{})
	}
}

func localTemplate() string{
	return `{
	"privateKey": "",
	"localMainFile": "",
	"scriptFile":""
}`
}




