package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdRoot(t *testing.T) {
	tests := []struct {
		desc    string
		app     App
		param   CommandLineRootParam
		wantErr bool
	}{
		{
			desc: "ok: print list",
			app: App{
				Config: Config{
					RelmaRoot: filepath.Join(testOutputDir, "test_cmd_list_1"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)
			assert.NoError(tt.app.CmdRoot(nil))
		})
	}
}
