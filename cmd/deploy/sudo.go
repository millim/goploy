package deploy

func sudo()string{
	if serverConfig.NeedSudo {
		return "sudo "
	}
	return ""
}