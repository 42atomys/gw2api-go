package gw2api

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type World struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population string `json:"population"`
}

func (r *Requestor) World(world *World, id int) *Requestor {
	r.request("/worlds", url.Values{"id": []string{fmt.Sprint(id)}}, &world)
	return r
}

func (r *Requestor) WorldIDs(s []int) *Requestor {
	r.request("/worlds", nil, &s)
	return r
}

func (r *Requestor) Worlds(worlds *[]*World, ids ...int) *Requestor {
	if len(ids) == 0 {
		r.err = errors.New("at least one id must be given")
	}

	sIds := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	r.request("/worlds", url.Values{"ids": strings.Split(sIds, ",")}, &worlds)
	return r
}
