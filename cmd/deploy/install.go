package deploy

import "fmt"

//Install create dir
func Install(){
	setting()
	defer session.Close()

	execCmd(fmt.Sprintf("%s mkdir -p %s", sudo(), serverConfig.ServerDir))
	execCmd(fmt.Sprintf("%schown -R %s %s", sudo(), serverConfig.User, serverConfig.ServerDir))
}
