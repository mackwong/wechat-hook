package manager

type GitlabInfo struct {
	Repo string
}

type Rule struct {
	WeChatHook string     `yaml:"wechat_hook"`
	Gitlab     GitlabInfo `yaml:"gitlab"`
	Type       string     `yaml:"type"`
	Content    string     `yaml:"content"`
}
