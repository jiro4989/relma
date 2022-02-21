package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jiro4989/relma/releases"
	"github.com/stretchr/testify/assert"
)

func TestCmdReleasesLockUnlock(t *testing.T) {
	tests := []struct {
		desc      string
		app       App
		ownerRepo string
		lock      bool
		ops       string
		before    releases.Releases
		want      releases.Releases
		wantErr   bool
	}{
		{
			desc: "ok: success",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "releases_test"),
				},
			},
			before: releases.Releases{
				{
					Owner:  "jiro4989",
					Repo:   "nimjson",
					Locked: false,
				},
			},
			ownerRepo: "jiro4989/nimjson",
			lock:      true,
			ops:       "lock",
			want: releases.Releases{
				{
					Owner:  "jiro4989",
					Repo:   "nimjson",
					Locked: true,
				},
			},
			wantErr: false,
		},
		{
			desc: "ng: not found owner/repo",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "releases_test"),
				},
			},
			before: releases.Releases{
				{
					Owner:  "jiro4989",
					Repo:   "sushi",
					Locked: false,
				},
			},
			ownerRepo: "jiro4989/nimjson",
			lock:      true,
			ops:       "lock",
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			// 書き換え前のファイルを準備する
			assert.NoError(os.MkdirAll(tt.app.Config.RelmaRoot, os.ModePerm))
			assert.NoError(tt.app.SaveReleases(tt.before))

			got, err := tt.app.cmdReleasesLockUnlock(tt.ownerRepo, tt.lock, tt.ops)
			if tt.wantErr {
				assert.Error(err)
				assert.Nil(got)
				return
			}
			assert.Equal(tt.want, got)
			assert.NoError(err)
		})
	}
}
