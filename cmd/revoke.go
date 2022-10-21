/*
Copyright © 2022 Sunggun Yu <sunggun.dev@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sunggun-yu/gh-app-access-token-cli/internal/utils"
	"github.com/sunggun-yu/gh-app-access-token-cli/pkg/installation"
)

// revokeCmd func represents the revoke command
func revokeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "revoke [access token string to revoke]",
		Short:   "Revoke an access token",
		Aliases: []string{"del"},
		Example: `  # revoke token in argument
  gh-app-access-token-cli revoke [access token string]

  # revoke the token passed into stdin
  cat [access-token-file] | gh-app-access-token-cli revoke -
  
  # revoke the token passed into stdin
  echo "access-token-value" | gh-app-access-token-cli revoke -`,
		Args: cobra.MatchAll(
			cobra.MinimumNArgs(1),
			cobra.MaximumNArgs(1),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			token := args[0]
			// read token from stdin if arg is `-`
			if args[0] == "-" {
				token = utils.ReadInOrStdin(cmd)
			}
			err := installation.RevokeAccessToken(cmd.Context(), token)
			return err
		},
	}
	return cmd
}

func init() {
	rootCmd.AddCommand(revokeCmd())
}
