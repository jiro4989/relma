package main

import (
	"os"
	"path/filepath"
	"testing"
)

var (
	testDir       = filepath.Join(".", "testdata")
	testOutputDir = filepath.Join(testDir, "out")
)

func TestMain(m *testing.M) {
	testBefore()
	exitCode := m.Run()
	testAfter()
	os.Exit(exitCode)
}

func testBefore() {
	os.RemoveAll(testOutputDir)

	if err := os.Mkdir(testOutputDir, os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.Mkdir(filepath.Join(testOutputDir, "bin"), os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.Mkdir(filepath.Join(testOutputDir, "releases"), os.ModePerm); err != nil {
		panic(err)
	}
}

func testAfter() {
	if err := os.RemoveAll(testOutputDir); err != nil {
		panic(err)
	}
}
