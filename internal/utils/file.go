package utils

import (
	"os"
	"strings"
)

// FilePath gets file path with expanding home dir `~`, $HOME. also, $PWD
func FilePath(f string) string {
	homeDir := func(value string) string {
		switch value {
		case "HOME":
			userHome, _ := os.UserHomeDir()
			return userHome
		case "PWD":
			userHome, _ := os.Getwd()
			return userHome
		}
		return ""
	}

	// replace tilde ~ to $HOME and use os.Expand
	if strings.HasPrefix(f, "~") {
		f = strings.Replace(f, "~", "$HOME", 1)
	}

	return os.Expand(f, homeDir)
}
