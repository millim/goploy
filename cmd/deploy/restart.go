package deploy

import (
	"fmt"
	"github.com/millim/goploy/session"
	"log"
)
//Restart <----
func Restart() {
	setting()
	if serverConfig.PidFile == ""{
		log.Println("not setting pidfile")
		return
	}
	cmd := fmt.Sprintf("kill -USR2 `cat %s`", serverConfig.PidFile)
	err := session.ExecCmd(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file restart done!")
}
