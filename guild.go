//go:generate easytags $GOFILE
package gw2api

import (
	"fmt"
	"net/url"
	"time"
)

type Guild struct {
	// The current level of the guild.
	Level int `json:"level"`
	// The message of the day written out in a single string.
	Motd string `json:"motd"`
	// The guild's current influence.
	Influence int `json:"influence"`
	// The guild's current aetherium level.
	Aetherium string `json:"aetherium"`
	// The guild's current level of favor.
	Favor int `json:"favor"`
	// The number of People currently in the Guild.
	MemberCount int `json:"member_count"`
	// The maximum number of People that can be in the Guild.
	// @see https://wiki.guildwars2.com/wiki/Guild#Membership
	MemberCapacity int `json:"member_capacity"`
	// The unique guild id.
	ID string `json:"id"`
	// The guild's name.
	Name string `json:"name"`
	// The 2 to 4 letter guild tag representing the guild.
	Tag string `json:"tag"`
	// The guild emblem
	Emblem GuildEmblem `json:"emblem"`
}

type GuildEmblem struct {
	// An object containing information of the background of the guild emblem.
	Background GuildEmblemPart `json:"background"`
	// An object containing information of the foreground of the guild emblem
	Foreground GuildEmblemPart `json:"foreground"`
	// An array containing the manipulations done to the logo.
	// Possible values:
	//   FlipBackgroundHorizontal
	//   FlipBackgroundVertical
	Flags []string `json:"flags"`
}

type GuildEmblemPart struct {
	//  The background id, to be resolved against
	// https://api.guildwars2.com/v2/emblem/{foregrounds,backgrounds}
	ID int `json:"id"`
	// An array of numbers containing the id of each color used.
	Colors []int `json:"colors"`
}

type GuildLog struct {
	// An ID to uniquely identify the log entry within the scope of the guild.
	// Not globally unique.
	ID int `json:"id"`
	//  ISO-8601 standard timestamp for when the log entry was created.
	Time time.Time `json:"time"`
	// The account name of the guild member who generated this log entry
	// Known not to return a user key when the type is equal to upgrade and
	// action is queued.
	//   `type == joined`       - user has joined the guild
	//   `type == invited`      - user has been invited to the guild
	//   `type == kick`         - user  has been kicked from the guild
	//   `type == rank_change`  - user  has been changed
	User string `json:"user"`
	// The type of log entry. Additional fields are given depending on the type.
	// Posible values:
	//   `joined`       - A player join the guild
	//   `invited`      - A player has invited on the guild
	//   `kick`         - A player has beed kicked of the guild
	//   `rank_change`  - Rank of a player has beed changed
	//   `treasury`     - A guild member has deposited an item into the guilds
	//                    treasury.
	//   `stash`        - A guild member has deposited/withdrawn an item into
	//                    the guild stash
	//   `motd`         - A guild member has changed the guilds MOTD
	//   `upgrade`      - A guild member has interacted with a guild upgrade.
	Type string `json:"type"`

	// Account name of the guild member which invited the player.
	// Only when type is equals to `invited`
	InvitedBy string `json:"invited_by"`
	// Account name of the guild member which kicked the player.
	// Only when type is equals to `kick`
	KickedBy string `json:"kicked_by"`
	// Account name of the guild member which changed the player rank.
	// Only when type is equals to `rank_change`
	ChangedBy string `json:"changed_by"`
	// Old rank name.
	// Only when type is equals to `rank_change`
	OldRank string `json:"old_rank"`
	// New rank name.
	// Only when type is equals to `rank_change`
	NewRank string `json:"new_rank"`
	// The item ID that was deposited
	// Only when type is equals to `treasury` and `stash`
	ItemID int `json:"item_id"`
	// How many of the specified item was deposited.
	// Only when type is equals to `treasury` and `stash`
	Count int `json:"count"`

	// Posible values:
	//   deposit
	//   withdraw
	//   move
	// Only when type is equals to `stash`
	Operation string `json:"operation"`
	// How many coins (in copper) were deposited.
	// Only when type is equals to `stash`
	Coins int `json:"coins"`
	// The new MOTD.
	// Only when type is equals to `motd`
	Motd string `json:"motd"`

	// The type of interaction had. Possible values:
	//   queued
	//   cancelled
	//   completed - Having this action will also generate a new count
	//               field indicating how many upgrades were added.
	//   sped_up
	// Only when type is equals to `upgrade`
	Action string `json:"action"`
	// The upgrade ID which was completed.
	// Only when type is equals to `upgrade`
	UpgradeID int `json:"upgrade_id"`
	// May be added if the upgrade was created through a scribe station by
	// a scribe.
	// Only when type is equals to `upgrade`
	RecipeID int `json:"recipe_id"`
}

