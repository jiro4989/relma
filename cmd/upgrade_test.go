package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jiro4989/relma/external/downloader"
	"github.com/jiro4989/relma/releases"
	"github.com/stretchr/testify/assert"
)

func TestCmdUpgrade(t *testing.T) {
	validReleases := releases.Releases{
		{
			URL:           "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
			Owner:         "Jiro4989",
			Repo:          "textimg",
			Version:       "v1.2.6",
			LatestVersion: "v1.2.8",
		},
	}
	invalidReleases := releases.Releases{
		{
			URL:           "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
			Owner:         "Jiro4989",
			Repo:          "textimg",
			Version:       "v1.2.6",
			LatestVersion: "v1.2.6",
		},
	}
	invalidReleases2 := releases.Releases{
		{
			URL:           "https://gitlab.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
			Owner:         "Jiro4989",
			Repo:          "textimg",
			Version:       "v1.2.6",
			LatestVersion: "v1.2.8",
		},
	}

	tests := []struct {
		desc    string
		app     App
		rels    releases.Releases
		param   *CmdUpgradeParam
		wantErr bool
		wantNil bool
	}{
		{
			desc: "ok: upgrade",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "cmd_upgrade_1"),
				},
				Downloader: &downloader.MockDownloader{
					Body: nimjson1_2_6,
				},
			},
			rels: validReleases,
			param: &CmdUpgradeParam{
				Yes: true,
			},
			wantErr: false,
		},
		{
			desc: "ng: illegal owner/repo",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "cmd_upgrade_2"),
				},
				Downloader: &downloader.MockDownloader{
					Body: nimjson1_2_6,
				},
			},
			rels: validReleases,
			param: &CmdUpgradeParam{
				Yes:       true,
				OwnerRepo: "jiro4989textimg",
			},
			wantErr: true,
		},
		{
			desc: "ng: no upgradable releases",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "cmd_upgrade_3"),
				},
				Downloader: &downloader.MockDownloader{
					Body: nimjson1_2_6,
				},
			},
			rels: invalidReleases,
			param: &CmdUpgradeParam{
				Yes: true,
			},
			wantErr: false,
			wantNil: true,
		},
		{
			desc: "ng: failed to install",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "cmd_upgrade_4"),
				},
				Downloader: &downloader.MockDownloader{
					Body: nimjson1_2_6,
				},
			},
			rels: invalidReleases2,
			param: &CmdUpgradeParam{
				Yes: true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			err := os.MkdirAll(tt.app.Config.RelmaRoot, os.ModePerm)
			assert.NoError(err)

			err = os.MkdirAll(tt.app.Config.BinDir(), os.ModePerm)
			assert.NoError(err)

			err = os.MkdirAll(tt.app.Config.ReleasesDir(), os.ModePerm)
			assert.NoError(err)

			err = tt.app.cmdUpgrade(tt.rels, tt.param)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
			if tt.wantNil {
				return
			}

			rels, err := tt.app.Config.ReadReleasesFile()
			assert.NoError(err)
			rel := rels[0]
			assert.Equal(rel.Version, rel.LatestVersion)
		})
	}
}

func TestSearchReleaseOrDefault(t *testing.T) {
	tests := []struct {
		desc      string
		rels      releases.Releases
		ownerRepo string
		want      releases.Releases
		wantErr   bool
	}{
		{
			desc: "ok: default releases",
			rels: releases.Releases{
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			ownerRepo: "",
			want: releases.Releases{
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			wantErr: false,
		},
		{
			desc: "ok: found owner/repo",
			rels: releases.Releases{
				{
					Owner: "jiro4989",
					Repo:  "sushi",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			ownerRepo: "JIRO4989/TEXTIMG",
			want: releases.Releases{
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			wantErr: false,
		},
		{
			desc: "ng: not found owner/repo",
			rels: releases.Releases{
				{
					Owner: "jiro4989",
					Repo:  "sushi",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			ownerRepo: "jiro4989/onigiri",
			want:      nil,
			wantErr:   true,
		},
		{
			desc: "ng: illegal owner/repo",
			rels: releases.Releases{
				{
					Owner: "jiro4989",
					Repo:  "sushi",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			ownerRepo: "jiro4989textimg",
			want:      nil,
			wantErr:   true,
		},
		{
			desc:      "ng: releases are nil",
			rels:      nil,
			ownerRepo: "jiro4989/textimg",
			want:      nil,
			wantErr:   true,
		},
		{
			desc:      "ng: releases are empty",
			rels:      releases.Releases{},
			ownerRepo: "jiro4989/textimg",
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := searchReleaseOrDefault(tt.rels, tt.ownerRepo)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}

func TestSearchRelease(t *testing.T) {
	tests := []struct {
		desc      string
		rels      releases.Releases
		ownerRepo string
		want      releases.Releases
		wantErr   bool
	}{
		{
			desc: "ok: found releases",
			rels: releases.Releases{
				{
					Owner: "Jiro",
					Repo:  "Test",
				},
				{
					Owner: "Jiro4989",
					Repo:  "TestA",
				},
				{
					Owner: "Jiro",
					Repo:  "TestA",
				},
				{
					Owner: "Jiro4989",
					Repo:  "Test",
				},
			},
			ownerRepo: "JIRO4989/TEST",
			want: releases.Releases{
				{
					Owner: "Jiro4989",
					Repo:  "Test",
				},
			},
			wantErr: false,
		},
		{
			desc:      "ok: releases is empty",
			rels:      releases.Releases{},
			ownerRepo: "",
			want:      nil,
			wantErr:   false,
		},
		{
			desc: "ng: ownerRepo is illegal",
			rels: releases.Releases{
				{
					Owner: "Jiro4989",
					Repo:  "Test",
				},
			},
			ownerRepo: "JIRO4989TEST",
			want:      nil,
			wantErr:   true,
		},
		{
			desc: "ng: ownerRepo is empty",
			rels: releases.Releases{
				{
					Owner: "Jiro4989",
					Repo:  "Test",
				},
			},
			ownerRepo: "",
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := searchRelease(tt.rels, tt.ownerRepo)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}
