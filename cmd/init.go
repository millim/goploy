package cmd

import (
	"fmt"
	"github.com/millim/goploy/lib/util"
	"log"
	"os"
	"text/template"
)

//const dirPath = "/Users/millim/my_project/github.com/millim/goto/goploy/template"

//const dirPath = "./goploy/template"

const DefaultDir = "goploy"


type NilTemplate struct {}

func Init() {
	createDir()
	createLocalFile()
}



func createDir(){
	if !util.FileExists(DefaultDir){
		fmt.Fprintf(os.Stdout, "%s: %s\n", util.FontColor("Create dir"), DefaultDir)
		os.MkdirAll(DefaultDir, os.ModePerm)
	}
}

func createLocalFile(){
	localFile := fmt.Sprintf("%s/local.json", DefaultDir)
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




