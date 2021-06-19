// +build !windows

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdInstall(t *testing.T) {
	tests := []struct {
		desc      string
		app       App
		param     *CmdInstallParam
		want      Releases
		wantCount int
		wantErr   bool
	}{
		{
			desc: "ok: installing",
			app: App{
				Config: Config{
					RelmaRoot: testOutputDir,
				},
			},
			param: &CmdInstallParam{
				URL: "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
			},
			want: Releases{
				{
					URL:           "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
					Owner:         "jiro4989",
					Repo:          "nimjson",
					Version:       "v1.2.6",
					AssetFileName: "nimjson_linux.tar.gz",
					InstalledFiles: InstalledFiles{
						{
							Src:  filepath.Join("bin", "nimjson"),
							Dest: "nimjson",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			desc: "ok: installing with releases.json",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "cmd_install_test"),
				},
			},
			param: &CmdInstallParam{
				File: filepath.Join(testDir, "cmd_install_releases.json"),
			},
			want: Releases{
				{
					URL:           "https://github.com/jiro4989/nimjson/releases/download/v1.2.8/nimjson_linux.tar.gz",
					Owner:         "jiro4989",
					Repo:          "nimjson",
					Version:       "v1.2.8",
					AssetFileName: "nimjson_linux.tar.gz",
					InstalledFiles: InstalledFiles{
						{
							Src:  filepath.Join("bin", "nimjson"),
							Dest: "nimjson",
						},
					},
				},
				{
					URL:           "https://github.com/jiro4989/monit/releases/download/1.1.0/monit_linux.tar.gz",
					Owner:         "jiro4989",
					Repo:          "monit",
					Version:       "1.1.0",
					AssetFileName: "monit_linux.tar.gz",
					InstalledFiles: InstalledFiles{
						{
							Src:  filepath.Join("bin", "monit"),
							Dest: "monit",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			assert.NoError(os.MkdirAll(tt.app.Config.BinDir(), os.ModePerm))
			assert.NoError(os.MkdirAll(tt.app.Config.ReleasesDir(), os.ModePerm))

			p := tt.app.Config.ReleasesFile()
			os.Remove(p)

			err := tt.app.CmdInstall(tt.param)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)

			b, err := ioutil.ReadFile(p)
			assert.NoError(err)

			var rels Releases
			err = json.Unmarshal(b, &rels)
			assert.NoError(err)
			assert.Equal(len(tt.want), len(rels))

			for i, want := range tt.want {
				rel := rels[i]
				assert.Equal(want.URL, rel.URL)
				assert.Equal(want.Owner, rel.Owner)
				assert.Equal(want.Repo, rel.Repo)
				assert.Equal(want.Version, rel.Version)
				assert.Equal(want.AssetFileName, rel.AssetFileName)
				assert.Equal(want.InstalledFiles, rel.InstalledFiles)
			}
		})
	}
}

func TestParseURL(t *testing.T) {
	tests := []struct {
		desc    string
		url     string
		want    *Release
		wantErr bool
	}{
		{
			desc: "ok: parsing",
			url:  "https://github.com/itchyny/mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
			want: &Release{
				URL:           "https://github.com/itchyny/mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
				Owner:         "itchyny",
				Repo:          "mmv",
				Version:       "v0.1.2",
				LatestVersion: "v0.1.2",
				AssetFileName: "mmv_v0.1.2_linux_amd64.tar.gz",
			},
			wantErr: false,
		},
		{
			desc: "ok: GITHUB.COM",
			url:  "https://GITHUB.COM/itchyny/mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
			want: &Release{
				URL:           "https://GITHUB.COM/itchyny/mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
				Owner:         "itchyny",
				Repo:          "mmv",
				Version:       "v0.1.2",
				LatestVersion: "v0.1.2",
				AssetFileName: "mmv_v0.1.2_linux_amd64.tar.gz",
			},
			wantErr: false,
		},
		{
			desc:    "ng: gitlab.com domain",
			url:     "https://gitlab.com/itchyny/mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (no owner)",
			url:     "https://github.com//mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (no repo)",
			url:     "https://github.com/hoge//releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (no version)",
			url:     "https://github.com/hoge/fuga/releases/download//mmv_v0.1.2_linux_amd64.tar.gz",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (no asset file)",
			url:     "https://github.com/hoge/fuga/releases/download/v0.1.2/",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (no asset file)",
			url:     "https://github.com/itchyny/mmv/releases/download/v0.1.2",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (empty)",
			url:     "",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got, err := parseURL(tt.url)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.Equal(tt.want, got)
			assert.NoError(err)
		})
	}
}

func TestDownloadFile(t *testing.T) {
	tests := []struct {
		desc     string
		url      string
		destDir  string
		destFile string
		want     string
		wantErr  bool
	}{
		{
			desc:     "ok: download file",
			url:      "https://github.com/jiro4989",
			destDir:  testOutputDir,
			destFile: "out.html",
			want:     filepath.Join(testOutputDir, "out.html"),
			wantErr:  false,
		},
		{
			desc:     "ng: empty url",
			url:      "",
			destDir:  testOutputDir,
			destFile: "out.html",
			want:     "",
			wantErr:  true,
		},
		{
			desc:     "ng: not exist directory",
			url:      "https://github.com/jiro4989",
			destDir:  "foobar",
			destFile: "out.html",
			want:     "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got, err := downloadFile(tt.url, tt.destDir, tt.destFile)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.Equal(tt.want, got)
			assert.NoError(err)

			_, err = os.Stat(got)
			// assert.True(os.IsExist(err)) // why error?
			assert.False(os.IsNotExist(err))

			b, err := ioutil.ReadFile(got)
			assert.NoError(err)
			assert.True(0 < len(b))
		})
	}
}

func TestInstallFiles(t *testing.T) {
	tests := []struct {
		desc    string
		srcDir  string
		destDir string
		want    int
		wantErr bool
	}{
		{
			desc:    "ok: install files",
			srcDir:  filepath.Join(testDir, "test_install_files"),
			destDir: testOutputDir,
			want:    2,
			wantErr: false,
		},
		{
			desc:    "ok: install file (nested directory)",
			srcDir:  filepath.Join(testDir, "test_install_files_2"),
			destDir: testOutputDir,
			want:    1,
			wantErr: false,
		},
		{
			desc:    "ok: install file (nested directory and bin directory)",
			srcDir:  filepath.Join(testDir, "test_install_files_3"),
			destDir: testOutputDir,
			want:    1,
			wantErr: false,
		},
		{
			desc:    "ng: src directory is not exist dir",
			srcDir:  "not_exist",
			destDir: testOutputDir,
			want:    0,
			wantErr: true,
		},
		{
			desc:    "ng: dest directory is not exist dir",
			srcDir:  filepath.Join(testDir, "test_install_files"),
			destDir: "not_exist",
			want:    0,
			wantErr: true,
		},
		{
			desc:    "ng: dest is file",
			srcDir:  filepath.Join(testDir, "test_install_files"),
			destDir: "go.mod",
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got, err := installFiles(tt.srcDir, tt.destDir)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.Len(got, tt.want)
			assert.NoError(err)
		})
	}
}
