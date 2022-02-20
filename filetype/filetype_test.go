package cmd

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsArchiveFile(t *testing.T) {
	tests := []struct {
		desc string
		path string
		want bool
	}{
		{
			desc: "ok: zip file is archive file",
			path: filepath.Join(testDir, "archivefile.zip"),
			want: true,
		},
		{
			desc: "ok: gzip file is archive file",
			path: filepath.Join(testDir, "archivefile.gz"),
			want: true,
		},
		{
			desc: "ok: tar.gz file is archive file",
			path: filepath.Join(testDir, "archivefile.tar.gz"),
			want: true,
		},
		{
			desc: "ok: shellscript is not archive file",
			path: filepath.Join(testDir, "archivefile.sh"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := IsArchiveFile(tt.path)
			assert.Equal(tt.want, got)
			assert.NoError(err)
		})
	}
}
