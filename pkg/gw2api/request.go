package gw2api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Requestor struct {
	userAgent     string
	schemaVersion string
	context       context.Context
	token         AuthToken
	lang          Lang
	timeout       time.Duration
	permissions   uint

	err error
}

type AuthToken string
type Lang string

const (
	LangFR Lang = "fr"
	LangEN Lang = "en"
	LangES Lang = "es"
	LangDE Lang = "de"
	LangZH Lang = "zh"
)

var (
	ErrTooManyRequest = errors.New("too many request: 429")
	BaseURL, _        = url.Parse("https://api.guildwars2.com/v2")
)

func NewRequestor() *Requestor {
	requestor := &Requestor{
		userAgent:     "gw2api-go:0.1",
		schemaVersion: "2021-08-14T00:00:00Z",
		timeout:       15 * time.Second,
		context:       context.TODO(),
	}

	return requestor
}

func (r *Requestor) Timeout(timeout time.Duration) *Requestor {
	r.timeout = timeout
	return r
}

func (r *Requestor) Auth(token string) *Requestor {
	r.token = AuthToken(token)
	r.permissions = 0
	tki, err := r.TokenInfo()
	if err != nil {
		r.err = err
	}

	for _, perm := range tki.Permissions {
		if p, e := permissionsMapping[perm]; e {
			setBitwise(&r.permissions, uint(p))
		} else {
			// TODO new permission detected : alert Atomys
			// log.Print("Found new permission: ", perm)
		}
	}

	return r
}

func (r *Requestor) Lang(lang Lang) *Requestor {
	r.lang = lang
	return r
}

func (r *Requestor) Err() error {
	return r.err
}

func (r *Requestor) request(endpoint string, queryParams url.Values, v interface{}) {
	url := BaseURL
	url.Path += endpoint

	if queryParams != nil {
		url.RawQuery = queryParams.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		r.err = err
		return
	}

	if r.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.token))
	}
	req.Header.Set("X-Schema-Version", r.schemaVersion)
	req.Header.Set("User-Agent", r.userAgent)
	req.Header.Set("Accept-Language", string(r.lang))
	req = req.WithContext(r.context)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		r.err = err
		return
	}

	switch response.StatusCode {
	case http.StatusOK, http.StatusNotModified:
		if err = json.NewDecoder(response.Body).Decode(&v); err != nil {
			r.err = err
		}
	case http.StatusTooManyRequests:
		r.err = ErrTooManyRequest
		return
	default:
		var apiErr *APIError
		if err = json.NewDecoder(response.Body).Decode(&apiErr); err != nil {
			r.err = err
			return
		}
		r.err = apiErr
	}
}
