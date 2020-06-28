package manager

import (
	"encoding/json"
	"fmt"
	"github.com/mackwong/gitllab-wechat-hook/pkg/wechat"
	"github.com/sirupsen/logrus"
	gitlab "github.com/xanzy/go-gitlab"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"time"
)

const MaxMessages = 100

type EventInfo struct {
	ProjectName string `json:"project_name"`
	EventType   gitlab.EventType
	Event       interface{} `json:"event"`
}

type Manager struct {
	s     *http.Server
	c     chan EventInfo
	rules []Rule
}

func NewManager(configFile string) (*Manager, error) {
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	rules := make([]Rule, 0)
	err = yaml.Unmarshal(conf, &rules)
	if err != nil {
		return nil, err
	}

	mgr := &Manager{
		c:     make(chan EventInfo, MaxMessages),
		rules: rules,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/gitlab", mgr.gitlabHandler)
	mgr.s = &http.Server{
		Addr:         "0.0.0.0:9999",
		WriteTimeout: time.Second * 5,
		Handler:      mux,
	}
	return mgr, nil
}

func (m *Manager) Run() error {
	go m.ProcessMessage()
	return m.s.ListenAndServe()
}

func (m *Manager) ProcessMessage() {
	for {
		event := <-m.c
		for _, r := range m.rules {
			for _, c := range r.Config {
				if c.EventType == event.EventType {
					if event.EventType == gitlab.EventTypePipeline {
						status := event.Event.(*gitlab.PipelineEvent).ObjectAttributes.Status
						if status == "success" {
							continue
						}
					}

					msg := wechat.Message{
						WeChat: r.WeChatHook,
						Message: wechat.WechatMessage{
							MsgType: c.Type,
						},
						Data: event.Event,
					}
					if c.Type == Markdown {
						msg.Message = wechat.WechatMessage{
							MsgType:  c.Type,
							Markdown: &wechat.WeChatMarkdown{Content: c.Content},
						}
					} else {
						msg.Message = wechat.WechatMessage{
							MsgType: c.Type,
							Text:    &wechat.WeChatText{Content: c.Content, MentionedList: c.MentionedList},
						}
					}
					if err := msg.Send(); err != nil {
						logrus.Errorf("Send message error: %s", err.Error())
					}
				}
			}
		}
	}
}

func (m *Manager) push(name string) {
	logrus.Info(name)
}

func (m *Manager) gitlabHandler(w http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logrus.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer func() {
		_ = r.Body.Close()
	}()

	eventType := gitlab.HookEventType(r)
	event, err := gitlab.ParseHook(eventType, payload)
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	o, _ := json.MarshalIndent(event, "", "  ")
	fmt.Printf("%s", o)

	info := EventInfo{
		Event:     event,
		EventType: eventType,
	}

	m.c <- info
	w.WriteHeader(http.StatusOK)
}
