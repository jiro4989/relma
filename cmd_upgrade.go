package main

type CmdUpgradeParam struct {
}

func (a *App) CmdUpgrade(p *CmdUpgradeParam) error {
	pkgDir := a.Config.ReleasesDir()
	pkgFiles, err := readReleasesFiles(pkgDir)
	if err != nil {
		return err
	}
	for _, pkgFile := range pkgFiles {
		_, err := readReleases(pkgFile)
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
