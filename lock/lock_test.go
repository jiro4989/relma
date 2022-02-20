package lock

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	tmpDir = ".tmp"
)

func TestMain(m *testing.M) {
	testBefore()
	exitCode := m.Run()
	testAfter()
	os.Exit(exitCode)
}

func testBefore() {
	os.RemoveAll(tmpDir)

	if err := os.Mkdir(tmpDir, os.ModePerm); err != nil {
		panic(err)
	}
}

func testAfter() {
	if err := os.RemoveAll(tmpDir); err != nil {
		panic(err)
	}
}

func TestLockAndUnlock(t *testing.T) {
	assert := assert.New(t)
	dir := filepath.Join(tmpDir, "lock")
	assert.NoError(lock(dir), "ok: successful creating a lock directory")
	assert.Error(lock(dir), "ok: raises error when a lock directory was existed")
	assert.NoError(unlock(dir), "ok: successful removing a lock directory")

	_, err := os.Stat(dir)
	assert.True(os.IsNotExist(err), "ok: a lock directory was removed")
}

func TestTransactionLock(t *testing.T) {
	assert := assert.New(t)

	var err error
	err = TransactionLock(func() error {
		return nil
	})
	assert.NoError(err)

	_, err = os.Stat(lockDir)
	assert.True(os.IsNotExist(err), "ok: a lock directory was removed")

	// 先にロックディレクトリを作ることでエラーを起こさせる
	assert.NoError(lock(lockDir))
	err = TransactionLock(func() error {
		return errors.New("error")
	})
	assert.Error(err, "ok: locked")
	_, err = os.Stat(lockDir)
	assert.False(os.IsNotExist(err), "ok: a lock directory was existed")
	assert.NoError(Unlock(), "ok: successful unlock")
}
