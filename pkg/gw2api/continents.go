//go:generate easytags $GOFILE
package gw2api

type Continent struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	ContinentDims []int  `json:"continent_dims"`
	MinZoom       int    `json:"min_zoom"`
	MaxZoom       int    `json:"max_zoom"`
	Floors        []int  `json:"floors"`
}

func (r *Requestor) ContinentIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/continents", &pointer)
	return r
}

func (r *Requestor) Continents(pointer *[]*Continent, ids ...int) *Requestor {
	r.collection("/continents", &pointer, ids)
	return r
}

func (r *Requestor) Continent(pointer *Continent, id int) *Requestor {
	r.singleton("/continents", &pointer, id)
	return r
}
