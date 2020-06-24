package wechat

import (
	"net/http"
	"strings"
)

var url = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=609a7fae-cba3-4eae-afe4-7159e3fcfada"

type Message struct {
	WeChat  string        `json:"wechat"`
	Message WechatMessage `json:"message"`
	Data    interface{}   `json:"data"`
}

type WeChatText struct {
	Context       string   `json:"context"`
	MentionedList []string `json:"mentioned_list,omitempty"`
}

type WeChatMarkdown struct {
	Context string `json:"context"`
}

type WechatMessage struct {
	MsgType  string         `json:"msgtype"`
	Markdown WeChatMarkdown `json:"markdown,omitempty"`
	Text     WeChatText     `json:"text,omitempty"`
}

func (m *Message) Send() error {

}

func Send() {
	d := `{
  "msgtype": "text",
  "text": {
    "content": "Hi，我是机器人test\n由于06月24日添加到群"
  }
}`
	r := strings.NewReader(d)
	http.Post(url, "application/json", r)
}
