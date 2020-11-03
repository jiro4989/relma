// +build windows

package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultConfig(t *testing.T) {
	tests := []struct {
		desc    string
		home    string
		want    Config
		wantErr bool
	}{
		{
			desc: "ok: get default config",
			home: `C:\Users\testuser`,
			want: Config{
				RelmaRoot: `C:\Users\testuser\` + appName,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			SetHome(tt.home)

			got, err := DefaultConfig()
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestConfigDir(t *testing.T) {
	tests := []struct {
		desc      string
		configDir string
		want      string
		wantErr   bool
	}{
		{
			desc:      "ok: get default config directory",
			configDir: `C:\Users\testuser\AppData\Roaming\`,
			want:      `C:\Users\testuser\AppData\Roaming\` + appName,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			SetConfigDir(tt.configDir)

			got, err := ConfigDir()
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestCreateConfigDir(t *testing.T) {
	tests := []struct {
		desc      string
		configDir string
		want      string
		wantErr   bool
	}{
		{
			desc:      "ok: create config directory",
			configDir: testOutputDir,
			want:      filepath.Join(testOutputDir, appName),
			wantErr:   false,
		},
		{
			desc:      "ok: config directory was existed",
			configDir: testOutputDir,
			want:      filepath.Join(testOutputDir, appName),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			SetConfigDir(tt.configDir)

			got, err := CreateConfigDir()
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestConfigFile(t *testing.T) {
	tests := []struct {
		desc      string
		configDir string
		want      string
		wantErr   bool
	}{
		{
			desc:      "ok: get default config file",
			configDir: `C:\Users\testuser`,
			want:      filepath.Join(`C:\Users\testuser`, appName, "config.json"),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			SetConfigDir(tt.configDir)

			got, err := ConfigFile()
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestCreateConfigFile(t *testing.T) {
	appDir := filepath.Join(testOutputDir, "test_create_config_file_1", appName)
	err := os.MkdirAll(appDir, os.ModePerm)
	assert.NoError(t, err)

	tests := []struct {
		desc      string
		configDir string
		config    Config
		want      string
		wantErr   bool
	}{
		{
			desc:      "ok: create config file",
			configDir: filepath.Join(testOutputDir, "test_create_config_file_1"),
			config: Config{
				RelmaRoot: "sushi",
			},
			want:    filepath.Join(appDir, "config.json"),
			wantErr: false,
		},
		{
			desc:      "ok: create config file",
			configDir: filepath.Join(testOutputDir, "test_create_config_file_2"),
			config: Config{
				RelmaRoot: "sushi",
			},
			want:    filepath.Join(appDir, "config.json"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			SetConfigDir(tt.configDir)

			got, err := CreateConfigFile(tt.config)
			if tt.wantErr {
				assert.Error(err)
				return
			}

			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestReadReleasesFile(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		want    Releases
		wantErr bool
	}{
		{
			desc: "ok: releases.json exists",
			path: filepath.Join(testDir, "releases.json"),
			want: Releases{
				{
					URL: "https://example.com",
				},
				{
					URL: "https://example2.com",
				},
			},
			wantErr: false,
		},
		{
			desc:    "ok: releases.json doesn't exist",
			path:    "not_found.json",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := ReadReleasesFile(tt.path)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}
