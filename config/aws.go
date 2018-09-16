package config

import (
	"path"
)

var (
	AwsConfigPath      = path.Join(homePath(), ".aws")
	AwsConfigFile      = path.Join(AwsConfigPath, "config")
	AwsCredentialsFile = path.Join(AwsConfigPath, "credentials")
)
