package commands

import (
	"github.com/spf13/cobra"
	"github.com/tienne/aws-profiles/setting"
)

func selectProfile(*cobra.Command, []string) {

	setting.InitSetting()

	// prompt := promptui.Select{
	// 	Label: " Day",
	// 	Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
	// 		"Saturday", "Sunday"},
	// }
	// _, result, err := prompt.Run()
	//
	// if err != nil {
	// 	fmt.Printf("Prompt failed %v\n", err)
	// 	return
	// }
	//
	// fmt.Printf("You choose %q\n", result)
}