type GuildMember struct {
	// The account name of the member.
	Name string `json:"name"`
	// The guild rank of the member. (See /v2/guild/:id/ranks.)
	Rank string `json:"rank"`
	// The time and date the member joined the guild (ISO-8601 standard).
	// May also be null -- the join date was not tracked before around
	// March 19th, 2013.
	Joined time.Time `json:"joined"`
}

type GuildRank struct {
	// The ID and name of the rank.
	ID string `json:"id"`
	// A number given to each rank to specify how they should be sorted,
	// lowest being the first and highest being the last.
	Order int `json:"order"`
	// An array of permission ids from /v2/guild/permissions that have been
	// given to this rank.
	Permissions []string `json:"permissions"`
	// A URL pointing to the image currently assigned to this rank.
	Icon string `json:"icon"`
}

type GuildStash struct {
	// ID of the guild upgrade that granted access to this section of the
	// guild vault.
	UpgradeID int `json:"upgrade_id"`
	// The number of slots in this section of the guild vault.
	Size int `json:"size"`
	// The number of coins deposited in this section of the guild vault.
	Coins int `json:"coins"`
	// The description set for this section of the guild's vault.
	Note string `json:"note"`
	// An array of objects describing the items in the guild stash of exactly
	// size length. Each object either contains the following properties if an
	// item is present, or is null if the slot is empty
	Inventory []GuildInventoryItem `json:"inventory"`
}

type GuildInventoryItem struct {
	// ID of the item in this slot.
	ID int `json:"id"`
	// Amount of items in this slot.
	Coint int `json:"coint"`
}

type GuildTreasury struct {
	// The item's ID.
	ItemID int `json:"item_id"`
	// How many of the item are currently in the treasury.
	Count int `json:"count"`
	// An array of objects describing which currently in-progress upgrades
	// are needing this item.
	NeededBy []GuildTreasuryNeededBy `json:"needed_by"`
}

type GuildTreasuryNeededBy struct {
	// The ID of the upgrade needing this item.
	UpgradeID int `json:"upgrade_id"`
	// The total amount the upgrade needs for this item.
	Count int `json:"count"`
}

type GuildPermission struct {
	// The permission id.
	ID string `json:"id"`
	// The name of the permission.
	Name string `json:"name"`
	// The description of the permission.
	Description string `json:"description"`
}

type GuildUpgrade struct {
	// The upgrade id.
	ID string `json:"id"`
	// The name of the upgrade.
	Name string `json:"name"`
	// The guild upgrade description.
	Description string `json:"description"`
	// The upgrade type. Some upgrade types will add additional fields
	// to describe them further.
	// Possible values:
	//   `AccumulatingCurrency`  – Used for mine capacity upgrades.
	//   `BankBag`               – Used for guild bank upgrades.
	//   `Boost`                 – Used for permanent guild buffs such as waypoint
	//                             cost reduction.
	//   `Claimable`             – Used for guild WvW tactics.
	//   `Consumable`            – Used for banners and guild siege.
	//   `Decoration`            – Used for decorations that must be crafted by
	//                             a Scribe.
	//   `GuildHall`             - Used for claiming a Guild Hall.
	//   `GuildHallExpedition`   - Used for the Expedition unlock.
	//   `Hub`                   – Used for the Guild Initiative office unlock.
	//   `Queue`                 - Used for Workshop Restoration 1.
	//   `Unlock`                – Used for permanent unlocks, such as merchants
	//                             and arena decorations.
	Type string `json:"type"`
	// The maximum item slots of the guild bank tab.
	// Only when type is equals to `BankBag`
	BagMaxItems int `json:"bag_max_items"`
	// The maximum amount of coins that can be stored in the bank tab.
	// Only when type is equals to `BankBag`
	BagMaxCoins int `json:"bag_max_coins"`
	// A URL pointing to an icon for the upgrade.
	Icon string `json:"icon"`
	// The time it takes to build the upgrade.
	BuildTime int `json:"build_time"`
	// The prerequisite level the guild must be at to build the upgrade.
	RequiredLevel int `json:"required_level"`
	// The amount of guild experience that will be awarded upon building the
	// upgrade.
	Experience int `json:"experience"`
	// An array of upgrade IDs that must be completed before this can be built.
	Prerequisites []int `json:"prerequisites"`
	// An array of objects describing the upgrade's cost.
	Costs []GuildUpgradeCost `json:"costs"`
}

type GuildUpgradeCost struct {
	// The type of cost. One of:
	//   Item, Collectible, Currency, Coins
	Type string `json:"type"`
	// The name of the cost.
	Name string `json:"name"`
	// The amount needed.
	Count int `json:"count"`
	// The id of the item, if applicable, resolvable against v2/items.
	ItemID int `json:"item_id"`
}

// This resource returns core details about a given guild. The end point will
// include more or less fields dependend on whether or not an API Key of a
// Leader or Member of the Guild with the guilds scope is included
// in the Request.
func (r *Requestor) Guild(pointer *Guild, id string) *Requestor {
	r.request(fmt.Sprintf("/guild/%s", id), nil, &pointer)
	return r
}

