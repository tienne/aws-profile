package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tienne/aws-profiles/config"
)

var (
	versionGlobalFlag bool
)

func init() {
	RootCmd.Flags().BoolVarP(&versionGlobalFlag, "version", "v", false, "Print aws-profiles version")
}

var RootCmd = &cobra.Command{
	Use:   config.AppName,
	Short: "aws cli profiles switch",
	Long:  fmt.Sprintf("%s is aws cli profile switch tool", config.AppName),
	RunE: func(c *cobra.Command, args []string) error {
		if versionGlobalFlag {
			printVersion(c, args)
			return nil
		}

		selectProfile(c, args)
		return nil
	},
}
