//go:generate easytags $GOFILE
package gw2api

type World struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Population string `json:"population"`
}

func (r *Requestor) World(pointer *World, id int) *Requestor {
	r.singleton("/worlds", &pointer, id)
	return r
}

func (r *Requestor) WorldIDs(pointer []int) *Requestor {
	r.collectionIDs("/worlds", &pointer)
	return r
}

func (r *Requestor) Worlds(worlds *[]*World, ids ...int) *Requestor {
	r.collection("/worlds", &worlds, ids)
	return r
}
