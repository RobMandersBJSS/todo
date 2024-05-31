package helpers

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func CreateTempFile(t testing.TB, name, contents string) (os.File, func()) {
	t.Helper()

	dir, removeDir := CreateTempDirectory(t)

	tempFile, err := os.CreateTemp(dir, name)
	AssertNoError(t, err)

	tempFile.Write([]byte(contents))
	removeFile := func() {
		tempFile.Close()
		removeDir()
	}

	return *tempFile, removeFile
}

func CreateTempDirectory(t testing.TB) (string, func()) {
	t.Helper()

	dir, err := os.MkdirTemp("", "test")
	AssertNoError(t, err)

	removeDir := func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatalf("Error removing temporary directory %q: %v", dir, err)
		}
	}

	return dir, removeDir
}

func GetRootDir() string {
	_, basePath, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(basePath), "../..")

	return rootDir
}
