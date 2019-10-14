package deploy

import (
	"fmt"
	"github.com/millim/goploy/session"
	"strings"
)

//Start run file
func Start() {
	setting()
	env := strings.Join(serverConfig.Env, " ")
	cmd := fmt.Sprintf("%s %s", env, runCmd())
	err := session.ExecCmd(cmd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file start done!")
}
