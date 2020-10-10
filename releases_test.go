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

func TestUnset(t *testing.T) {
	tests := []struct {
		desc string
		rels Releases
		i    int
		want Releases
	}{
		{
			desc: "ok: unset 0",
			rels: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			i: 0,
			want: Releases{
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
		},
		{
			desc: "ok: unset 1",
			rels: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			i: 1,
			want: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
			},
		},
		{
			desc: "ng: illegal index",
			rels: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			i: 99,
			want: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got := unset(tt.rels, tt.i)
			assert.Equal(tt.want, got)
		})
	}
}

func TestRemoveRelease(t *testing.T) {
	tests := []struct {
		desc string
		rels Releases
		rel  *Release
		want Releases
	}{
		{
			desc: "ok: remove 0",
			rels: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			rel: &Release{
				Owner: "jiro4989",
				Repo:  "monit",
			},
			want: Releases{
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
		},
		{
			desc: "ok: remove 1",
			rels: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			rel: &Release{
				Owner: "jiro4989",
				Repo:  "textimg",
			},
			want: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
			},
		},
		{
			desc: "ok: no remove",
			rels: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
			rel: &Release{
				Owner: "jiro4989",
				Repo:  "sushi",
			},
			want: Releases{
				{
					Owner: "jiro4989",
					Repo:  "monit",
				},
				{
					Owner: "jiro4989",
					Repo:  "textimg",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got := RemoveRelease(tt.rels, tt.rel)
			assert.Equal(tt.want, got)
		})
	}
}

func TestRelease_EqualRelease(t *testing.T) {
	tests := []struct {
		desc string
		a, b Release
		want bool
	}{
		{
			desc: "ok: match",
			a: Release{
				Owner: "jiro4989",
				Repo:  "textimg",
			},
			b: Release{
				Owner: "JIRO4989",
				Repo:  "textimg",
			},
			want: true,
		},
		{
			desc: "ok: unmatch",
			a: Release{
				Owner: "jiro4989",
				Repo:  "textimg",
			},
			b: Release{
				Owner: "jiro4989",
				Repo:  "monit",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got := tt.a.EqualRelease(&tt.b)
			assert.Equal(tt.want, got)
		})
	}
}
