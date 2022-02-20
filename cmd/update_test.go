package cmd

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jiro4989/relma/external/github"
	"github.com/jiro4989/relma/releases"
	"github.com/stretchr/testify/assert"
)

func TestCmdUpdate(t *testing.T) {
	tests := []struct {
		desc    string
		app     App
		rel     releases.Releases
		param   *CmdUpdateParam
		wantErr bool
	}{
		{
			desc: "ok: update",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "test_cmd_update"),
				},
				GitHubClient: &github.MockClient{
					LatestTag: "v1.2.6",
					Err:       nil,
				},
			},
			param: &CmdUpdateParam{},
			rel: releases.Releases{
				{
					URL:           "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
					Owner:         "jiro4989",
					Repo:          "nimjson",
					Version:       "v1.2.6",
					AssetFileName: "nimjson_linux.tar.gz",
					InstalledFiles: releases.InstalledFiles{
						{},
					},
				},
			},
			wantErr: false,
		},
		{
			desc: "ng: raise error",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "test_cmd_update"),
				},
				GitHubClient: &github.MockClient{
					Err: errors.New("error"),
				},
			},
			param:   &CmdUpdateParam{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			dir := tt.app.Config.RelmaRoot
			err := os.MkdirAll(dir, os.ModePerm)
			assert.NoError(err)

			f := tt.app.Config.ReleasesFile()
			b, err := json.Marshal(tt.rel)
			assert.NoError(err)
			err = ioutil.WriteFile(f, b, os.ModePerm)
			assert.NoError(err)

			err = tt.app.CmdUpdate(tt.param)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)

			b, err = ioutil.ReadFile(f)
			assert.NoError(err)

			var rs releases.Releases
			err = json.Unmarshal(b, &rs)
			assert.NoError(err)
			for i, r := range rs {
				rel := tt.rel[i]
				assert.NotEqual(r.LatestVersion, rel.LatestVersion)
				r.LatestVersion = rel.LatestVersion
				assert.Equal(r, rel)
			}
		})
	}
}
