package patreon

import (
	"errors"
	"fmt"

	log "github.com/kirillDanshin/dlog"
	json "github.com/pquerna/ffjson/ffjson"
	http "github.com/valyala/fasthttp"
	"golang.org/x/oauth2"
)

const baseURL = "https://api.patreon.com"

func post(token *oauth2.Token, dst []byte, requestURI string) ([]byte, error) {
	return request(token, dst, "POST", requestURI)
}

func get(token *oauth2.Token, requestURI string) ([]byte, error) {
	return request(token, nil, "GET", requestURI)
}

func request(token *oauth2.Token, dst []byte, method, requestURI string) ([]byte, error) {
	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.Header.Set("Authorization", fmt.Sprintln(token.Type(), token.AccessToken))
	req.SetRequestURI(requestURI)
	req.Header.SetMethod(method)
	if dst != nil {
		req.SetBody(dst)
	}

	resp := http.AcquireResponse()
	defer http.ReleaseResponse(resp)
	err := http.Do(req, resp)
	log.D(req)
	log.D(resp)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		var data []Error
		err = json.Unmarshal(resp.Body(), &data)
		return nil, errors.New(fmt.Sprintln(data[0].CodeName, data[0].Detail))
	}

	return resp.Body(), err
}
