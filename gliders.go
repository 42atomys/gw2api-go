//go:generate easytags $GOFILE
package gw2api

type Glider struct {
	// The id of the glider.
	ID int `json:"id"`
	// An array of item ids used to unlock the glider.
	// Can be resolved against v2/items
	UnlockItems []int `json:"unlock_items"`
	// The order in which the glider appears in a list.
	// This value is not unique.
	Order int `json:"order"`
	// The icon uri for the glider.
	Icon string `json:"icon"`
	// The name of the glider as it appears in-game.
	Name string `json:"name"`
	// The in-game glider description.
	Description string `json:"description"`
	// List of dye ids. Can be resolved against v2/colors.
	DefaultDyes []int `json:"default_dyes"`
}

// This resource returns a list of the gliders
// Return an array of ids for each type of currency.
func (r *Requestor) GliderIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/gliders", &pointer)
	return r
}

// This resource returns a list of the gliders
// Return a list of response objects
func (r *Requestor) Gliders(pointer *[]*Glider, ids ...int) *Requestor {
	r.collection("/gliders", &pointer, ids)
	return r
}

// This resource returns a list of the gliders
// Return an object
func (r *Requestor) Glider(pointer *Glider, id int) *Requestor {
	r.singleton("/gliders", &pointer, id)
	return r
}
