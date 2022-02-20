//go:build windows
// +build windows

package main

import (
	"os"
)

func SetHome(path string) {
	os.Setenv("USERPROFILE", path)
}

func SetHomeWithRecoverFunc(path string) func() {
	const key = "USERPROFILE"
	orgHome := os.Getenv(key)
	os.Setenv(key, path)
	return func() {
		os.Setenv(key, orgHome)
	}
}

func SetConfigDir(path string) {
	os.Setenv("AppData", path)
}
