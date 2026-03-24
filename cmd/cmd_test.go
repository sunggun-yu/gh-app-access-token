package cmd

import (
	"bytes"
	"testing"
)

// TestRootCmd_Execute verifies that the root command executes without error (displays help).
func TestRootCmd_Execute(t *testing.T) {
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("root command execution failed: %v", err)
	}
}

// TestRootCmd_HasSubcommands verifies that the root command has both "generate" and "revoke" subcommands registered.
func TestRootCmd_HasSubcommands(t *testing.T) {
	cmds := rootCmd.Commands()
	if len(cmds) == 0 {
		t.Error("expected root command to have subcommands")
	}
	names := make(map[string]bool)
	for _, c := range cmds {
		names[c.Name()] = true
	}
	if !names["generate"] {
		t.Error("missing 'generate' subcommand")
	}
	if !names["revoke"] {
		t.Error("missing 'revoke' subcommand")
	}
}

// TestSetVersion verifies that SetVersion correctly sets the version string on the root command.
func TestSetVersion(t *testing.T) {
	SetVersion("1.2.3")
	if rootCmd.Version != "1.2.3" {
		t.Errorf("expected version 1.2.3, got %s", rootCmd.Version)
	}
}

// TestGenerateCmd_RequiresFlags verifies that the generate command fails when required flags
// (--app-id, --installation-id, --private-key) are not provided.
func TestGenerateCmd_RequiresFlags(t *testing.T) {
	cmd := generateCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	if err == nil {
		t.Error("expected error when required flags are missing")
	}
}

// TestGenerateCmd_Aliases verifies that the generate command has "create" and "gen" as aliases.
func TestGenerateCmd_Aliases(t *testing.T) {
	cmd := generateCmd()
	expected := []string{"create", "gen"}
	if len(cmd.Aliases) != len(expected) {
		t.Errorf("expected %d aliases, got %d", len(expected), len(cmd.Aliases))
		return
	}
	for i, a := range cmd.Aliases {
		if a != expected[i] {
			t.Errorf("alias[%d] = %q, want %q", i, a, expected[i])
		}
	}
}

// TestGenerateCmd_RejectsPositionalArgs verifies that the generate command rejects
// unexpected positional arguments (cobra.NoArgs).
func TestGenerateCmd_RejectsPositionalArgs(t *testing.T) {
	cmd := generateCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"--app-id", "1", "--installation-id", "2", "--private-key", "dummy", "extra-arg"})
	err := cmd.Execute()
	if err == nil {
		t.Error("expected error when positional args are passed to generate")
	}
}

// TestRevokeCmd_RequiresArg verifies that the revoke command fails when no access token argument is provided.
func TestRevokeCmd_RequiresArg(t *testing.T) {
	cmd := revokeCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{})
	err := cmd.Execute()
	if err == nil {
		t.Error("expected error when no argument is provided to revoke")
	}
}

// TestRevokeCmd_TooManyArgs verifies that the revoke command fails when more than one argument is provided.
func TestRevokeCmd_TooManyArgs(t *testing.T) {
	cmd := revokeCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"token1", "token2"})
	err := cmd.Execute()
	if err == nil {
		t.Error("expected error when too many args are passed to revoke")
	}
}

// TestRevokeCmd_Aliases verifies that the revoke command has "del" as an alias.
func TestRevokeCmd_Aliases(t *testing.T) {
	cmd := revokeCmd()
	if len(cmd.Aliases) != 1 || cmd.Aliases[0] != "del" {
		t.Errorf("expected aliases [del], got %v", cmd.Aliases)
	}
}
