package setting

import (
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/tienne/aws-profiles/config"
	"os"
)

var (
	awsProfileNames []string
)

// if not settings initialize setting directory
func InitSetting() error {
	_, err := os.Stat(config.SettingPath)

	if os.IsNotExist(err) {
		// err := os.Mkdir(config.SettingPath, 0775)
		//
		// if err != nil {
		// 	return errors.New("Do Not Create setting directory: " + err.Error())
		// }

		// fmt.Print(config.AwsConfigFile)
		// var configs AwsProfileConfig
		configs, err := ini.Load(config.AwsConfigFile)

		if err != nil {
			return errors.New("aws config file load fail: " + err.Error())
		}

		credentialConfig, err := ini.Load(config.AwsCredentialsFile)

		if err != nil {
			return errors.New("Aws credentials file load fail: " + err.Error())
		}

		// credentialConfig.Section("default").MapTo(credential)
		// fmt.Println(credential.AccessKeyId)
		for _, name := range configs.SectionStrings() {
			if name == "DEFAULT" {
				continue
			}

			awsProfileNames = append(awsProfileNames, name)
			// var credential = new(AwsCredential)
		}
		//
		fmt.Println(credentialConfig.SectionStrings())
		fmt.Println(awsProfileNames)
		// fmt.Println(configs.GetSection("default"))
		// fmt.Println(err)
	}

	// already init settings or success
	return nil
}
