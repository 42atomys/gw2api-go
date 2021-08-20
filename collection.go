package gw2api

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// This resource returns a list of the dungeons
// Return an array of ids for each type of currency.
func (r *Requestor) collectionIDs(endpoint string, pointer interface{}) *Requestor {
	r.request(endpoint, nil, &pointer)
	return r
}

// This resource returns a list of the dungeons
// Return a list of response objects
func (r *Requestor) collection(endpoint string, pointer interface{}, ids ...interface{}) *Requestor {
	if len(ids) == 0 {
		r.err = errors.New("at least one id must be given")
		return r
	}

	var urlValues url.Values
	sIds := strings.Trim(strings.Replace(fmt.Sprint(ids...), " ", ",", -1), "[]")
	urlValues = url.Values{"ids": strings.Split(sIds, ",")}

	r.request(endpoint, urlValues, &pointer)
	return r
}

// This resource returns a list of the dungeons
// Return an object
func (r *Requestor) singleton(endpoint string, pointer interface{}, id interface{}) *Requestor {
	r.request(endpoint, url.Values{"id": []string{fmt.Sprint(id)}}, &pointer)
	return r
}
