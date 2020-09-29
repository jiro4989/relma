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
	os.Setenv("HOME", p)

	app := App{}
	err = app.CmdInit()
	assert.NoError(err)

	// config file
	_, err = os.Stat(filepath.Join(p, ".config", appName, "config.json"))
	assert.False(os.IsNotExist(err))

	// application directory
	_, err = os.Stat(filepath.Join(p, appName))
	assert.False(os.IsNotExist(err))

	// bin directory
	_, err = os.Stat(filepath.Join(p, appName, "bin"))
	assert.False(os.IsNotExist(err))

	// releases directory
	_, err = os.Stat(filepath.Join(p, appName, "releases"))
	assert.False(os.IsNotExist(err))

	// re run
	err = app.CmdInit()
	assert.NoError(err)
}
