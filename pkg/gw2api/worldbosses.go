//go:generate easytags $GOFILE
package gw2api

type WorldBoss struct {
	ID string `json:"id"`
}

func (r *Requestor) WorldBoss(pointer *WorldBoss, id string) *Requestor {
	r.singleton("/worldbosses", &pointer, id)
	return r
}

func (r *Requestor) WorldBosses(pointer *[]string) *Requestor {
	r.collectionIDs("/worldbosses", &pointer)
	return r
}
