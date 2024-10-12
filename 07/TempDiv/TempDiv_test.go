// このTempDir()は何度も呼び出すことができ、TestCreateProfileが終了すると自動でテンポラリディレクトリが削除される
package main

import (
	"path/filepath"
	"testing"
)

func CreateProfile(filename string) (bool, error) {
	return true, nil
}

func TestCreateProfile(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "profile.json")
	got, err := CreateProfile(filename)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := true
	if got != want {
		t.Fatalf("expected: %v, got %v:", want, got)
	}
}
