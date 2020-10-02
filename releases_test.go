package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelease_FormatSimpleInformation(t *testing.T) {
	tests := []struct {
		desc string
		rel  Release
		want string
	}{
		{
			desc: "ok: generate string",
			rel: Release{
				Owner:   "jiro4989",
				Repo:    "nimjson",
				Version: "v1.2.6",
			},
			want: "jiro4989/nimjson v1.2.6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got := tt.rel.FormatSimpleInformation()
			assert.Equal(tt.want, got)
		})
	}
}

func TestInstalledFiles_FixPath(t *testing.T) {
	tests := []struct {
		desc  string
		files InstalledFiles
		src   string
		dest  string
		want  InstalledFiles
	}{
		{
			desc: "ok: fix path",
			files: InstalledFiles{
				{
					Src:  "/home/foobar/sample",
					Dest: "/home/foobar/bin/sample",
				},
			},
			src:  "/home/foobar",
			dest: "/home/foobar/bin",
			want: InstalledFiles{
				{
					Src:  "sample",
					Dest: "sample",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			tt.files.FixPath(tt.src, tt.dest)
			assert.Equal(tt.want, tt.files)
		})
	}
}
