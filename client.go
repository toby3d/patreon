package patreon

import (
	"context"
	"path"

	"golang.org/x/oauth2"
)

type Client struct {
	Config *oauth2.Config
	Token  *oauth2.Token
}

func NewClient(clientID, clientSecret, redirectURL string, scopes ...string) *Client {
	tokenURL := defaultAPIURI
	tokenURL.Path = path.Join("oauth2", "token")

	authURL := defaultURI
	authURL.Path = path.Join("oauth2", "authorize")

	cfg := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL.String(),
			TokenURL: tokenURL.String(),
		},
		RedirectURL: redirectURL,
		Scopes:      scopes,
	}

	return &Client{Config: &cfg}
}

func (client *Client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return client.Config.AuthCodeURL(state, opts...)
}

func (client *Client) Exchange(ctx context.Context, code string) error {
	var err error
	client.Token, err = client.Config.Exchange(ctx, code)
	return err
}
