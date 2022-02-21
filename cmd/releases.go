package cmd

import (
	"errors"
	"fmt"

	"github.com/jiro4989/relma/lock"
	"github.com/jiro4989/relma/releases"
	"github.com/spf13/cobra"
)

func init() {
	commandReleases.AddCommand(commandReleasesLock)
	commandReleases.AddCommand(commandReleasesUnlock)
	rootCmd.AddCommand(commandReleases)
}

var commandReleases = &cobra.Command{
	Use:   "releases",
	Short: "lock or unlock releases",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var commandReleasesLock = &cobra.Command{
	Use:   "lock",
	Short: "lock specific version",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a <owner/repo>")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return lock.TransactionLock(func() error {
			a, err := NewApp()
			if err != nil {
				return err
			}
			_, err = a.cmdReleasesLockUnlock(args[0], true, "lock")
			return err
		})
	},
}

var commandReleasesUnlock = &cobra.Command{
	Use:   "unlock",
	Short: "unlock specific version",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a <owner/repo>")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return lock.TransactionLock(func() error {
			a, err := NewApp()
			if err != nil {
				return err
			}
			_, err = a.cmdReleasesLockUnlock(args[0], false, "unlock")
			return err
		})
	},
}

func (a *App) cmdReleasesLockUnlock(ownerRepo string, lock bool, ops string) (releases.Releases, error) {
	rels, err := a.Config.ReadReleasesFile()
	if err != nil {
		return nil, err
	}

	if err := rels.Lock(ownerRepo, lock); err != nil {
		return nil, err
	}

	if err := a.SaveReleases(rels); err != nil {
		return nil, err
	}

	msg := fmt.Sprintf("%sed '%s'", ops, ownerRepo)
	fmt.Println(msg)

	return rels, nil
}
