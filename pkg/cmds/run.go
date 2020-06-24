package cmds

import (
	"github.com/mackwong/gitllab-wechat-hook/pkg/manager"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch Wechat Hooker",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := manager.NewManager("config.yaml")
			if err != nil {
				logrus.Fatal(err.Error())
			}

			if err = m.Run(); err != nil {
				log.Fatal(err.Error())
			}
			return nil
		},
	}

	return cmd
}
