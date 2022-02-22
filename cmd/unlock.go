package cmd

import (
	"fmt"

	"github.com/jiro4989/relma/lock"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commandUnlock)
}

var commandUnlock = &cobra.Command{
	Use:   "unlock",
	Short: "Unlock a lock file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := lock.Unlock(); err != nil {
			return err
		}
		fmt.Println("unlock successfully")
		return nil
	},
}
