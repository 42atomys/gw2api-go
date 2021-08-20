//go:generate easytags $GOFILE
package gw2api

import (
	"net/url"
	"strings"
	"time"
)

type TokenInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

type Subtoken struct {
	Subtoken string `json:"subtoken"`
}

func (r *Requestor) TokenInfo() (token TokenInfo, err error) {
	r.request("/tokeninfo", nil, &token)
	return
}

// This resource allows for the creation of Subtokens; essentially API keys
// with a more limited set of permissions, which can be used as a
// substitute for them.
//   expire      - An ISO-8601 datetime specifing when the generated
//                 Subtoken will expire.
//   permissions - A comma separated list of permissions to inherit.
//                 Unrecognized permissions as well as permissions that are
//                 specified but are not granted to the API Key used in the
//                 Request are silently ignored.
//                 v2/tokeninfo may be queried for a list of avaiable choices
//                 to use for a given API key.
//   urls -        A comma separated list of Endpoints that will be accessible
//                 using this Subtoken.
//                 If no Endpoints are specified all Endpoints, not
//                 otherwise limited by 'permissions', will be accessible.
//
// Return a JSON Web Token which can be used like an API key but only with the requested limitations.
func (r *Requestor) CreateSubToken(pointer *Subtoken, expireAt time.Time, perms []string, urls []string) *Requestor {
	urlValues := url.Values{
		"expire":      []string{expireAt.Format(time.RFC3339)},
		"permissions": []string{strings.Join(perms, ",")},
		"urls":        []string{strings.Join(urls, ",")},
	}

	r.
		needPerms(TokenPermissionAccount).
		request("/createsubtoken", urlValues, &pointer)
	return r
}
