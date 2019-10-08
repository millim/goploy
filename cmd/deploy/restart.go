package deploy

import (
	"fmt"
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
	execCmd(cmd)
	fmt.Println("file restart done!")
}
