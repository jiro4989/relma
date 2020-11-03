// +build windows

package main

import (
	"os"
)

func SetHome(path string) {
	os.Setenv("USERPROFILE", path)
}

func SetConfigDir(path string) {
	os.Setenv("AppData", path)
}
