//go:generate easytags $GOFILE
package gw2api

import "time"

type Account struct {
	// The unique persistent account GUID.
	ID string `json:"id"`
	// The unique account name with numerical suffix.
	// It is possible that the name change.
	// Do not rely on the name, use id instead.
	// @see https://forum-en.gw2archive.eu/forum/community/api/Launching-v2-account-w-Authentication/page/1#post4830070
	Name string `json:"name"`
	// The age of the account in seconds.
	Age int `json:"age"`
	// The id of the home world the account is assigned to.
	// Can be resolved against /v2/worlds.
	World int `json:"world"`
	// A list of guilds assigned to the given account.
	Guilds []string `json:"guilds"`
	// A list of guilds the account is leader of. Requires the additional
	// guilds scope.
	GuildLeader []string `json:"guild_leader"`
	// An ISO-8601 standard timestamp of when the account was created.
	Created time.Time `json:"created"`
	// A list of what content this account has access to.
	// Possible values:
	// None – should probably never happen
	// PlayForFree – has not yet purchased the game
	// GuildWars2 – has purchased the base game
	// HeartOfThorns – has purchased Heart of Thorns
	// PathOfFire – has purchased Path of Fire
	Access []string `json:"access"`
	// True if the player has bought a commander tag.
	Commander bool `json:"commander"`
	// The account's personal fractal reward level.
	//* Requires the additional `progression` scope.
	FractalLevel int `json:"fractal_level"`
	// The daily AP the account has
	//* Requires the additional `progression` scope.
	DailyAp int `json:"daily_ap"`
	// The monthly AP the account has
	//* Requires the additional `progression` scope.
	MonthlyAp int `json:"monthly_ap"`
	// The account's personal wvw rank
	//* Requires the additional `progression` scope.
	WvwRank int `json:"wvw_rank"`
	// An ISO-8601 standard timestamp of when the account information last
	// changed as perceived by the API. This field is only present when a
	// Schema version of 2019-02-21T00:00:00Z or later is requested.
	// @see https://gitter.im/arenanet/api-cdi?at=5c7091129e430b3086af0943
	LastModified time.Time `json:"last_modified"`
}

type AccountAchievement struct {
	// The achievement id.
	ID int `json:"id"`
	// This attribute contains an array of numbers,
	// giving more specific information on the progress for the achievement.
	// The meaning of each value varies with each achievement.
	// Bits start at zero. If an achievement is done, the in-progress bits
	// are not displayed.
	Bits []int `json:"bits"`
	// The player's current progress towards the achievement.
	Current int `json:"current"`
	// The amount needed to complete the achievement.
	Max int `json:"max"`
	// Whether or not the achievement is done.
	Done bool `json:"done"`
	// The number of times the achievement has been completed
	// if the achievement is repeatable.
	Repeated int `json:"repeated"`
	// Whether or not the achievement is unlocked.
	// Note that if this property does not exist, the achievement
	// is unlocked as well.
	Unlocked bool `json:"unlocked"`
}

type AccountInventoryItem struct {
	// The item's ID.
	ID int `json:"id"`
	// The amount of items in the item stack.
	Count int `json:"count"`
	// The amount of charges remaining on the item.
	Charges int `json:"charges"`
	// The skin applied to the item, if it is different from its original.
	// Can be resolved against /v2/skins.
	Skin int `json:"skin"`
	// The IDs of the dyes applied to the item.
	// Can be resolved against /v2/colors.
	Dyes []int `json:"dyes"`
	// The item IDs of the runes or sigills applied to the item.
	Upgrades []int `json:"upgrades"`
	// The slot occupied by the upgrade at the corresponding position in
	// upgrades.
	UpgradeSlotIndices []interface{} `json:"upgrade_slot_indices"`
	// An array of item IDs for each infusion applied to the item.
	Infusions []interface{} `json:"infusions"`
	// The current binding of the item. Either Account or Character if present.
	Biding string `json:"biding"`
	// If binding is Character, this field tells which character it is bound to.
	BoundTo string `json:"bound_to"`
	// The stats of the item
	Stats AccountBankStat `json:"stats"`
}

type AccountBankStat struct {
	// The ID of the item's stats. Can be resolved against /v2/itemstats.
	ID string `json:"id"`
	// The list of stats provided by this item. May include the following:
	// AgonyResistance - Agony Resistance
	// BoonDuration - Concentration
	// ConditionDamage – Condition Damage
	// ConditionDuration - Expertise
	// CritDamage – Ferocity
	// Healing – Healing Power
	// Power – Power
	// Precision – Precision
	// Toughness – Toughness
	// Vitality – Vitality
	Attributes map[string]int `json:"attributes"`
}

