package deploy

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"os/exec"
	"strings"
)

var session *ssh.Session
var client *ssh.Client
var serverConfig ServerConfig
var localConfig LocalConfig

//Deploy 部署文件
func Deploy() {
	setting()
	deleteCmd := fmt.Sprintf("rm %s", mainFilePath())

	err := execCmd(deleteCmd)
	if err != nil {
		fmt.Println("rm error ->", err)
	}

	var env []string
	for _, v := range serverConfig.Env {
		env = append(env, fmt.Sprintf("export %s", v))
	}

	err = execCmd(fmt.Sprintf(`echo -e "#! /bin/bash\n%s\n%s" > %s`, strings.Join(env, "\n"), runCmd(), localConfig.ScriptFile))
	if err != nil {
		fmt.Println("copy scripts error ->", err)
	}
	log.Print(fmt.Sprintf("%s %s %s","scp", localConfig.LocalMainFile, fmt.Sprintf("%s@%s:%s", serverConfig.User, serverConfig.SSHHost, mainFilePath())))
	cmd := exec.Command("scp", localConfig.LocalMainFile, fmt.Sprintf("%s@%s:%s", serverConfig.User, serverConfig.SSHHost, mainFilePath()))
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd exec failed: %s", fmt.Sprint(err))
	} else {
		chmodCmd := fmt.Sprintf("chmod 777 %s", mainFilePath())
		session2, _ := client.NewSession()
		err := session2.Run(chmodCmd)
		defer session2.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("done.")
		fmt.Println("copy file to remote server finished!")
	}
}



