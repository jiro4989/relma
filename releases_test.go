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
