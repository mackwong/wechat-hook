package server

import (
	"encoding/json"
	"fmt"
	"github.com/mackwong/gitllab-wechat-hook/pkg/wechat"
	gitlab "github.com/xanzy/go-gitlab"
	"io/ioutil"
	"net/http"
)

type Event struct {
	ObjectKind string `json:"object_kind"`
}

func gitlabHandler(w http.ResponseWriter, r *http.Request) {
	c, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer r.Body.Close()

	var e Event
	err = json.Unmarshal(c, &e)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch e.ObjectKind {
	case "push":
		pushHandler(c)
	default:
		fmt.Printf("can not know %s\n", e.ObjectKind)
	}
	return
}

func pushHandler(b []byte) {
	var e gitlab.PushEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	c, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(c))
	wechat.Send()

}
