package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/jiro4989/relma/logger"
	"github.com/spf13/cobra"
)

var Error = logger.Error

func init() {
	rootCmd.AddCommand(commandInit)
}

var commandInit = &cobra.Command{
	Use:   "init",
	Short: "Setup relma environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		// init での初期化で初めて設定ファイルを作成するため
		// init 処理では NewApp 関数は呼び出さない
		var a App
		if err := a.SetUserEnv(); err != nil {
			logger.Error(err)
			return err
		}
		return a.CmdInit()
	},
}

func (a *App) CmdInit() error {
	_, err := a.CreateConfigDir()
	if err != nil {
		logger.Error(err)
		return err
	}

	conf := a.DefaultConfig()

	_, err = a.CreateConfigFile(conf)
	if err != nil {
		logger.Error(err)
		return err
	}

	paths := []string{
		conf.RelmaRoot,
		conf.BinDir(),
		conf.ReleasesDir(),
	}
	for _, path := range paths {
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			continue
		}

		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Error(err)
			return err
		}
	}

	fmt.Println("initialize successful")

	return nil
}

func Sleep() {
	time.Sleep(1 * time.Second)
}
