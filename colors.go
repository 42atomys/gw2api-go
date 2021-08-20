//go:generate easytags $GOFILE
package gw2api

import "net/url"

type Color struct {
	// The color id.
	ID int `json:"id"`
	// The color name
	Name string `json:"name"`
	// The base RGB values.
	BaseRGB []int `json:"base_rgb"`
	// Detailed information on its appearance when applied on cloth armor.
	Cloth ColorStruct `json:"cloth"`
	// Detailed information on its appearance when applied on leather armor.
	Leather ColorStruct `json:"leather"`
	// Detailed information on its appearance when applied on metal armor.
	Metal ColorStruct `json:"metal"`
	// Detailed information on its appearance when applied on fur armor.
	Fur ColorStruct `json:"fur"`
	// ID of the dye item.
	Item int `json:"item"`
	// Color categories. Possible values:
	//   Hue: `Gray`, `Brown`, `Red`, `Orange`, `Yellow`, `Green`, `Blue`, `Purple`
	//   Material: `Vibrant`, `Leather`, `Metal`
	//   Rarity: `Starter`, `Common`, `Uncommon`, `Rare`, `Exclusive`
	Categories []string `json:"categories"`
}

type ColorStruct struct {
	// The brightness.
	Brightness float32 `json:"brightness"`
	// The contrast.
	Contrast float32 `json:"contrast"`
	// The hue in the HSL colorspace.
	// @see https://en.wikipedia.org/wiki/HSL_and_HSV
	Hue float32 `json:"hue"`
	// The saturation in the HSL colorspace.
	// @see https://en.wikipedia.org/wiki/Colorfulness#Saturation
	Saturation float32 `json:"saturation"`
	// The lightness in the HSL colorspace.
	// @see https://en.wikipedia.org/wiki/Lightness
	Lightness float32 `json:"lightness"`
	// A list containing precalculated RGB values.
	RGB []int `json:"rgb"`
}

// This resource returns information about the Biography answers that are in the game.
func (r *Requestor) Color(color *Color, id string) *Requestor {
	r.request("/backstory/answers", url.Values{"id": []string{id}}, &color)
	return r
}

// This resource returns information about the Biography answers that are in the game.
func (r *Requestor) Colors(colors *[]*Color, ids ...string) *Requestor {
	r.request("/backstory/answers", url.Values{"id": ids}, &colors)
	return r
}

// This resource returns information about the Biography answers that are in the game.
func (r *Requestor) ColorsIDs(s []string) *Requestor {
	r.request("/backstory/answers", nil, &s)
	return r
}
