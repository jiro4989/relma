// +build !windows

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
			home: "/home/testuser",
			want: Config{
				RelmaRoot: "/home/testuser/" + appName,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			a := App{UserHomeDir: tt.home}
			got := a.DefaultConfig()
			assert.Equal(tt.want, got)
		})
	}
}

func TestConfigDir(t *testing.T) {
	tests := []struct {
		desc    string
		conf    string
		want    string
		wantErr bool
	}{
		{
			desc:    "ok: get default config directory",
			conf:    "/home/testuser/.config",
			want:    "/home/testuser/.config/" + appName,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			a := App{UserConfigDir: tt.conf}
			got := a.ConfigDir()
			assert.Equal(tt.want, got)
		})
	}
}

func TestCreateConfigDir(t *testing.T) {
	tests := []struct {
		desc    string
		home    string
		conf    string
		want    string
		wantErr bool
	}{
		{
			desc:    "ok: create config directory",
			home: testOutputDir,
			conf:    filepath.Join(testOutputDir, ".config"),
			want:    filepath.Join(testOutputDir, ".config", appName),
			wantErr: false,
		},
		{
			desc:    "ok: config directory was existed",
			home: testOutputDir,
			conf:    filepath.Join(testOutputDir, ".config"),
			want:    filepath.Join(testOutputDir, ".config", appName),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			a := App{
				UserHomeDir: tt.home,
				UserConfigDir: tt.conf,
			}
			got, err := a.CreateConfigDir()
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestConfigFile(t *testing.T) {
	tests := []struct {
		desc    string
		home    string
		conf    string
		want    string
		wantErr bool
	}{
		{
			desc:    "ok: get default config file",
			home:    "/home/testuser",
			conf:    "/home/testuser/.config",
			want:    filepath.Join("/home/testuser", ".config", appName, "config.json"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			a := App{
				UserHomeDir: tt.home,
				UserConfigDir: tt.conf,
			}
			got := a.ConfigFile()
			assert.Equal(tt.want, got)
		})
	}
}

func TestCreateConfigFile(t *testing.T) {
	p := filepath.Join(testOutputDir, "test_create_config_file_1", ".config", appName)
	err := os.MkdirAll(p, os.ModePerm)
	assert.NoError(t, err)

	tests := []struct {
		desc    string
		home    string
		conf string
		config  Config
		want    string
		wantErr bool
	}{
		{
			desc: "ok: create config file",
			home: filepath.Join(testOutputDir, "test_create_config_file_1"),
			conf: filepath.Join(testOutputDir, "test_create_config_file_1", ".config"),
			config: Config{
				RelmaRoot: "sushi",
			},
			want:    filepath.Join(testOutputDir, "test_create_config_file_1", ".config", appName, "config.json"),
			wantErr: false,
		},
		{
			desc: "ok: create config file",
			home: filepath.Join(testOutputDir, "test_create_config_file_2"),
			conf: filepath.Join(testOutputDir, "test_create_config_file_2", ".config"),
			config: Config{
				RelmaRoot: "sushi",
			},
			want:    filepath.Join(testOutputDir, "test_create_config_file_2", ".config", appName, "config.json"),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			a := App{
				UserHomeDir: tt.home,
				UserConfigDir: tt.conf,
			}
			got, err := a.CreateConfigFile(tt.config)
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
