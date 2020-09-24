package main

type CmdUpgradeParam struct {
}

func (a *App) CmdUpgrade(p *CmdUpgradeParam) error {
	pkgDir := a.Config.PackageDir()
	pkgFiles, err := readPackageFiles(pkgDir)
	if err != nil {
		return err
	}
	for _, pkgFile := range pkgFiles {
		_, err := readPackage(pkgFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func readPackageFiles(dir string) ([]string, error) {
	return nil, nil
}

func readPackage(path string) (*Package, error) {
	return nil, nil
}
