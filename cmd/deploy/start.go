package deploy

import (
	"fmt"
	"strings"
)

//Start run file
func Start() {
	setting()
	env := strings.Join(serverConfig.Env, " ")
	cmd := fmt.Sprintf("%s %s", env, runCmd())
	execCmd(cmd)
	fmt.Println("file start done!")
}
