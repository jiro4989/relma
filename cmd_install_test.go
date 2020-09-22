package main

import (
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
			url:     "https://gitlab.com//mmv/releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (no repo)",
			url:     "https://gitlab.com/hoge//releases/download/v0.1.2/mmv_v0.1.2_linux_amd64.tar.gz",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (no version)",
			url:     "https://gitlab.com/hoge/fuga/releases/download//mmv_v0.1.2_linux_amd64.tar.gz",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (no asset file)",
			url:     "https://gitlab.com/hoge/fuga/releases/download/v0.1.2/",
			want:    nil,
			wantErr: true,
		},
		{
			desc:    "ng: illegal URL (no asset file)",
			url:     "https://gitlab.com/itchyny/mmv/releases/download/v0.1.2",
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
