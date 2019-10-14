package deploy

import (
	"fmt"
	"github.com/millim/goploy/config"
	"github.com/millim/goploy/session"
	"os/exec"
	"strings"
)

var serverConfig *config.ServerConfig
var localConfig *config.LocalConfig

//Deploy 部署文件
func Deploy() {
	setting()
	deleteCmd := fmt.Sprintf("rm %s", mainFilePath())

	err := session.ExecCmd(deleteCmd)
	if err != nil {
		fmt.Println("rm error ->", err)
	}

	var env []string
	for _, v := range serverConfig.Env {
		env = append(env, fmt.Sprintf("export %s", v))
	}

	err = session.ExecCmd(fmt.Sprintf(`echo -e "#! /bin/bash\n%s\n%s" > %s`, strings.Join(env, "\n"), runCmd(), localConfig.ScriptFile))
	if err != nil {
		fmt.Println("copy scripts error ->", err)
	}

	var cmd *exec.Cmd
	if serverConfig.SSHPort == "22" || serverConfig.SSHPort == ""{
		cmd = exec.Command("scp", localConfig.LocalMainFile, fmt.Sprintf("%s@%s:%s", serverConfig.User, serverConfig.SSHHost, mainFilePath()))
	}else{
		cmd = exec.Command("scp","-P", serverConfig.SSHPort, localConfig.LocalMainFile, fmt.Sprintf("%s@%s:%s", serverConfig.User, serverConfig.SSHHost, mainFilePath()))
	}
	if err := cmd.Run(); err != nil {
		fmt.Printf("cmd exec failed: %s", fmt.Sprint(err))
	} else {
		chmodCmd := fmt.Sprintf("chmod 777 %s", mainFilePath())
		session.ExecCmd(chmodCmd)
		fmt.Println("done.")
		fmt.Println("copy file to remote server finished!")
	}
}



