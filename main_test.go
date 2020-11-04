// +build !windows

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	tests := []struct {
		desc    string
		home    string
		confDir string
		args    []string
		wantErr bool
	}{
		{
			desc:    "",
			home:    filepath.Join(testOutputDir, "test_main_1"),
			confDir: filepath.Join(testOutputDir, "test_main_1", ".config", appName),
			args:    []string{"init"},
			wantErr: false,
		},
		{
			desc:    "",
			home:    filepath.Join(testOutputDir, "test_main_2"),
			confDir: filepath.Join(testOutputDir, "test_main_2", ".config", appName),
			args:    []string{"hogefuga"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			assert.NoError(os.MkdirAll(tt.home, os.ModePerm))
			SetHome(tt.home)
			assert.NoError(os.MkdirAll(tt.confDir, os.ModePerm))

			err := Main(tt.args)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
		})
	}
}
