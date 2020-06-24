package main

import (
	"github.com/mackwong/gitllab-wechat-hook/pkg/cmds"
	"log"
)

func main() {
	if err := cmds.NewRootCmd().Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
