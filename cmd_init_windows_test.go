// +build windows

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdInit(t *testing.T) {
	assert := assert.New(t)

	p := filepath.Join(testOutputDir, "test_cmd_init")
	err := os.MkdirAll(p, os.ModePerm)
	assert.NoError(err)

	conf := filepath.Join(p, "AppData", "Roaming")
	app := App{
		UserHomeDir: p,
		UserConfigDir: conf,
	}
	err = app.CmdInit()
	assert.NoError(err)

	// config file
	_, err = os.Stat(filepath.Join(conf, appName, "config.json"))
	assert.False(os.IsNotExist(err))

	// application directory
	appDir := filepath.Join(p, appName)
	_, err = os.Stat(appDir)
	assert.False(os.IsNotExist(err))

	// bin directory
	_, err = os.Stat(filepath.Join(appDir, "bin"))
	assert.False(os.IsNotExist(err))

	// releases directory
	_, err = os.Stat(filepath.Join(appDir, "releases"))
	assert.False(os.IsNotExist(err))

	// re run
	err = app.CmdInit()
	assert.NoError(err)
}
