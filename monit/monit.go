package monit

import (
	"bytes"
	"fmt"
	"github.com/millim/goploy/config"
	"github.com/millim/goploy/goflag"
	"github.com/millim/goploy/session"
	"log"
	"strings"
	"text/template"
)

const installUbuntu = "apt-get install monit"
const installCentOS = "yum install monit -y"

var serverConfig *config.ServerConfig
var localConfig *config.LocalConfig

//InstallMonitTo 安装monit，只能CentOS或者Ubuntu
func InstallMonitTo(){
	serverConfig, _ = config.LoadConfigFile()
	s, err := session.ExecCmdResponse("cat /etc/issue")
	if err != nil {
		log.Println(err)
		return
	}

	cmdMod := ""
	if serverConfig.NeedSudo {
		cmdMod = "sudo "
	}

	isCentOS := strings.Contains(s, "Kernel \\r on an \\m")
	if isCentOS {
		cmdMod = fmt.Sprintf("%s%s",cmdMod, installCentOS)
	}

	isUbuntu := strings.Contains(s, "Ubuntu")
	if isUbuntu {
		cmdMod = fmt.Sprintf("%s%s",cmdMod, installUbuntu)
	}

	session.ExecCmd(cmdMod)
}

func SetConfig(){
	serverConfig, localConfig = config.LoadConfigFile()
	args := goflag.Get()
	if len(args) < 3 {
		log.Println("设置monit配置，goploy DirName monit-config 配置文件名")
		return
	}

	s, err := session.ExecCmdResponse("cat /etc/issue")
	if err != nil {
		log.Println(err)
		return
	}

	cmdMod := ""
	if serverConfig.NeedSudo {
		cmdMod = "sudo "
	}


	overText := ""
	isCentOS := strings.Contains(s, "Kernel \\r on an \\m")
	if isCentOS {
		s := getT()
		file := fmt.Sprintf("/etc/monit.d/%s", args[2])
		cmdMod = fmt.Sprintf(`%secho -e "%s"|%stee %s`,cmdMod, s,cmdMod, file)
		overText = fmt.Sprintf("请修改 %s 文件", file)
	}

	isUbuntu := strings.Contains(s, "Ubuntu")
	if isUbuntu {
		s := getT()
		file := fmt.Sprintf("/etc/monit/conf-enabled/%s", args[2])
		cmdMod = fmt.Sprintf(`%secho -e "%s"|%stee %s`,cmdMod, s,cmdMod, file)
		overText = fmt.Sprintf("请修改 %s 文件", file)
	}

	err = session.ExecCmd(cmdMod)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(overText)
}


func getT() string {
	t := template.New("monit_config")
	t, _ = t.Parse(configTemplate())
	b := new(bytes.Buffer)
	cfa := ConfigFileArgs{
		User: serverConfig.User,
		ScriptAddress: localConfig.ScriptFile,
	}
	t.Execute(b, cfa)
	return b.String()
}

type ConfigFileArgs struct {
	FileName string
	User string
	ScriptAddress string
}

//check host gla with address 10.173.167.240
//	start program = "/bin/sh /home/deploy/scripts/run_gla.sh"
//	as uid deploy and gid deploy
//	if failed port 2233 protocol http  request "/api/ping" status 404 for 5 cycles then start
func configTemplate()string{
	s := make([]string,0)
	s = append(s, `check host {{.User}} with address {MODIFY_HOST}`)
	s = append(s, `\tstart program = \"/bin/sh {{.ScriptAddress}}\"`)
	s = append(s, `\tas uid {{.User}} and gid {{.User}}`)
	s = append(s, `\tif failed port {MODIFY_PORT} protocol http  request \"{MODIFY_PING_API}\" status 404 for 5 cycles then start`)
	return strings.Join(s, `\n`)
}





