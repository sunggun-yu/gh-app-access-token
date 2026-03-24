package utils

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

// TestReadInOrStdin verifies reading input from stdin with various formats:
// - simple input is returned as-is
// - trailing newlines and surrounding whitespace are trimmed
// - multiline input preserves internal newlines
// - empty input returns an empty string
func TestReadInOrStdin(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "simple input",
			input: "hello world",
			want:  "hello world",
		},
		{
			name:  "trailing newline",
			input: "hello world\n",
			want:  "hello world",
		},
		{
			name:  "leading and trailing whitespace",
			input: "  hello world  \n",
			want:  "hello world",
		},
		{
			name:  "multiline input",
			input: "line1\nline2\nline3",
			want:  "line1\nline2\nline3",
		},
		{
			name:  "empty input",
			input: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.SetIn(bytes.NewBufferString(tt.input))
			got := ReadInOrStdin(cmd)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
