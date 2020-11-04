// +build !windows

package main

import (
	"os"
)

func SetHome(path string) {
	os.Setenv("HOME", path)
}

func SetConfigDir(path string) {
	// nothing to do
}
