//go:generate easytags $GOFILE
package gw2api

type Finisher struct {
	// The id of the finisher.
	ID int `json:"id"`
	// A description explaining how to acquire the finisher.
	UnlockDetails string `json:"unlock_details"`
	// An array of item ids used to unlock the finisher.
	// Can be resolved against v2/items
	UnlockItems []int `json:"unlock_items"`
	// The order in which the finisher appears in a list.
	Order int `json:"order"`
	// The icon uri for the finisher.
	Icon string `json:"icon"`
	// The name of the finisher as it appears in-game.
	Name string `json:"name"`
}

// This resource returns a list of the finishers
// Return an array of ids for each type of currency.
func (r *Requestor) FinisherIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/finishers", &pointer)
	return r
}

// This resource returns a list of the finishers
// Return a list of response objects
func (r *Requestor) Finishers(pointer *[]*Finisher, ids ...int) *Requestor {
	r.collection("/finishers", &pointer, ids)
	return r
}

// This resource returns a list of the finishers
// Return an object
func (r *Requestor) Finisher(pointer *Finisher, id int) *Requestor {
	r.singleton("/finishers", &pointer, id)
	return r
}
