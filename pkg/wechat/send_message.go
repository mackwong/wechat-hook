package wechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Message struct {
	WeChat  string        `json:"wechat"`
	Message WechatMessage `json:"message"`
	Data    interface{}   `json:"data"`
}

type WeChatText struct {
	Content       string   `json:"content"`
	MentionedList []string `json:"mentioned_list,omitempty"`
}

type WeChatMarkdown struct {
	Content string `json:"content"`
}

type WechatMessage struct {
	MsgType  string          `json:"msgtype"`
	Markdown *WeChatMarkdown `json:"markdown,omitempty"`
	Text     *WeChatText     `json:"text,omitempty"`
}

func (m *Message) Send() (err error) {
	var tmpl *template.Template
	if m.Message.MsgType == "markdown" {
		tmpl, err = template.New("").Parse(m.Message.Markdown.Content)
	} else {
		tmpl, err = template.New("").Parse(m.Message.Text.Content)
	}
	if err != nil {
		return err
	}
	if tmpl == nil {
		return errors.New("template is nil")
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, m.Data)
	if err != nil {
		return err
	}
	if m.Message.MsgType == "markdown" {
		m.Message.Markdown.Content = buf.String()
		m.Message.Text = nil
	} else {
		m.Message.Text.Content = buf.String()
		m.Message.Markdown = nil
	}

	out, err := json.MarshalIndent(m.Message, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("%s", out)

	resp, err := http.Post(m.WeChat, "application/json", bytes.NewBuffer(out))
	if err != nil {
		return err
	}
	respText, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", respText)
	return err
}
