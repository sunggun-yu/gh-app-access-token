package installation

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

// TestGenerateAccessToken_InvalidPrivateKey verifies that an error is returned when
// a malformed private key is provided.
func TestGenerateAccessToken_InvalidPrivateKey(t *testing.T) {
	ctx := context.Background()
	_, err := GenerateAccessToken(ctx, 12345, 67890, []byte("not-a-valid-key"))
	if err == nil {
		t.Error("expected error for invalid private key, got nil")
	}
}

// TestGenerateAccessToken_EmptyPrivateKey verifies that an error is returned when
// an empty private key is provided.
func TestGenerateAccessToken_EmptyPrivateKey(t *testing.T) {
	ctx := context.Background()
	_, err := GenerateAccessToken(ctx, 12345, 67890, []byte(""))
	if err == nil {
		t.Error("expected error for empty private key, got nil")
	}
}

// TestRevokeAccessToken_Success verifies that revoking a valid access token succeeds
// by mocking the GitHub API to return 204 No Content.
func TestRevokeAccessToken_Success(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/installation/token", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE, got %s", r.Method)
		}
		if r.Header.Get("Authorization") == "" {
			t.Error("expected Authorization header")
		}
		w.WriteHeader(http.StatusNoContent)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "test-token"})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	client.BaseURL, _ = client.BaseURL.Parse(server.URL + "/")

	_, err := client.Apps.RevokeInstallationToken(ctx)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

// TestRevokeAccessToken_Failure verifies that revoking an invalid/unauthorized token
// returns an error by mocking the GitHub API to return 401 Unauthorized.
func TestRevokeAccessToken_Failure(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/installation/token", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"message": "Bad credentials"})
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "invalid-token"})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	client.BaseURL, _ = client.BaseURL.Parse(server.URL + "/")

	_, err := client.Apps.RevokeInstallationToken(ctx)
	if err == nil {
		t.Error("expected error for unauthorized token, got nil")
	}
}
