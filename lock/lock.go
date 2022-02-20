package lock

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jiro4989/relma/logger"
)

var lockDir = filepath.Join(os.TempDir(), "relma.lock")

func lock(dir string) error {
	return os.Mkdir(dir, os.ModePerm)
}

func unlock(dir string) error {
	return os.Remove(dir)
}

func Unlock() error {
	return unlock(lockDir)
}

func TransactionLock(f func() error) error {
	// ディレクトリの作成に失敗したらロック済みなので処理を中断する
	if err := lock(lockDir); err != nil {
		msg := fmt.Sprintf("relma is running. remove %s directory with 'relma unlock' if relma is not running now", lockDir)
		logger.Error(msg)
		return err
	}

	// エラーの有無に関わらずアンロックする
	err := f()

	// アンロックでエラーが出たら無視
	unlock(lockDir)

	return err
}
