package installation

import (
	"context"
	"net/http"

	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

// GenerateAccessToken ...
func GenerateAccessToken(ctx context.Context, appID, installationID int64, privateKey []byte) (string, error) {
	// Shared transport to reuse TCP connections.
	tr := http.DefaultTransport

	// Wrap the shared transport for use with the app ID 1 authenticating with installation ID 99.
	itr, err := ghinstallation.New(tr, appID, installationID, privateKey)
	if err != nil {
		return "", err
	}
	accessToken, err := itr.Token(ctx)
	return accessToken, err
}

// RevokeAccessToken ...
func RevokeAccessToken(ctx context.Context, accessToken string) error {
	// github client auth with input access token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// revoke the token
	_, err := client.Apps.RevokeInstallationToken(ctx)
	return err
}
