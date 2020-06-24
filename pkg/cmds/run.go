package cmds

import (
	"github.com/mackwong/gitllab-wechat-hook/pkg/server"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch AppsCode Service Broker",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			s := server.NewServer()
			if err := s.ListenAndServe(); err != nil {
				log.Fatal(err.Error())
			}
			return nil
		},
	}

	return cmd
}
