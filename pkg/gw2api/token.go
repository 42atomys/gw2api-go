package gw2api

type TokenInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

func (r *Requestor) TokenInfo() (token TokenInfo, err error) {
	r.request("/tokeninfo", nil, &token)
	return
}