// This resource returns information about certain events in a guild's log.
// The endpoint requires the scope guilds, and will only work if the API key is
// from the guild leader's account.
// The endpoint will return roughly the last 100 events of each type at max.
//
//   @params `since` - The optional parameter since can be added to a guild log
//                     query. The since parameter takes an event id value as a
//                     value. This will filter out all log entries not newer than
//                     the id given. Any events when an id is equal to, or less
//                     than since will be omitted.
//                     (set since to empty string to avoid)
func (r *Requestor) GuildLogs(pointer *[]*GuildLog, guildID string, sinceID string) *Requestor {
	urlValues := url.Values{}
	if sinceID != "" {
		urlValues["since"] = []string{sinceID}
	}

	r.
		needPerms(TokenPermissionAccount, TokenPermissionGuilds).
		request(fmt.Sprintf("/guild/%s/log", guildID), urlValues, &pointer)

	return r
}

// This resource returns information about the members of a specified guild.
// The endpoint requires the scope guilds, and will only work if the API key
// is from the guild leader's account.
func (r *Requestor) GuildMembers(pointer *[]*GuildMember, guildID string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionGuilds).
		request(fmt.Sprintf("/guild/%s/members", guildID), nil, &pointer)
	return r
}

// This resource returns information about the ranks of a specified guild.
// The endpoint requires the scope guilds, and will only work if the API key
// is from the guild leader's account.
func (r *Requestor) GuildRanks(pointer *[]*GuildRank, guildID string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionGuilds).
		request(fmt.Sprintf("/guild/%s/ranks", guildID), nil, &pointer)
	return r
}

// This resource returns information about the items in a guild's vault.
// The endpoint requires the scope guilds, and will only work if the API key
// is from the guild leader's account.
func (r *Requestor) GuildStash(pointer *[]*GuildStash, guildID string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionGuilds).
		request(fmt.Sprintf("/guild/%s/stash", guildID), nil, &pointer)
	return r
}

// This resource returns information about the items in a guild's storage.
// The endpoint requires the scope guilds, and will only work if the API key
//is from the guild leader's account.
func (r *Requestor) GuildStorage(pointer *[]*GuildInventoryItem, guildID string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionGuilds).
		request(fmt.Sprintf("/guild/%s/storage", guildID), nil, &pointer)
	return r
}

// This resource returns information about the items in a guild's treasury.
// The endpoint requires the scope guilds, and will only work if the API key
// is from the guild leader's account.
func (r *Requestor) GuildTreasury(pointer *[]*GuildTreasury, guildID string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionGuilds).
		request(fmt.Sprintf("/guild/%s/treasury", guildID), nil, &pointer)
	return r
}

// This resource returns information about the guild's upgrades. The endpoint
// requires the scope guilds, and will only work if the API key is
// from the guild leader's account.
func (r *Requestor) GuildUpgrades(pointer *[]int, guildID string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionGuilds).
		request(fmt.Sprintf("/guild/%s/upgrades", guildID), nil, &pointer)
	return r
}

// This resource returns a list of the guild permissions
// Return an array of ids for each type of currency.
func (r *Requestor) GuildPermissionIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/guild/permissions", &pointer)
	return r
}

// This resource returns a list of the guild permissions
// Return a list of response objects
func (r *Requestor) GuildPermissions(pointer *[]*GuildPermission, ids ...string) *Requestor {
	r.collection("/guild/permissions", &pointer, ids)
	return r
}

// This resource returns a list of the guild permissions
// Return an object
func (r *Requestor) GuildPermission(pointer *GuildPermission, id string) *Requestor {
	r.singleton("/guild/permissions", &pointer, id)
	return r
}

// This resource returns information on guild ids to be used for other
// API queries.
//   @param name - The guild name must be given in order to obtain the
//                 relevant id.
func (r *Requestor) GuildSearch(pointer *[]string, name string) *Requestor {
	r.singleton("/guild/permissions", &pointer, name)
	return r
}

// This resource returns a list of the guild upgrades
// Return an array of ids for each type of currency.
func (r *Requestor) UnscopedGuildUpgradeIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/guild/upgrades", &pointer)
	return r
}

// This resource returns a list of the guild upgrades
// Return a list of response objects
func (r *Requestor) UnscopedGuildUpgrades(pointer *[]*GuildUpgrade, ids ...string) *Requestor {
	r.collection("/guild/upgrades", &pointer, ids)
	return r
}

// This resource returns a list of the guild upgrades
// Return an object
func (r *Requestor) UnscopedGuildUpgrade(pointer *GuildUpgrade, id string) *Requestor {
	r.singleton("/guild/upgrades", &pointer, id)
	return r
}
