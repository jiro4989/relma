package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmdUpgrade(t *testing.T) {
	// tests := []struct {
	// 	desc      string
	// 	app       App
	// 	url       string
	// 	want      Releases
	// 	wantCount int
	// 	wantErr   bool
	// }{
	// 	{
	// 		desc: "ok: installing",
	// 		app: App{
	// 			Config: Config{
	// 				RelmaRoot: testOutputDir,
	// 			},
	// 		},
	// 		url: "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
	// 		want: Releases{
	// 			{
	// 				URL:           "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
	// 				Owner:         "jiro4989",
	// 				Repo:          "nimjson",
	// 				Version:       "v1.2.6",
	// 				AssetFileName: "nimjson_linux.tar.gz",
	// 				InstalledFiles: InstalledFiles{
	// 					{
	// 						Src:  filepath.Join("bin", "nimjson"),
	// 						Dest: "nimjson",
	// 					},
	// 				},
	// 			},
	// 		},
	// 		wantErr: false,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.desc, func(t *testing.T) {
	// 		assert := assert.New(t)
	//
	// 		p := filepath.Join(testOutputDir, "releases.json")
	// 		os.Remove(p)
	//
	// 		err := tt.app.CmdInstall(tt.url)
	// 		if tt.wantErr {
	// 			assert.Error(err)
	// 			return
	// 		}
	// 		assert.NoError(err)
	//
	// 		b, err := ioutil.ReadFile(p)
	// 		assert.NoError(err)
	//
	// 		var rels Releases
	// 		err = json.Unmarshal(b, &rels)
	// 		assert.NoError(err)
	// 		assert.Equal(len(tt.want), len(rels))
	//
	// 		for i, want := range tt.want {
	// 			rel := rels[i]
	// 			assert.Equal(want.URL, rel.URL)
	// 			assert.Equal(want.Owner, rel.Owner)
	// 			assert.Equal(want.Repo, rel.Repo)
	// 			assert.Equal(want.Version, rel.Version)
	// 			assert.Equal(want.AssetFileName, rel.AssetFileName)
	// 			assert.Equal(want.InstalledFiles, rel.InstalledFiles)
	// 		}
	// 	})
	// }
}

func TestSearchRelease(t *testing.T) {
	tests := []struct {
		desc      string
		rels      Releases
		ownerRepo string
		want      Releases
		wantErr   bool
	}{
		{
			desc: "ok: found releases",
			rels: Releases{
				{
					Owner: "Jiro",
					Repo:  "Test",
				},
				{
					Owner: "Jiro4989",
					Repo:  "TestA",
				},
				{
					Owner: "Jiro",
					Repo:  "TestA",
				},
				{
					Owner: "Jiro4989",
					Repo:  "Test",
				},
			},
			ownerRepo: "JIRO4989/TEST",
			want: Releases{
				{
					Owner: "Jiro4989",
					Repo:  "Test",
				},
			},
			wantErr: false,
		},
		{
			desc:      "ok: releases is empty",
			rels:      Releases{},
			ownerRepo: "",
			want:      nil,
			wantErr:   false,
		},
		{
			desc: "ng: ownerRepo is illegal",
			rels: Releases{
				{
					Owner: "Jiro4989",
					Repo:  "Test",
				},
			},
			ownerRepo: "JIRO4989TEST",
			want:      nil,
			wantErr:   true,
		},
		{
			desc: "ng: ownerRepo is empty",
			rels: Releases{
				{
					Owner: "Jiro4989",
					Repo:  "Test",
				},
			},
			ownerRepo: "",
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := searchRelease(tt.rels, tt.ownerRepo)
			if tt.wantErr {
				assert.Error(err)
				return
			}
			assert.NoError(err)
			assert.Equal(tt.want, got)
		})
	}
}
