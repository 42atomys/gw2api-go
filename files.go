//go:generate easytags $GOFILE
package gw2api

type File struct {
	// The file identifier.
	ID string `json:"id"`
	// The URL to the image.
	Icon string `json:"icon"`
}

// This resource returns a list of the files
// Return an array of ids for each type of currency.
func (r *Requestor) FileIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/files", &pointer)
	return r
}

// This resource returns a list of the files
// Return a list of response objects
func (r *Requestor) Files(pointer *[]*File, ids ...string) *Requestor {
	r.collection("/files", &pointer, ids)
	return r
}

// This resource returns a list of the files
// Return an object
func (r *Requestor) File(pointer *File, id string) *Requestor {
	r.singleton("/files", &pointer, id)
	return r
}
