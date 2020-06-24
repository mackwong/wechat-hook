package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
)

type version struct {
	Version         string `json:"version,omitempty"`
	VersionStrategy string `json:"versionStrategy,omitempty"`
	CommitHash      string `json:"commitHash,omitempty"`
	GitBranch       string `json:"gitBranch,omitempty"`
	GitTag          string `json:"gitTag,omitempty"`
	CommitTimestamp string `json:"commitTimestamp,omitempty"`
	GoVersion       string `json:"goVersion,omitempty"`
	Compiler        string `json:"compiler,omitempty"`
	Platform        string `json:"platform,omitempty"`
	// Deprecated
	Os string `json:"os,omitempty"`
	// Deprecated
	Arch string `json:"arch,omitempty"`
	// Deprecated
	BuildTimestamp string `json:"buildTimestamp,omitempty"`
	// Deprecated
	BuildHost string `json:"buildHost,omitempty"`
	// Deprecated
	BuildHostOs string `json:"buildHostOs,omitempty"`
	// Deprecated
	BuildHostArch string `json:"buildHostArch,omitempty"`
}

func (v *version) Print() {
	fmt.Printf("Version = %v\n", v.Version)
	fmt.Printf("VersionStrategy = %v\n", v.VersionStrategy)
	fmt.Printf("GitTag = %v\n", v.GitTag)
	fmt.Printf("GitBranch = %v\n", v.GitBranch)
	fmt.Printf("CommitHash = %v\n", v.CommitHash)
	fmt.Printf("CommitTimestamp = %v\n", v.CommitTimestamp)

	if v.GoVersion != "" {
		fmt.Printf("GoVersion = %v\n", v.GoVersion)
	}
	if v.Compiler != "" {
		fmt.Printf("Compiler = %v\n", v.Compiler)
	}
	if v.Platform != "" {
		fmt.Printf("Platform = %v\n", v.Platform)
	}
}

var Version version

func NewCmdVersion() *cobra.Command {
	var short bool
	cmd := &cobra.Command{
		Use:               "version",
		Short:             "Prints binary version number.",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if short {
				fmt.Print(Version.Version)
			} else {
				Version.Print()
			}
		},
	}
	cmd.Flags().BoolVar(&short, "short", false, "Print just the version number.")
	return cmd
}
