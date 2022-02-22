package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jiro4989/relma/lock"
	"github.com/jiro4989/relma/prompt"
	"github.com/jiro4989/relma/releases"
	"github.com/spf13/cobra"
)

type CommandLineUpgradeParam struct {
	Upgrade   bool
	Yes       bool
	OwnerRepo string
}

func init() {
	commandUpgrade.Flags().BoolVarP(&commandLineUpgradeParam.Yes, "yes", "y", false, "yes")

	rootCmd.AddCommand(commandUpgrade)
}

var commandLineUpgradeParam CmdUpgradeParam

var commandUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade installed GitHub Releases",
	RunE: func(cmd *cobra.Command, args []string) error {
		return lock.TransactionLock(func() error {
			a, err := NewApp()
			if err != nil {
				return err
			}
			if 0 < len(args) {
				commandLineUpgradeParam.OwnerRepo = args[0]
			}
			return a.CmdUpgrade(&commandLineUpgradeParam)
		})
	},
}

type CmdUpgradeParam struct {
	OwnerRepo string
	Yes       bool
}

func (a *App) CmdUpgrade(p *CmdUpgradeParam) error {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return err
	}

	return a.cmdUpgrade(rels, p)
}

func (a *App) cmdUpgrade(rels releases.Releases, p *CmdUpgradeParam) error {
	rels, err := searchReleaseOrDefault(rels, p.OwnerRepo)
	if err != nil {
		return err
	}

	targets := rels.UpgradableReleases()
	if len(targets) < 1 {
		fmt.Println("no upgradable releases")
		return nil
	}

	if !p.Yes {
		if yes, err := prompt.PromptYesNo("upgrade? [yes/no]"); err != nil {
			return err
		} else if !yes {
			fmt.Println("upgrade was canceled")
			return nil
		}
	}

	for _, rel := range targets {
		url := strings.ReplaceAll(rel.URL, rel.Version, rel.LatestVersion)
		err := a.CmdInstall(&CmdInstallParam{URL: url})
		if err != nil {
			return err
		}

		Sleep()
	}
	fmt.Println("upgrade successful")

	return nil
}

func searchReleaseOrDefault(rels releases.Releases, ownerRepo string) (releases.Releases, error) {
	if len(rels) < 1 {
		return nil, errors.New("installed releases don't exist")
	}

	if ownerRepo != "" {
		var err error
		rels, err = searchRelease(rels, ownerRepo)
		if err != nil {
			return nil, err
		}
		if len(rels) < 1 {
			msg := fmt.Sprintf("%s was not installed", ownerRepo)
			return nil, errors.New(msg)
		}
	}

	return rels, nil
}

func searchRelease(rels releases.Releases, ownerRepo string) (releases.Releases, error) {
	var retRels releases.Releases
	for _, rel := range rels {
		if ok, err := rel.EqualRepo(ownerRepo); err != nil {
			return nil, err
		} else if !ok {
			continue
		}
		retRels = append(retRels, rel)
		return retRels, nil
	}
	return nil, nil
}
