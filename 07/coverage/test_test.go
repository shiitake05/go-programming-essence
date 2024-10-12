// テストコード
package coverage

import (
	"testing"
)

func TestStringDistance(t *testing.T) {
	got := StringDistance("foo", "foh")
	want := 1
	if got != want {
		t.Fatalf("expected: %v, got %v:", want, got)
	}

	got = StringDistance("foo", "foobar")
	want = -1
	if got != want {
		t.Fatalf("expected: %v, got %v:", want, got)
	}
}
