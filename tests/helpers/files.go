package helpers

import (
	"os"
	"testing"
)

func CreateTempFile(t testing.TB, name, contents string) (os.File, func()) {
	t.Helper()

	tempFile, err := os.CreateTemp("", name)
	AssertNoError(t, err)

	tempFile.Write([]byte(contents))
	closeFile := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return *tempFile, closeFile
}
