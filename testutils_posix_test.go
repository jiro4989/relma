// +build !windows

package main

import (
	"os"
)

func SetHome(path string) {
	os.Setenv("HOME", path)
}

func SetHomeWithRecoverFunc(path string) func() {
	const key = "HOME"
	orgHome := os.Getenv(key)
	os.Setenv(key, path)
	return func() {
		os.Setenv(key, orgHome)
	}
}

func SetConfigDir(path string) {
	// nothing to do
}
