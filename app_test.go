package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadReleasesFile(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		want    Releases
		wantErr bool
	}{
		{
			desc: "ok: releases.json exists",
			path: filepath.Join(testDir, "releases.json"),
			want: Releases{
				{
					URL: "https://example.com",
				},
				{
					URL: "https://example2.com",
				},
			},
			wantErr: false,
		},
		{
			desc:    "ok: releases.json doesn't exist",
			path:    "not_found.json",
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := readReleasesFile(tt.path)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}
