package patreon

import (
	"fmt"

	log "github.com/kirillDanshin/dlog"
	http "github.com/valyala/fasthttp"
)

func (client *Client) post(dst []byte, requestURI string) ([]byte, error) {
	return client.request(dst, "POST", requestURI)
}

func (client *Client) get(requestURI string) ([]byte, error) {
	return client.request(nil, "GET", requestURI)
}

func (client *Client) request(dst []byte, method, requestURI string) ([]byte, error) {
	req := http.AcquireRequest()
	defer http.ReleaseRequest(req)
	req.SetRequestURI(requestURI)
	req.Header.SetMethod(method)
	if client.Token != nil {
		req.Header.Set(
			"Authorization",
			fmt.Sprint(client.Token.Type(), " ", client.Token.AccessToken),
		)
	}
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

	return resp.Body(), errMap[resp.StatusCode()]
}