type AccountBuildStorage struct {
	// The name of the template
	Name string `json:"name"`
	// The profession of the template build
	Profession string `json:"profession"`
	//Specializations
	Specializations AccountBuildStorageSpecialization `json:"specializations"`
	Skills          AccountBuildStorageSkills         `json:"skills"`
	AquaticSkills   AccountBuildStorageSkills         `json:"aquatic_skills"`

	// Ranger specific structure for Account BuildStorage
	Pets AccountBuildStoragePets `json:"pets"`

	// Revenant specific structure for Account BuildStorage
	Legends        []string `json:"legends"`
	AquaticLegends []string `json:"aquatic_legends"`
}

type AccountBuildStorageSpecialization struct {
	ID     int   `json:"id"`
	Traits []int `json:"traits"`
}

type AccountBuildStorageSkills struct {
	Heal      int   `json:"heal"`
	Utilities []int `json:"utilities"`
	Elite     int   `json:"elite"`
}

// Ranger specific structure for Account BuildStorage
type AccountBuildStoragePets struct {
	Terrestrial []int `json:"terrestrial"`
	Aquatic     []int `json:"aquatic"`
}

type AccountFinisher struct {
	// The id of the finisher resolvable against /v2/finishers.
	ID int `json:"id"`
	// Indicates if the finisher is permanent or temporary.
	Permanent bool `json:"permanent"`
	// If permanent is false, this field will indicate the number
	// of uses the finisher has remaining.
	Quantity int `json:"quantity"`
}

type AccountLegendaryArmory struct {
	// The id of the armory items, resolvable against /v2/items and
	// /v2/legendaryarmory.
	ID int `json:"id"`
	// The count of that item available for use in a single equipment template.
	Count int `json:"count"`
}

type AccountLuck struct {
	// The string "luck".
	ID string `json:"id"`
	// The amount of luck consumed
	Value int `json:"value"`
}

type AccountMastery struct {
	// The id of the mastery resolvable against /v2/masteries.
	ID int `json:"id"`
	// Indicates the level at which the mastery is on the account.
	// Is a 0-indexed reference to the /v2/masteries.levels array indicating
	// the maximum level unlocked by the user. If omitted,
	// this mastery hasn't been started.
	Level int `json:"level"`
}

type AccountMasteryPoint struct {
	Totals []struct {
		// The mastery region. Current possible options:
		// Tyria, Maguuma, Desert and Tundra.
		Region string `json:"region"`
		// Amount of masteries of this region spent in mastery tracks.
		Spent int `json:"spent"`
		// Amount of masteries of this region earned for the account.
		Earned int `json:"earned"`
	} `json:"totals"`

	// Array of unlocked mastery ids.
	Unlocked []int `json:"unlocked"`
}

type AccountMaterial struct {
	// The item ID of the material.
	ID int `json:"id"`
	// The material category the item belongs to.
	// Can be resolved against /v2/materials.
	Category int `json:"category"`
	// The binding of the material. Either Account or omitted.
	Binding string `json:"binding"`
	// The number of the material that is stored in the account vault.
	Count int `json:"count"`
}

type AccountCurrency struct {
	// The currency's ID that can be resolved against /v2/currencies.
	ID int `json:"id"`
	// The amount of this currency.
	Value int `json:"value"`
}

// This resource returns information about player accounts.
// This endpoint is only accessible with a valid API key.
func (r *Requestor) Account(account *Account) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionAccount).
		request("/account", nil, &account)
	return r
}

// This resource returns an account's progress towards all their achievements.
// The response shows every achievement that the account has progress on by ID
// and how far the player has progressed. For each achievement,
// the following object is given:
func (r *Requestor) AccountAchievements(achievements *[]*AccountAchievement) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionProgression).
		request("/account/achievements", nil, &achievements)
	return r
}

// This resource returns the items stored in a player's vault
// (not including material storage).
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array of objects, each representing an item slot
// in the vault.
// If a slot is empty, it will return null. The amount of slots/bank tabs is
// implied by the length of the array.
func (r *Requestor) AccountBank(accountBank *[]*AccountInventoryItem) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionInventory).
		request("/account/bank", nil, &accountBank)
	return r
}

// This resource returns the templates stored in a player's build storage.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array of objects, each representing a template
// slot in the build storage. The amount of templates is implied by the
// length of the array.
func (r *Requestor) AccountBuildStorage(accountBuildStorage *[]*AccountBuildStorage) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionAccount).
		request("/account/buildstorage", nil, &accountBuildStorage)
	return r
}

