//go:build !windows
// +build !windows

package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExecutableFile(t *testing.T) {
	tests := []struct {
		desc string
		path string
		want bool
	}{
		{
			desc: "ok: executable shellscript",
			path: filepath.Join(testDir, "script.sh"),
			want: true,
		},
		{
			desc: "ok: executable batch file",
			path: filepath.Join(testDir, "script.bat"),
			want: true,
		},
		{
			desc: "ok: executable binary (linux)",
			path: filepath.Join(testDir, "bin"),
			want: true,
		},
		{
			desc: "ok: executable binary (windows)",
			path: filepath.Join(testDir, "bin.exe"),
			want: true,
		},
		// TODO:
		// {
		// 	desc: "ok: executable binary (darwin)",
		// 	path: filepath.Join(testDir, "darwin"),
		// 	want: true,
		// },
		{
			desc: "ng: text file",
			path: filepath.Join(testDir, "text.txt"),
			want: false,
		},
		{
			desc: "ng: directory",
			path: testDir,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			file, err := os.Open(tt.path)
			assert.NoError(err)
			defer file.Close()
			fi, err := file.Stat()
			assert.NoError(err)

			got, err := IsExecutableFile(fi, tt.path)
			assert.Equal(tt.want, got)
			assert.NoError(err)
		})
	}
}
