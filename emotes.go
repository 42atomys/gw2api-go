//go:generate easytags $GOFILE
package gw2api

type Emote struct {
	// The id of the emote.
	ID string `json:"id"`
	// List of ids of available commands for the emote.
	Commands []string `json:"commands"`
	// List of ids of the items resolvable against v2/items.
	UnlockItems []int `json:"unlock_items"`
}

// This resource returns a list of the emotes
// Return an array of ids for each type of currency.
func (r *Requestor) EmoteIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/emotes", &pointer)
	return r
}

// This resource returns a list of the emotes
// Return a list of response objects
func (r *Requestor) Emotes(pointer *[]*Emote, ids ...string) *Requestor {
	r.collection("/emotes", &pointer, ids)
	return r
}

// This resource returns a list of the emotes
// Return an object
func (r *Requestor) Emote(pointer *Emote, id string) *Requestor {
	r.singleton("/emotes", &pointer, id)
	return r
}
