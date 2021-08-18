//go:generate easytags $GOFILE
package gw2api

type Currency struct {
	// The currency's ID.
	ID int `json:"id"`
	// The currency's name.
	Name string `json:"name"`
	// A description of the currency.
	Description string `json:"description"`
	// A URL to an icon for the currency.
	Icon string `json:"icon"`
	// A number that can be used to sort the list of currencies when ordered
	// from least to greatest.
	Order int `json:"order"`
}

// This resource returns a list of the currencies contained in the
// account wallet.
// Return an array of ids for each type of currency.
func (r *Requestor) CurrencyIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/currencies", &pointer)
	return r
}

// This resource returns a list of the currencies contained in the
// account wallet.
// Return a list of response objects
func (r *Requestor) Currencys(pointer *[]*Currency, ids ...int) *Requestor {
	r.collection("/currencies", &pointer, ids)
	return r
}

// This resource returns a list of the currencies contained in the
// account wallet.
// Return an object
func (r *Requestor) Currency(pointer *Currency, id int) *Requestor {
	r.singleton("/currencies", &pointer, id)
	return r
}
