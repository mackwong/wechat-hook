package manager

import (
	"github.com/sirupsen/logrus"
	gitlab "github.com/xanzy/go-gitlab"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"time"
)

const MaxMessages = 100

type Manager struct {
	s     *http.Server
	c     chan interface{}
	rules []Rule
}

func NewManager(configFile string) (*Manager, error) {
	conf, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	logrus.Infof("loading config:\n%s\n", conf)
	rules := make([]Rule, 0)
	err = yaml.Unmarshal(conf, &rules)
	if err != nil {
		return nil, err
	}

	mgr := &Manager{
		c:     make(chan interface{}, MaxMessages),
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
	return m.s.ListenAndServe()
}

func (m *Manager) ProcessMessage() {
	go func() {
		for {
			event := <-m.c
			switch event.(type) {
			case gitlab.PushEvent:
				m.push()
			}
		}
	}()
}

func (m *Manager) push() {

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

	m.c <- event
	w.WriteHeader(http.StatusOK)
}
