/*
Copyright © 2022 Sunggun Yu <sunggun.dev@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/sunggun-yu/gh-app-access-token/internal/utils"
	"github.com/sunggun-yu/gh-app-access-token/pkg/installation"
)

// flags struct for show command
type generateCmdFlags struct {
	appID          int64
	installationID int64
	privateKeyFile string
}

// generateCmd represents the generate command
func generateCmd() *cobra.Command {
	var flags generateCmdFlags

	cmd := &cobra.Command{
		Use:     "generate",
		Short:   "Generate the Github App Installation access token",
		Aliases: []string{"create", "gen"},
		Args:    cobra.NoArgs,
		Example: `  # generate the Github App access token
  gh-app-access-token generate \
    --app-id [app-id] \
    --installation-id [installation-id] \
    --private-key [private-key-file-path]
  
	# generate the Github App access token with text in private key file passed into stdin
  cat [private-key-file-path] | gh-app-access-token generate \
    --app-id [app-id] \
    --installation-id [installation-id] \
    --private-key -
  
  # generate the Github App access token with private key text passed into stdin
  echo "private-key-text" | gh-app-access-token generate \
    --app-id [app-id] \
    --installation-id [installation-id] \
    --private-key -`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var privateKey []byte
			if flags.privateKeyFile == "-" {
				// read private key from stdin if flag value is `-`
				privateKey = []byte(utils.ReadInOrStdin(cmd))
			} else {
				privateKey, _ = os.ReadFile(flags.privateKeyFile)
			}

			token, err := installation.GenerateAccessToken(cmd.Context(), flags.appID, flags.installationID, privateKey)
			if err != nil {
				return err
			}
			fmt.Printf("%s", strings.TrimSpace(token))
			return nil
		},
	}

	cmd.Flags().Int64VarP(&flags.appID, "app-id", "a", 0, "The unique identifier of the Github App")
	cmd.Flags().Int64VarP(&flags.installationID, "installation-id", "i", 0, "The unique identifier of the installation")
	cmd.Flags().StringVarP(&flags.privateKeyFile, "private-key", "f", "", "The private key file path of Github App")

	cmd.MarkFlagRequired("app-id")
	cmd.MarkFlagRequired("installation-id")
	cmd.MarkFlagRequired("private-key")
	return cmd
}

func init() {
	rootCmd.AddCommand(generateCmd())
}
