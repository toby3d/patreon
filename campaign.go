package patreon

import (
	"net/url"
	"strings"

	log "github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	"golang.org/x/oauth2"
)

func GetCreatorProfile(token *oauth2.Token, includes ...string) (*ResponseList, error) {
	requestURI := &url.URL{
		Scheme: "https",
		Host:   "www.patreon.com",
		Path:   "/api/oauth2/api/current_user/campaigns",
	}
	if len(includes) > 0 {
		query := requestURI.Query()
		query.Add("includes", strings.Join(includes, ","))
		requestURI.RawQuery = query.Encode()
	}

	resp, err := get(token, requestURI.String())
	if err != nil {
		return nil, err
	}

	var data ResponseList
	err = json.Unmarshal(resp, &data)
	log.D(data)
	return &data, err
}
