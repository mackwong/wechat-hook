package wechat

import (
	"net/http"
	"strings"
)

var url = "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=609a7fae-cba3-4eae-afe4-7159e3fcfada"

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
