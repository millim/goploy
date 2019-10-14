package deploy

import (
	"fmt"
	"github.com/millim/goploy/session"
	"log"
)

//Stop <----
func Stop() {
	setting()
	if serverConfig.PidFile == ""{
		log.Println("not setting pidfile")
		return
	}
	chmodCmd := fmt.Sprintf("kill `cat %s`", serverConfig.PidFile)
	err := session.ExecCmd(chmodCmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file stop done!")
}
