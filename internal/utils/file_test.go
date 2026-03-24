package utils

import (
	"os"
	"path/filepath"
	"testing"
)

// TestFilePath_Tilde verifies that tilde (~) prefix is expanded to the user's home directory.
func TestFilePath_Tilde(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	got := FilePath("~/some/path")
	want := filepath.Join(home, "some/path")
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestFilePath_HomeEnv verifies that $HOME environment variable is expanded to the user's home directory.
func TestFilePath_HomeEnv(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		t.Fatal(err)
	}
	got := FilePath("$HOME/some/path")
	want := filepath.Join(home, "some/path")
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestFilePath_PWD verifies that $PWD is expanded to the current working directory.
func TestFilePath_PWD(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	got := FilePath("$PWD/file.txt")
	want := filepath.Join(pwd, "file.txt")
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// TestFilePath_PlainPath verifies that an absolute path is returned unchanged.
func TestFilePath_PlainPath(t *testing.T) {
	got := FilePath("/absolute/path/to/file")
	if got != "/absolute/path/to/file" {
		t.Errorf("got %q, want /absolute/path/to/file", got)
	}
}

// TestFilePath_RelativePath verifies that a relative path is returned unchanged.
func TestFilePath_RelativePath(t *testing.T) {
	got := FilePath("relative/path")
	if got != "relative/path" {
		t.Errorf("got %q, want relative/path", got)
	}
}

// TestFilePath_Empty verifies that an empty string input returns an empty string.
func TestFilePath_Empty(t *testing.T) {
	got := FilePath("")
	if got != "" {
		t.Errorf("got %q, want empty string", got)
	}
}
