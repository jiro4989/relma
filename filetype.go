package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
)

// IsExecutableFile returns true if `f` is executable.
func IsExecutableFile(f os.FileInfo, path string) (bool, error) {
	mode := f.Mode()
	if !mode.IsRegular() {
		return false, nil
	}

	if mode&0111 != 0 {
		return true, nil
	}

	ext := filepath.Ext(path)
	switch strings.ToLower(ext) {
	case ".bat", ".cmd":
		return true, nil
	}

	typ, err := filetype.MatchFile(path)
	if err != nil {
		return false, err
	}
	switch typ.Extension {
	case "elf", "exe":
		return true, nil
	}

	return false, nil
}

// IsArchiveFile returns true if `f` is archive file (zip or gz).
func IsArchiveFile(path string) (bool, error) {
	typ, err := filetype.MatchFile(path)
	if err != nil {
		return false, err
	}
	switch typ.Extension {
	case "gz", "zip":
		return true, nil
	}

	return false, nil
}
