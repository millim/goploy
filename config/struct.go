package config

//ServerConfig 服务器端配置
type ServerConfig struct {
	User       string   `json:"user"`       //登录名
	NeedSudo	 bool `json:"needSudo"`
	SSHHost    string   `json:"sshHost"`    //ssh登录地址
	SSHPort    string   `json:"sshPort"`    //ssh登录端口，默认22
	ServerDir  string   `json:"serverDir"`  //服务器端放置的目录名
	ServerFile string   `json:"serverFile"` //服务器端放置的文件名
	PidFile    string `json:"pidFile"`
	Command    []string `json:"command"`    //执行的参数
	Env        []string `json:"env"`        //环境变量
}

//LocalConfig 本地的配置
type LocalConfig struct {
	PrivateKey    string `json:"privateKey"`    //本地key的绝对地址
	LocalMainFile string `json:"localMainFile"` //本地上传文件的绝对地址
	ScriptFile    string `json:"scriptFile"`
}
