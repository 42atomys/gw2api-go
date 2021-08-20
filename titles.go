//go:generate easytags $GOFILE
package gw2api

type Title struct {
	// The id of the title.
	ID int `json:"id"`
	// The name of the title (this is also the title displayed over
	// a character in-game.)
	Name string `json:"name"`
	// The id of the achievement that grants this title;
	// resolvable against v2/achievements.
	// (Now depreciated)
	Achievement int `json:"achievement"`
	// The id of the achievement that grants this title; resolvable
	// against v2/achievements.
	Achievements []int `json:"achievements"`
	// The amount of AP required to have said title.
	ApRequired int `json:"ap_required"`
}

// This resource returns information about the titles that are in the game.
// Return an array of ids for each type of currency.
func (r *Requestor) TitleIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/titles", &pointer)
	return r
}

// This resource returns information about the titles that are in the game.
// Return a list of response objects
func (r *Requestor) Titles(pointer *[]*Title, ids ...int) *Requestor {
	r.collection("/titles", &pointer, ids)
	return r
}

// This resource returns information about the titles that are in the game.
// Return an object
func (r *Requestor) Title(pointer *Title, id int) *Requestor {
	r.singleton("/titles", &pointer, id)
	return r
}
