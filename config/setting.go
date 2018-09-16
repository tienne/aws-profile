package config

import (
	"os"
	"path"
	"runtime"
)

func homePath() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}

	return os.Getenv("HOME")
}

var (
	HomePath        = homePath()
	SettingPath     = path.Join(homePath(), ".aws-profiles")
	SettingFileName = "config"
	SettingFilePath = path.Join(SettingPath, SettingPath)
)
