package helpers

import (
	"reflect"
	"testing"
)

func AssertEqual[T comparable](t testing.TB, actual, expected T) {
	t.Helper()

	if actual != expected {
		t.Errorf("Got '%v', expected '%v'", actual, expected)
	}
}

func AssertSliceEqual[T any](t testing.TB, actual, expected []T) {
	t.Helper()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Got %+v, expected %+v", actual, expected)
	}
}

func AssertError(t testing.TB, err error) {
	t.Helper()

	if err == nil {
		t.Error("Expected an error, but no error was returned.")
	}
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("expected no error, but the following error was returned: %v", err)
	}
}
