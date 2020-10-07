package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestCmdInstall(t *testing.T) {
// 	tests := []struct {
// 		desc      string
// 		app       App
// 		url       string
// 		want      Releases
// 		wantCount int
// 		wantErr   bool
// 	}{
// 		{
// 			desc: "ok: installing",
// 			app: App{
// 				Config: Config{
// 					RelmaRoot: testOutputDir,
// 				},
// 			},
// 			url: "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
// 			want: Releases{
// 				{
// 					URL:           "https://github.com/jiro4989/nimjson/releases/download/v1.2.6/nimjson_linux.tar.gz",
// 					Owner:         "jiro4989",
// 					Repo:          "nimjson",
// 					Version:       "v1.2.6",
// 					AssetFileName: "nimjson_linux.tar.gz",
// 					InstalledFiles: InstalledFiles{
// 						{
// 							Src:  filepath.Join("bin", "nimjson"),
// 							Dest: "nimjson",
// 						},
// 					},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.desc, func(t *testing.T) {
// 			assert := assert.New(t)
//
// 			p := filepath.Join(testOutputDir, "releases.json")
// 			os.Remove(p)
//
// 			err := tt.app.CmdInstall(tt.url)
// 			if tt.wantErr {
// 				assert.Error(err)
// 				return
// 			}
// 			assert.NoError(err)
//
// 			b, err := ioutil.ReadFile(p)
// 			assert.NoError(err)
//
// 			var rels Releases
// 			err = json.Unmarshal(b, &rels)
// 			assert.NoError(err)
// 			assert.Equal(len(tt.want), len(rels))
//
// 			for i, want := range tt.want {
// 				rel := rels[i]
// 				assert.Equal(want.URL, rel.URL)
// 				assert.Equal(want.Owner, rel.Owner)
// 				assert.Equal(want.Repo, rel.Repo)
// 				assert.Equal(want.Version, rel.Version)
// 				assert.Equal(want.AssetFileName, rel.AssetFileName)
// 				assert.Equal(want.InstalledFiles, rel.InstalledFiles)
// 			}
// 		})
// 	}
// }

func TestUninstallableRelease(t *testing.T) {
	tests := []struct {
		desc      string
		rels      Releases
		ownerRepo string
		want      *Release
		wantErr   bool
	}{
		{
			desc: "ok: uninstallable",
			rels: Releases{
				{
					Owner: "JIRO4989",
					Repo:  "sushi",
				},
				{
					Owner: "JIRO4989",
					Repo:  "TEXTIMG",
				},
			},
			want: &Release{
				Owner: "JIRO4989",
				Repo:  "TEXTIMG",
			},
			ownerRepo: "jiro4989/textimg",
			wantErr:   false,
		},
		{
			desc: "ng: un match",
			want: &Release{
				Owner: "jiro4989",
				Repo:  "textimg",
			},
			ownerRepo: "jiro4989/monit",
			wantErr:   true,
		},
		{
			desc: "ng: illegal ownerRepo",
			want: &Release{
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
