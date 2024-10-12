// テストコード
package test

import (
	"testing"
)

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	// want := 2
	want := 1
	if got != want {
		t.Fatalf("expected: %v, got %v:", want, got)
	}
}
