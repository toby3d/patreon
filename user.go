package patreon

import (
	"path"
	"strconv"
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
)

func (client *Client) GetCurrentUser() (*User, error) {
	requestURI := defaultURI
	requestURI.Path = path.Join("api", "oauth2", "api", "current_user")

	resp, err := client.get(requestURI.String())
	if err != nil {
		return nil, err
	}

	var dst User
	err = json.Unmarshal(resp, &dst)
	return &dst, err
}

func (client *Client) GetUser(userID int, includes ...string) (*Campaign, error) {
	requestURI := defaultURI
	requestURI.Path = path.Join("api", "user", strconv.Itoa(userID))

	if len(includes) > 0 {
		query := requestURI.Query()
		query.Add("includes", strings.Join(includes, ","))
		requestURI.RawQuery = query.Encode()
	}

	resp, err := client.get(requestURI.String())
	if err != nil {
		return nil, err
	}

	var dst Campaign
	err = json.Unmarshal(resp, &dst)
	return &dst, err
}
