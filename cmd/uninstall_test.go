package cmd

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jiro4989/relma/releases"
	"github.com/stretchr/testify/assert"
)

func TestUninstallableRelease(t *testing.T) {
	tests := []struct {
		desc      string
		rels      releases.Releases
		ownerRepo string
		want      *releases.Release
		wantErr   bool
	}{
		{
			desc: "ok: uninstallable",
			rels: releases.Releases{
				{
					Owner: "JIRO4989",
					Repo:  "sushi",
				},
				{
					Owner: "JIRO4989",
					Repo:  "TEXTIMG",
				},
			},
			want: &releases.Release{
				Owner: "JIRO4989",
				Repo:  "TEXTIMG",
			},
			ownerRepo: "jiro4989/textimg",
			wantErr:   false,
		},
		{
			desc: "ng: un match",
			want: &releases.Release{
				Owner: "jiro4989",
				Repo:  "textimg",
			},
			ownerRepo: "jiro4989/monit",
			wantErr:   true,
		},
		{
			desc: "ng: illegal ownerRepo",
			want: &releases.Release{
				Owner: "jiro4989",
				Repo:  "textimg",
			},
			ownerRepo: "jiro4989monit",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := uninstallableRelease(tt.rels, tt.ownerRepo)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.Equal(tt.want, got)
			assert.NoError(err)
		})
	}
}

func TestUninstallRelease(t *testing.T) {
	dir1 := filepath.Join(testOutputDir, "cmd_uninstall_release_test_1")
	conf1 := Config{
		RelmaRoot: dir1,
	}
	for _, dir := range []string{dir1, conf1.BinDir(), conf1.ReleasesDir()} {
		err := os.MkdirAll(dir, os.ModePerm)
		assert.NoError(t, err)
	}
	for _, file := range []string{"sample1", "sample2"} {
		f := filepath.Join(conf1.BinDir(), file)
		err := ioutil.WriteFile(f, []byte{1}, os.ModePerm)
		assert.NoError(t, err)
	}

	dir2 := filepath.Join(testOutputDir, "cmd_uninstall_release_test_2")
	conf2 := Config{
		RelmaRoot: dir2,
	}
	for _, dir := range []string{dir2, conf2.BinDir(), conf2.ReleasesDir()} {
		err := os.MkdirAll(dir, os.ModePerm)
		assert.NoError(t, err)
	}
	for _, file := range []string{"sample1"} {
		f := filepath.Join(conf2.BinDir(), file)
		err := ioutil.WriteFile(f, []byte{1}, os.ModePerm)
		assert.NoError(t, err)
	}

	tests := []struct {
		desc      string
		app       App
		rel       *releases.Release
		param     *CmdUninstallParam
		wantCount int
		wantErr   bool
	}{
		{
			desc: "ok: success uninstall",
			app: App{
				Config: conf1,
			},
			rel: &releases.Release{
				Owner: "jiro4989",
				Repo:  "textimg",
				InstalledFiles: releases.InstalledFiles{
					{
						Dest: "sample1",
					},
					{
						Dest: "sample2",
					},
				},
			},
			param: &CmdUninstallParam{
				OwnerRepo: "jiro4989/textimg",
			},
			wantCount: 3,
			wantErr:   false,
		},
		{
			desc: "ng: uninstall target doesn't exist",
			app: App{
				Config: conf2,
			},
			rel: &releases.Release{
				Owner: "jiro4989",
				Repo:  "textimg",
				InstalledFiles: releases.InstalledFiles{
					{
						Dest: "sample2",
					},
				},
			},
			param: &CmdUninstallParam{
				OwnerRepo: "jiro4989/textimg",
			},
			wantCount: 0,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := tt.app.uninstallRelease(tt.rel, tt.param)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.Equal(tt.wantCount, len(got))
			assert.NoError(err)
		})
	}
}
