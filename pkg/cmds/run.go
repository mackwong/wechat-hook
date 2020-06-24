package cmds

import (
	"github.com/mackwong/gitllab-wechat-hook/pkg/manager"
	"github.com/spf13/cobra"
	"log"
)

func NewCmdRun() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch Wechat Hooker",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			m := manager.NewManager()
			if err := m.Run(); err != nil {
				log.Fatal(err.Error())
			}
			return nil
		},
	}

	return cmd
}