// This resource returns information about time-gated
// recipes that have been crafted by the account since daily-reset.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array of strings, each representing a time-gated
// recipe name that can be resolved against /v2/dailycrafting.
// If no timed-gated recipes have been crafted since daily-reset by the account,
// it will return an empty array ([]).
func (r *Requestor) AccountDailyCrafting(accountDailyCraft *[]string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionProgression).
		request("/account/dailycrafting", nil, &accountDailyCraft)
	return r
}

// This resource returns the dungeons completed since daily dungeon reset.
// The endpoint returns an array, each value being the ID of a dungeon path
// that can be resolved against /v2/dungeons. Note that this ID indicates a
// path and not the dungeon itself.
func (r *Requestor) AccountDungeons(accountDungeons *[]string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionProgression).
		request("/account/dungeons", nil, &accountDungeons)
	return r
}

// This resource returns the unlocked dyes of the account.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array, each value being the ID
// of a color resolved against /v2/colors.
func (r *Requestor) AccountDyes(accountDyes *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/dyes", nil, &accountDyes)
	return r
}

// This resource returns the player's unlocked emotes.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array of strings, each representing an emote.
func (r *Requestor) AccountEmotes(accountEmotes *[]string) *Requestor {
	r.request("/account/emotes", nil, &accountEmotes)
	return r
}

// This resource returns information about finishers that are unlocked
// for an account.
// This request will return an array of objects
func (r *Requestor) AccountFinishers(accountFinishers *[]*AccountFinisher) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/finishers", nil, &accountFinishers)
	return r
}

// This resource returns information about gliders
// that are unlocked for an account.
// This request will return an array of integer values
// resolvable against /v2/gliders.
func (r *Requestor) AccountGliders(accountGliders *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/gliders", nil, &accountGliders)
	return r
}

// This resource returns a list of available sub-endpoints.
// This request will return an array of the endpoints that may be requested.
func (r *Requestor) AccountHome(accountHome *[]string) *Requestor {
	r.request("/account/home", nil, &accountHome)
	return r
}

// This resource returns information about unlocked home instance cats.
// This request will return an array of integers.
// Each integer represents the id of a particular cat that can be resolved
// against /v2/home/cats.
func (r *Requestor) AccountHomeCats(accountHomeCats *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/home/cats", nil, &accountHomeCats)
	return r
}

// This resource returns information about unlocked home instance nodes.
// This request will return an array of strings. Each string represents
// the id of a particular node that can be resolved against /v2/home/nodes.
func (r *Requestor) AccountHomeNodes(accountHomeNodes *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/home/nodes", nil, &accountHomeNodes)
	return r
}

// This resource returns the shared inventory slots in an account.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array of objects, each representing an item slot in
// the shared inventory. If a slot is empty, it will return null.
// The amount of slots is implied by the length of the array.
func (r *Requestor) AccountInventory(accountInventory *[]*AccountInventoryItem) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionInventory).
		request("/account/inventory", nil, &accountInventory)
	return r
}

// This resource returns information about the Legendary Armory
// items that are unlocked for an account.
// This request will return an array of objects
func (r *Requestor) AccountLegendaryArmory(accountLegendaryArmory *[]*AccountLegendaryArmory) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionInventory, TokenPermissionUnlocks).
		request("/account/legendaryarmory", nil, &accountLegendaryArmory)
	return r
}

// This resource returns the total amount of luck consumed
// on an account. This endpoint is only accessible with a valid API key.
// The endpoint returns an array with a singular object
// NOTE: The response is an array due to the way luck is stored
// internally on ArenaNET servers.
func (r *Requestor) AccountLuck(accountLuck *[]*AccountLuck) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionProgression, TokenPermissionUnlocks).
		request("/account/legendaryarmory", nil, &accountLuck)
	return r
}

// This resource returns information about mail carriers that are
// unlocked for an account.
// This request will return an array of integer values that can be
// resolved against /v2/mailcarriers.
func (r *Requestor) AccountMailCarriers(accountMailCarriers *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/mailcarriers", nil, &accountMailCarriers)
	return r
}

// This resource returns information about Hero's Choice Chests
// acquired by the account since daily-reset. This endpoint is only
// accessible with a valid API key.
// The endpoint returns an array of strings, each representing a
// Hero's Choice Chest that can be resolved against /v2/mapchests.
// If no Hero's Choice Chest have been acquired since
// daily-reset by the account, it will return an empty array ([]).
func (r *Requestor) AccountMapChests(accountMapChests *[]string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionProgression).
		request("/account/mapchests", nil, &accountMapChests)
	return r
}

// This resource returns information about masteries that are unlocked
// for an account.
// A tallied up total of the account's mastery points can be
// found at /v2/account/mastery/points.
// This request will return an array of objects
func (r *Requestor) AccountMasteries(accountMasteries *[]*AccountMastery) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionProgression).
		request("/account/masteries", nil, &accountMasteries)
	return r
}

