//go:generate easytags $GOFILE
package gw2api

type Emblem struct {
	// The ID of the emblem part.
	ID int `json:"id"`
	// An array of URLs to images that make up the various parts of the emblem.
	Layers []string `json:"layers"`
}

// This resource returns image resources that are needed to render the
// background of guild emblems.
// Return an array of ids for each type of currency.
func (r *Requestor) EmblemBackgroundIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/emblem/backgrounds", &pointer)
	return r
}

// This resource returns image resources that are needed to render the
// background of guild emblems.
// Return a list of response objects
// This endpoint is limited to 200 ids.
//
// Requesting more than 200 ids will return:
//   {"text":"id list too long; this endpoint is limited to 200 ids at once"}
func (r *Requestor) EmblemBackgrounds(pointer *[]*Emblem, ids ...int) *Requestor {
	r.collection("/emblem/backgrounds", &pointer, ids)
	return r
}

// This resource returns image resources that are needed to render the
// background of guild emblems.
// Return an object
func (r *Requestor) EmblemBackground(pointer *Emblem, id int) *Requestor {
	r.singleton("/emblem/backgrounds", &pointer, id)
	return r
}

// This resource returns image resources that are needed to render the
// foreground  of guild emblems.
// Return an array of ids for each type of currency.
func (r *Requestor) EmblemForegroundIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/emblem/foregrounds", &pointer)
	return r
}

// This resource returns image resources that are needed to render the
// foreground  of guild emblems.
// Return a list of response objects
// This endpoint is limited to 200 ids.
//
// Requesting more than 200 ids will return:
//   {"text":"id list too long; this endpoint is limited to 200 ids at once"}
func (r *Requestor) EmblemForegrounds(pointer *[]*Emblem, ids ...int) *Requestor {
	r.collection("/emblem/foregrounds", &pointer, ids)
	return r
}

// This resource returns image resources that are needed to render the
// foreground  of guild emblems.
// Return an object
func (r *Requestor) EmblemForeground(pointer *Emblem, id int) *Requestor {
	r.singleton("/emblem/foregrounds", &pointer, id)
	return r
}
