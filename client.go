package patreon

import "golang.org/x/oauth2"

const (
	ScopeUsers       = "users"
	ScopePledgesToMe = "pledges-to-me"
	ScopeMyCampaign  = "my-campaign"
)

func NewClient(clientID, clientSecret, redirectURL string, scopes ...string) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.patreon.com/oauth2/authorize",
			TokenURL: "https://api.patreon.com/oauth2/token",
		},
		RedirectURL: redirectURL,
		Scopes:      scopes,
	}
}
