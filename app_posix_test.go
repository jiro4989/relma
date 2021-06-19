// +build !windows env

// 並列にテストが走ったときに環境変数が上書きされるせいか、テストがコケてしまう
// ため、明示的にテストしないとダメにする

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

			recoverFunc := SetHomeWithRecoverFunc(tt.home)
			defer recoverFunc()

			got, err := DefaultConfig()
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestConfigDir(t *testing.T) {
	tests := []struct {
		desc    string
		home    string
		want    string
		wantErr bool
	}{
		{
			desc:    "ok: get default config directory",
			home:    "/home/testuser",
			want:    "/home/testuser/.config/" + appName,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			recoverFunc := SetHomeWithRecoverFunc(tt.home)
			defer recoverFunc()

			got, err := ConfigDir()
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestCreateConfigDir(t *testing.T) {
	tests := []struct {
		desc    string
		home    string
		want    string
		wantErr bool
	}{
		{
			desc:    "ok: create config directory",
			home:    testOutputDir,
			want:    filepath.Join(testOutputDir, ".config", appName),
			wantErr: false,
		},
		{
			desc:    "ok: config directory was existed",
			home:    testOutputDir,
			want:    filepath.Join(testOutputDir, ".config", appName),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			recoverFunc := SetHomeWithRecoverFunc(tt.home)
			defer recoverFunc()

			got, err := CreateConfigDir()
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestConfigFile(t *testing.T) {
	tests := []struct {
		desc    string
		home    string
		want    string
		wantErr bool
	}{
		{
			desc:    "ok: get default config file",
			home:    "/home/testuser",
			want:    filepath.Join("/home/testuser", ".config", appName, "config.json"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			recoverFunc := SetHomeWithRecoverFunc(tt.home)
			defer recoverFunc()

			got, err := ConfigFile()
			assert.NoError(err)
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
		config  Config
		want    string
		wantErr bool
	}{
		{
			desc: "ok: create config file",
			home: filepath.Join(testOutputDir, "test_create_config_file_1"),
			config: Config{
				RelmaRoot: "sushi",
			},
			want:    filepath.Join(testOutputDir, "test_create_config_file_1", ".config", appName, "config.json"),
			wantErr: false,
		},
		{
			desc: "ok: create config file",
			home: filepath.Join(testOutputDir, "test_create_config_file_2"),
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

			recoverFunc := SetHomeWithRecoverFunc(tt.home)
			defer recoverFunc()

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
