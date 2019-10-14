package deploy

import (
	"fmt"
	"github.com/millim/goploy/session"
)

//Install create dir
func Install(){
	setting()
	session.ExecCmd(fmt.Sprintf("%s mkdir -p %s", sudo(), serverConfig.ServerDir))
	session.ExecCmd(fmt.Sprintf("%schown -R %s %s", sudo(), serverConfig.User, serverConfig.ServerDir))
}
