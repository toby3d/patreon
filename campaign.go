package patreon

import (
	"path"
	"strconv"
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
)

// Can use next includes: rewards, creator, goals, pledges
func (client *Client) GetCurrentCampaign(includes ...string) (*Campaign, error) {
	requestURI := defaultURI
	requestURI.Path = path.Join("api", "oauth2", "api", "current_user", "campaigns")

	if len(includes) > 0 {
		query := requestURI.Query()
		query.Add("includes", strings.Join(includes, ","))
		requestURI.RawQuery = query.Encode()
	}

	resp, err := client.get(requestURI.String())
	if err != nil {
		return nil, err
	}

	var data Campaign
	err = json.Unmarshal(resp, &data)
	return &data, err
}

func (client *Client) GetCampaign(campaignID int, includes ...string) (*Pledge, error) {
	requestURI := defaultURI
	requestURI.Path = path.Join("api", "campaigns", strconv.Itoa(campaignID))

	if len(includes) > 0 {
		query := requestURI.Query()
		query.Add("includes", strings.Join(includes, ","))
		requestURI.RawQuery = query.Encode()
	}

	resp, err := client.get(requestURI.String())
	if err != nil {
		return nil, err
	}

	var data Pledge
	err = json.Unmarshal(resp, &data)
	return &data, err
}
