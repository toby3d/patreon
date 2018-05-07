package patreon

import (
	"path"
	"strconv"
	"strings"

	json "github.com/pquerna/ffjson/ffjson"
)

func (client *Client) GetPledges(campaignID int, includes ...string) (*Pledge, error) {
	requestURI := defaultURI
	requestURI.Path = path.Join("api", "oauth2", "api", "campaigns", strconv.Itoa(campaignID), "pledges")

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
