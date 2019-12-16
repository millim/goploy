package deploy

import (
	"fmt"
	"github.com/millim/goploy/session"
	"strings"
)

//Start run file
func Start() {
	setting()
	pro := fmt.Sprintf("ls /proc/`cat %s`", serverConfig.PidFile)
	_, err := session.ExecCmdResponse(pro)
	if err != nil {
		env := strings.Join(serverConfig.Env, " ")
		cmd := fmt.Sprintf("%s %s", env, runCmd())
		err := session.ExecCmd(cmd)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("file start done!")
	} else {
		fmt.Println("Proc is run!")
	}

}
