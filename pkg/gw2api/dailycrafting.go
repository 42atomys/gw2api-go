//go:generate easytags $GOFILE
package gw2api

// This resource returns information about time-gated recipes that can
// be crafted in-game.
func (r *Requestor) DailyCrafting(pointer *[]string) *Requestor {
	r.request("/dailycrafting", nil, &pointer)
	return r
}
