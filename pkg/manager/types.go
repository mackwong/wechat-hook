package manager

import "github.com/xanzy/go-gitlab"

type Config struct {
	EventType gitlab.EventType `yaml:"event"`
	Type      string           `yaml:"type"`
	Content   string           `yaml:"content"`
}

type Rule struct {
	WeChatHook string   `yaml:"wechat_hook"`
	Repo       string   `yaml:"repo"`
	Config     []Config `yaml:"config"`
}
