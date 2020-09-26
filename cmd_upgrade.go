package main

type CmdUpgradeParam struct {
}

func (a *App) CmdUpgrade(p *CmdUpgradeParam) error {
	releasesDir := a.Config.ReleasesDir()
	releasesFiles, err := readReleasesFiles(releasesDir)
	if err != nil {
		return err
	}
	for _, releasesFile := range releasesFiles {
		_, err := readReleases(releasesFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func readReleasesFiles(dir string) ([]string, error) {
	return nil, nil
}

func readReleases(path string) (*Releases, error) {
	return nil, nil
}
