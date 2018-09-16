package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tienne/aws-profiles/config"
	"os"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show aws-profiles version",
	Run:   printVersion,
}

func printVersion(*cobra.Command, []string) {
	fmt.Fprint(os.Stderr, config.AsciiLogo)
	fmt.Println(config.CurrentBuildInfo)
}
