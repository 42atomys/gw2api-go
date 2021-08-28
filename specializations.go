//go:generate easytags $GOFILE
package gw2api

type Specialization struct {
	// The specialization's ID.
	ID int `json:"id"`
	// The name of the specialization.
	Name string `json:"name"`
	// The profession that this specialization belongs to.
	Profession string `json:"profession"`
	// true if this specialization is an Elite specialization, false otherwise.
	Elite bool `json:"elite"`
	// A URL to an icon of the specialization.
	Icon string `json:"icon"`
	// A URL to the background image of the specialization.
	Background string `json:"background"`
	// Contains a list of IDs specifying the minor traits in the specialization.
	MinorTraits []int `json:"minor_traits"`
	// Contains a list of IDs specifying the major traits in the specialization.
	MajorTraits []int `json:"major_traits"`
}

// This resource returns information about the specializations that are in the game.
// Return an array of ids for each specializations.
func (r *Requestor) SpecializationIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/specializations", &pointer)
	return r
}

// This resource returns information about the specializations that are in the game.
// Return a list of response objects
func (r *Requestor) Specializations(pointer *[]*Specialization, ids ...int) *Requestor {
	r.collection("/specializations", &pointer, ids)
	return r
}

// This resource returns information about the specializations that are in the game.
// Return an object
func (r *Requestor) Specialization(pointer *Specialization, id int) *Requestor {
	r.singleton("/specializations", &pointer, id)
	return r
}
