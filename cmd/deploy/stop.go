package deploy

import (
	"fmt"
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
	session, _ := client.NewSession()
	defer session.Close()
	if err := session.Run(chmodCmd); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("file stop done!")
	}
}
