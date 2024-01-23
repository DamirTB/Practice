package assert
import (
	"strings" // New import
	"testing"
)

func StringContains(t *testing.T, actual, expectedSubstring string) {
	t.Helper()
	if !strings.Contains(actual, expectedSubstring) {
	t.Errorf("got: %q; expected to contain: %q", actual, expectedSubstring)
	}
}

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}