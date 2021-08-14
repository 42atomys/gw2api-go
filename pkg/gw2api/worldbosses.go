//go:generate easytags $GOFILE
package gw2api

import (
	"net/url"
)

type WorldBoss struct {
	ID string `json:"id"`
}

func (r *Requestor) WorldBoss(worldBoss *WorldBoss, id string) *Requestor {
	r.request("/worldbosses", url.Values{"id": []string{id}}, &worldBoss)
	return r
}

func (r *Requestor) WorldBosses(s []string) *Requestor {
	r.request("/worldbosses", nil, &s)
	return r
}
