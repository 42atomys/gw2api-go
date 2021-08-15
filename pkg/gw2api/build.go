//go:generate easytags $GOFILE
package gw2api

type Build struct {
	// ID containing the current build id as a number.
	ID int `json:"id"`
}

// This resource returns the current build id of the game.
// This can be used, for example, to register when event timers reset
// due to server restarts.
func (r *Requestor) Build(build *Build, id int) *Requestor {
	r.request("/builds", nil, &build)
	return r
}
