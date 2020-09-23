package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseURL(t *testing.T) {
	tests := []struct {
		desc    string
		url     string
		want    *Package
		wantErr bool
	}{
		{
			desc: "ok: parsing",
			url:  "https://github.com/itchyny/mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
			want: &Package{
				URL:           "https://github.com/itchyny/mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
				Owner:         "itchyny",
				Repo:          "mmv",
				Version:       "v0.1.2",
				AssetFileName: "mmv_v0.1.2_linux_amd64.tar.gz",
			},
			wantErr: false,
		},
		{
			desc: "ok: GITHUB.COM",
			url:  "https://GITHUB.COM/itchyny/mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
			want: &Package{
				URL:           "https://GITHUB.COM/itchyny/mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
				Owner:         "itchyny",
				Repo:          "mmv",
				Version:       "v0.1.2",
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

func TestIsExecutableFile(t *testing.T) {
	f1, err := os.Open(filepath.Join(testDir, "script.sh"))
	assert.NoError(t, err)
	defer f1.Close()
	f1s, err := f1.Stat()
	assert.NoError(t, err)

	f2, err := os.Open(filepath.Join(testDir, "text.txt"))
	assert.NoError(t, err)
	defer f2.Close()
	f2s, err := f2.Stat()
	assert.NoError(t, err)

	f3, err := os.Open(testDir)
	assert.NoError(t, err)
	defer f3.Close()
	f3s, err := f3.Stat()
	assert.NoError(t, err)

	tests := []struct {
		desc string
		file os.FileInfo
		want bool
	}{
		{
			desc: "ok: executable shellscript",
			file: f1s,
			want: true,
		},
		{
			desc: "ng: text file",
			file: f2s,
			want: false,
		},
		{
			desc: "ng: directory",
			file: f3s,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			got := isExecutableFile(tt.file)
			assert.Equal(tt.want, got)
		})
	}
}
