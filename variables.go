package patreon

import (
	"errors"
	"net/url"

	http "github.com/valyala/fasthttp"
)

var (
	defaultURI = url.URL{
		Scheme: "https",
		Host:   "www.patreon.com",
	}

	defaultAPIURI = url.URL{
		Scheme: "https",
		Host:   "api.patreon.com",
	}
)

var (
	ErrBadRequest          = errors.New("something was wrong with your request (syntax, size too large, etc.)")
	ErrUnauthorized        = errors.New("authentication failed (bad API key, invalid OAuth token, incorrect scopes, etc.)")
	ErrForbidden           = errors.New("the requested is hidden for administrators only")
	ErrNotFound            = errors.New("the specified resource could not be found")
	ErrMethodNotAllowed    = errors.New("you tried to access a resource with an invalid method")
	ErrNotAcceptable       = errors.New("you requested a format that isn't json")
	ErrGone                = errors.New("the resource requested has been removed from our servers")
	ErrTooManyRequests     = errors.New("slow down!")
	ErrInternalServerError = errors.New("our server ran into a problem while processing this request, please try again later")
	ErrServiceUnavailable  = errors.New("we're temporarily offline for maintenance, please try again later")

	errMap = map[int]error{
		http.StatusBadRequest:          ErrBadRequest,
		http.StatusUnauthorized:        ErrUnauthorized,
		http.StatusForbidden:           ErrForbidden,
		http.StatusNotFound:            ErrNotFound,
		http.StatusMethodNotAllowed:    ErrMethodNotAllowed,
		http.StatusNotAcceptable:       ErrNotAcceptable,
		http.StatusGone:                ErrGone,
		http.StatusTooManyRequests:     ErrTooManyRequests,
		http.StatusInternalServerError: ErrInternalServerError,
		http.StatusServiceUnavailable:  ErrServiceUnavailable,
	}
)
