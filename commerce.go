//go:generate easytags $GOFILE
package gw2api

import (
	"fmt"
	"net/url"
	"time"
)

type CommerceDelivery struct {
	// The amount of coins ready for pickup.
	Coins int `json:"coins"`
	// Array of items
	Items []CommerceItem `json:"items"`
}

type CommerceItem struct {
	// The item's id, resolvable against v2/items.
	ID int `json:"id"`
	// The amount of this item ready for pickup.
	Count int `json:"count"`
}

type CommerceExchange struct {
	// The number of coins you required for a single gem.
	CoinsPerGem int `json:"coins_per_gem"`
	// The number of gems you get for the specified quantity of coins.
	Quantity int `json:"quantity"`
}

type CommerceListings struct {
	// The item id.
	ID int `json:"id"`
	// A list of all buy listings, ascending from lowest buy order.
	Buys []CommerceListingsItem `json:"buys"`
	// A list of all sell listings, ascending from lowest sell offer.
	Sells []CommerceListingsItem `json:"sells"`
}

type CommerceListingsItem struct {
	// The number of individual listings this object refers to
	// (e.g two players selling at the same price will end up in the same listing)
	Listings int `json:"listings"`
	// The sell offer or buy order price in coins.
	UnitPrice int `json:"unit_price"`
	// The amount of items being sold/bought in this listing.
	Quantity int `json:"quantity"`
}

type CommercePrices struct {
	// The item id.
	ID int `json:"id"`
	// A list of all buy listings, ascending from lowest buy order.
	Buys []CommercePricesItem `json:"buys"`
	// A list of all sell listings, ascending from lowest sell offer.
	Sells []CommercePricesItem `json:"sells"`
	// Indicates whether or not a free to play account can purchase or sell this item on the trading post.
	Whitelisted bool `json:"whitelisted"`
}

type CommercePricesItem struct {
	// The highest buy order or lowest sell offer price in coins.
	UnitPrice int `json:"unit_price"`
	// The amount of items being sold/bought.
	Quantity int `json:"quantity"`
}

type CommerceTransaction struct {
	// Id of the transaction.
	ID int `json:"id"`
	// The item id.
	ItemID int `json:"item_id"`
	// The price in coins.
	Price int `json:"price"`
	// The quantity of the item.
	Quantity int `json:"quantity"`
	// The date of creation, using ISO-8601 standard.
	Created time.Time `json:"created"`
	// The date of purchase, using ISO-8601 standard. Not shown in current
	// second-level endpoint.
	Purchased time.Time `json:"purchased"`
}

// This resource returns a list of accepted resources for the gem exchange.
func (r *Requestor) CommerceDelivery(pointer *CommerceDelivery) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionTradingpost).
		request("/commerce/delivery", nil, &pointer)
	return r
}

// This resource returns the current coins to gems exchange rate.
// The amount of coins to exchange for gems.
func (r *Requestor) CommerceExchangeCoins(pointer *CommerceExchange, amount int) *Requestor {
	r.request("/commerce/exchange/coins", url.Values{"quantity": []string{fmt.Sprint(amount)}}, &pointer)
	return r
}

// This resource returns the current gem to coins exchange rate.
// The amount of coins to exchange for gems.
func (r *Requestor) CommerceExchangeGems(pointer *CommerceExchange, amount int) *Requestor {
	r.request("/commerce/exchange/gems", url.Values{"quantity": []string{fmt.Sprint(amount)}}, &pointer)
	return r
}

// This resource returns current buy and sell listings from the trading post.
// Return a list of IDs
func (r *Requestor) CommerceListingsIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/commerce/listings", &pointer)
	return r
}

// This resource returns current buy and sell listings from the trading post.
// Return an array of given Listings
func (r *Requestor) CommerceListings(pointer *[]*CommerceListings, ids ...int) *Requestor {
	r.collection("/commerce/listings", &pointer, ids)
	return r
}

// This resource returns current buy and sell listings from the trading post.
// Return a specific listing
func (r *Requestor) CommerceListing(pointer *CommerceListings, id int) *Requestor {
	r.singleton("/commerce/listings", &pointer, id)
	return r
}

// This resource returns current aggregated buy and sell listing
// information from the trading post.
// Return a list of IDs
func (r *Requestor) CommercePricesIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/commerce/prices", &pointer)
	return r
}

// This resource returns current aggregated buy and sell listing
// information from the trading post.
// Return an array of given Prices
func (r *Requestor) CommercePrices(pointer *[]*CommercePrices, ids ...int) *Requestor {
	r.collection("/commerce/prices", &pointer, ids)

	return r
}

// This resource returns current aggregated buy and sell listing
// information from the trading post.
// Return a specific listing
func (r *Requestor) CommercePrice(pointer *CommercePrices, id int) *Requestor {
	r.singleton("/commerce/prices", &pointer, id)
	return r
}

// This resource provides access to the current and historical transactions
// of a player. This is an authenticated endpoint.
// Results are cached for five minutes.
// @see https://forum-en.gw2archive.eu/forum/community/api/Launching-v2-commerce-transactions/first
// Currently unfulfilled transactions.
func (r *Requestor) CommerceTransactionsCurrentBuys(pointer *[]*CommerceTransaction) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionTradingpost).
		request("/commerce/transactions/current/buys", nil, &pointer)
	return r
}

// This resource provides access to the current and historical transactions
// of a player. This is an authenticated endpoint.
// Results are cached for five minutes.
// @see https://forum-en.gw2archive.eu/forum/community/api/Launching-v2-commerce-transactions/first
// Currently unfulfilled transactions.
func (r *Requestor) CommerceTransactionsCurrentSells(pointer *[]*CommerceTransaction) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionTradingpost).
		request("/commerce/transactions/current/sells", nil, &pointer)
	return r
}

// This resource provides access to the current and historical transactions
// of a player. This is an authenticated endpoint.
// Results are cached for five minutes.
// @see https://forum-en.gw2archive.eu/forum/community/api/Launching-v2-commerce-transactions/first
// Fulfilled transactions of the past 90 days.
func (r *Requestor) CommerceTransactionsHistoryBuys(pointer *[]*CommerceTransaction) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionTradingpost).
		request("/commerce/transactions/history/buys", nil, &pointer)
	return r
}

// This resource provides access to the current and historical transactions
// of a player. This is an authenticated endpoint.
// Results are cached for five minutes.
// @see https://forum-en.gw2archive.eu/forum/community/api/Launching-v2-commerce-transactions/first
// Fulfilled transactions of the past 90 days.
func (r *Requestor) CommerceTransactionsHistorySells(pointer *[]*CommerceTransaction) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionTradingpost).
		request("/commerce/transactions/history/sells", nil, &pointer)
	return r
}
