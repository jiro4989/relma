// +build windows

package main

func SetHome(path string) {
	os.Setenv("USERPROFILE", path)
}

func SetConfigDir(path string) {
	os.Setenv("AppData", path)
}
