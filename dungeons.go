//go:generate easytags $GOFILE
package gw2api

type Dungeon struct {
	// The name of the dungeon
	ID string `json:"id"`
	// An array of all paths on current dungeon
	Paths []DungeonPath `json:"paths"`
}

type DungeonPath struct {
	// The given name for the dungeon path.
	ID string `json:"id"`
	// The type of the chosen path. Can be either Story or Explorable
	Type string `json:"type"`
}

// This resource returns a list of the dungeons
// Return an array of ids for each type of currency.
func (r *Requestor) DungeonIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/dungeons", &pointer)
	return r
}

// This resource returns a list of the dungeons
// Return a list of response objects
func (r *Requestor) Dungeons(pointer *[]*Dungeon, ids ...string) *Requestor {
	r.collection("/dungeons", &pointer, ids)
	return r
}

// This resource returns a list of the dungeons
// Return an object
func (r *Requestor) Dungeon(pointer *Dungeon, id string) *Requestor {
	r.singleton("/dungeons", &pointer, id)
	return r
}
