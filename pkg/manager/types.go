package manager

import "github.com/xanzy/go-gitlab"

const (
	Markdown string = "markdown"
	Text     string = "text"
)

type Config struct {
	EventType     gitlab.EventType `yaml:"event"`
	Type          string           `yaml:"type"`
	Content       string           `yaml:"content"`
	MentionedList []string         `json:"mentioned_list,omitempty"`
}

type Rule struct {
	WeChatHook string   `yaml:"wechat_hook"`
	Config     []Config `yaml:"config"`
}
