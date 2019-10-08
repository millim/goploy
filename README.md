### goploy 
golang 单机部署用配置

目前用于 github.com/facebookgo/grace/gracehttp 部署时，方便本机直接部署重启使用.
此内容直接通过ssh进行操作，尝试用项目，一定程度上没有直接shell用着方便。

### 安装
go install github.com/millim/goploy

### 命令说明
goploy init DIRNAME 
项目中使用，初始化一套配置文件，其中有本地文件和服务器配置文件.

goploy/local.json
```json
{
  "privateKey": "/xxx/.ssh/id_rsa",
  "localMainFile": "main",
  "scriptFile":"~/scripts/job.sh"
}
```
`privateKey:` 本地的key文件，建议完整路径
`localMainFile:` 指定要copy的文件（一般为编译后的上传文件)，建议完整路径
`scriptFile:` 当执行deploy命令时，将会把完整的执行命令以sh文件放到此处

goploy/DIRNAME/config.json
```json
{
	"user":"",
	"sshHost": "",
	"sshPort": "",
	"serverDir": "",
	"serverFile": "",
	"needSudo": false,
	"pidFile": "",
	"env": ["GIN_MODE=release"],
	"command": ["--port=6061"]
}
```
`user:` 登录的用户名
`sshHost:` 登录的host
`sshPort:` 登录的port，默认为22
`serverDir:` 服务器端的目录名
`serverFile:` 服务器端的文件名
`needSudo:` 是否需要使用sudo命令来创建服务器端的目录
`pidFile:` 指定pid文件，没有时将无法使用stop和restart命令
`env:` 数组，start命令时的环境变量设置
`command:` 数组，start命令时的参数设置

---

`goploy DIRNAME install` 创建目录
`goploy DIRNAME deploy`  部署文件
`goploy DIRNAME start`   执行服务端文件
`goploy DIRNAME stop`    停止服务端文件
`goploy DIRNAME restart` 重启服务端文件