// This resource returns information about the total amount of masteries
// that are unlocked for an account. A detailed mastery track completion
// break down is available at /v2/account/masteries.
// This request will return an object
func (r *Requestor) AccountMasteryPoints(accountMasteryPoints *[]*AccountMasteryPoint) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionProgression).
		request("/account/mastery/points", nil, &accountMasteryPoints)
	return r
}

// This resource returns the materials stored in a player's vault.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array of objects, each representing a material
// that can be stored in the vault. Every material will be returned,
// even if they have a count of 0.
func (r *Requestor) AccountMaterials(accountMaterials *[]*AccountMaterial) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionInventory).
		request("/account/materials", nil, &accountMaterials)
	return r
}

// This resource returns the unlocked miniatures of the account.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array, each value being the ID of a
// miniature that can be resolved against /v2/minis.
func (r *Requestor) AccountMinis(accountMinis *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/minis", nil, &accountMinis)
	return r
}

// This resource returns a list of available sub-endpoints.
// This request will return an array of the endpoints that may be requested.
func (r *Requestor) AccountMounts(accountMounts *[]string) *Requestor {
	r.request("/account/mounts", nil, &accountMounts)
	return r
}

// This resource returns the unlocked mount skins of the account.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array of numbers which can be
// resolved against to /v2/mounts/skins.
func (r *Requestor) AccountMountsSkins(accountMountsSkins *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/mounts/skins", nil, &accountMountsSkins)
	return r
}

// This resource returns the unlocked mounts of the account.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array of strings (mount names) which
// can be compared to /v2/mounts/types.
func (r *Requestor) AccountMountsTypes(accountMountsTypes *[]string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/mounts/types", nil, &accountMountsTypes)
	return r
}

// This resource returns information about novelties that are unlocked
// for an account.
// This request will return an array of integer values
// resolvable against /v2/novelties.
func (r *Requestor) AccountNovelties(accountNovelties *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/novelties", nil, &accountNovelties)
	return r
}

// This resource returns information about outfits that are
// unlocked for an account.
// This request will return an array of integer values resolvable
// against /v2/outfits.
func (r *Requestor) AccountOutfils(accountOutfils *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/outfils", nil, &accountOutfils)
	return r
}

// This resource returns information about pvp heroes that
// are unlocked for an account.
// This request will return an array of integer values
// resolvable against /v2/pvp/heroes.
func (r *Requestor) AccountPvpHeroes(accountPvpHeroes *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/pvp/heroes", nil, &accountPvpHeroes)
	return r
}

// This resource returns the completed raid encounters
// since weekly raid reset.
// The endpoints returns an array, each value being the ID
// of a raid encounter that can be resolved against /v2/raids.
// Note that this ID indicates an encounter and not the raid wing itself.
func (r *Requestor) AccountRaids(accountRaids *[]string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionProgression).
		request("/account/raids", nil, &accountRaids)
	return r
}

// This resource returns information about recipes that
// are unlocked for an account.
// The endpoint returns an array, each value being the ID of a
// recipe that can be resolved against /v2/recipes.
func (r *Requestor) AccountReceipes(accountReceipes *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/recipes", nil, &accountReceipes)
	return r
}

// This resource returns the unlocked skins of the account.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array, each value being the ID of a skin
// that can be resolved against /v2/skins.
func (r *Requestor) AccountSkins(accountSkins *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/skins", nil, &accountSkins)
	return r
}

// This resource returns information about titles that are
// unlocked for an account.
// This request will return an array of integer values
// resolvable against /v2/titles.
func (r *Requestor) AccountTitles(accountTitles *[]int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionUnlocks).
		request("/account/titles", nil, &accountTitles)
	return r
}

// This resource returns the currencies of the account.
// This endpoint is only accessible with a valid API key.
// The endpoint returns an array of objects, each representing a currency.
func (r *Requestor) AccountWallet(accountWallet *[]*AccountCurrency) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionWallet).
		request("/account/wallet", nil, &accountWallet)
	return r
}

// This resource returns information about which world bosses have been
// killed by the account since daily-reset. This endpoint is only
// accessible with a valid API key.
// The endpoint returns an array of strings, each representing a world boss
// name that can be compared with the full array within /v2/worldbosses.
// If no bosses have been killed since daily-reset by the account,
// it will return an empty array ([]).
func (r *Requestor) AccountWorldBosses(accountWorldBosses *[]string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionProgression).
		request("/account/worldbosses", nil, &accountWorldBosses)
	return r
}
