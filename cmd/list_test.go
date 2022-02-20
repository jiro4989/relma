package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/jiro4989/relma/releases"
	"github.com/stretchr/testify/assert"
)

func TestCmdList(t *testing.T) {
	tests := []struct {
		desc    string
		app     App
		rels    releases.Releases
		wantErr bool
	}{
		{
			desc: "ok: print list",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "test_cmd_list_1"),
				},
			},
			rels: releases.Releases{
				{
					Owner:   "jiro4989",
					Repo:    "nimjson",
					Version: "v1.2.6",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			err := os.MkdirAll(tt.app.Config.RelmaRoot, os.ModePerm)
			assert.NoError(err)

			b, err := json.MarshalIndent(&tt.rels, "", "  ")
			assert.NoError(err)

			f := tt.app.Config.ReleasesFile()
			err = ioutil.WriteFile(f, b, os.ModePerm)
			assert.NoError(err)

			err = tt.app.CmdList()
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
		})
	}
}
