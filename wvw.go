//go:generate easytags $GOFILE
package gw2api

import "time"

type WvwAbility struct {
	// The id of the abilities.
	ID int `json:"id"`
	// The given name for the WvW ability.
	Name string `json:"name"`
	// The given description for the WvW ability.
	Description string `json:"description"`
	// The uri for the ability's icon.
	Icon  string           `json:"icon"`
	Ranks []WvwAbilityRank `json:"ranks"`
}

type WvwAbilityRank struct {
	// The cost in WvW experience points to purchase the ability.
	Cost int `json:"cost"`
	// The effect given to players for obtaining the given ability rank.
	Effect string `json:"effect"`
}

type WvwMatch struct {
	WvwMatchOverview
	// An object containing the score of the three servers, under the values
	// red, blue, and green.
	Scores RGBStat `json:"scores"`
	// An object containing the total deaths of the three servers, under the
	// values red, blue, and green.
	Deaths RGBStat `json:"deaths"`
	// An object containing the total kills of the three servers, under the
	// values red, blue, and green.
	Kills RGBStat `json:"kills"`
	// An object containing the victory points of the three servers, under the
	// values red, blue, and green.
	VictoryPoints RGBStat `json:"victory_points"`
	// A list of objects containing detailed information about each of the four
	// maps.
	Maps       []WvwMatchMap      `json:"maps"`
	Skirmishes []WvwMatchSkirmish `json:"skirmishes"`
}

type WvwMatchOverview struct {
	// The WvW match id.
	ID string `json:"id"`
	// The starting time of the matchup. (ISO-8601 Standard)
	StartTime time.Time `json:"start_time"`
	// The ending time of the matchup. (ISO-8601 Standard)
	EndTime time.Time `json:"end_time"`
	// An object containing the IDs of the three primary matchup worlds.
	Worlds RGBStat `json:"worlds"`
	// An object containing an array of objects with the world IDs of the three
	// servers, under the values red, blue, and green.
	AllWorlds WvwMatchAllWorlds `json:"all_worlds"`
}

type WvwMatchMap struct {
	// The map id
	ID int `json:"id"`
	// The identifier for the map. Can be either RedHome, GreenHome or BlueHome
	// for the borderlands or Center for Eternal Battlegrounds.
	Type string `json:"type"`
	//  An object containing the score of the three servers for only the
	// specified map, under the values red, blue, and green.
	Scores RGBStat `json:"scores"`
	// An object containing the total kills of the three servers for only
	// the specified map, under the values red, blue, and green.
	Kills RGBStat `json:"kills"`
	// An object containing the total deaths of the three servers for only
	// the specified map, under the values red, blue, and green.
	Deaths RGBStat `json:"deaths"`
	// A list of objective objects for this map.
	Objectives []WvwMatchMapObjective `json:"objectives"`
	// A list of all bonuses being granted by this map. If no player team
	// owns a bonus from the map, this list is empty.
	Bonuses []WvwMatchMapBonus `json:"bonuses"`
}

type WvwMatchMapObjective struct {
	// The objective id.
	ID string `json:"id"`
	// The objective type, with possible values Spawn, Camp, Ruins, Tower,
	// Keep, Castle, and Mercenary (Dredge/Krait/Frogs).
	Type string `json:"type"`
	// The current owner of the objective. Can be any one of Red, Green,
	// Blue or Neutral.
	Owner string `json:"owner"`
	// The time at which this objective was last captured by a server.
	// (ISO-8601 Standard)
	LastFlipped string `json:"last_flipped"`
	// The guild id of the guild currently claiming the objective, or
	// null if not claimed. (Not present for unclaimable objectives.)
	ClaimedBy string `json:"claimed_by"`
	// The time the objective was claimed by the claimed_by guild
	// (ISO-8601 Standard), or null if not claimed.
	// (Not present for unclaimable objectives.)
	ClaimedAt *time.Time `json:"claimed_at"`
	// The amount of points per tick the given objective yields.
	PointsTick int `json:"points_tick"`
	// The amount of points awarded for capturing the objective.
	PointsCapture int `json:"points_capture"`
	// An array of ids, resolvable against v2/guild/upgrades, showing which
	// guild upgrades are currently slotted.
	GuildUpgrades []int `json:"guild_upgrades"`
	// The total amount of yak shipments delivered to the objective. Not
	// limited to the shipments within the current tier only.
	YaksDelivered int `json:"yaks_delivered"`
}

type WvwMatchMapBonus struct {
	// A shorthand name for the bonus. Currently can only be Bloodlust.
	Type string `json:"type"`
	// The color representing which world group owns the bloodlust.
	Owner string `json:"owner"`
}

type WvwMatchSkirmish struct {
	// id number) - The skirmish id
	ID int `json:"id"`
	// scores object) - Object containing total scores for each team color,
	// under the values red, blue, and green.
	Scores RGBStat `json:"scores"`
	// map_scores array of objects) - Contains the map specific scores for the
	// specific skirmish.
	MapScores []WvwMatchSkirmishMapScore `json:"map_scores"`
}

type WvwMatchSkirmishMapScore struct {
	// Which map is being looked at, can have the values "Center", "RedHome",
	// "BlueHome", or "GreenHome"
	Type string `json:"type"`
	// Object containing total scores for each team color on the selected map,
	// under the values red, blue, and green.
	Scores RGBStat `json:"scores"`
}

type RGBStat struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

type WvwMatchAllWorlds struct {
	Red   []int `json:"red"`
	Green []int `json:"green"`
	Blue  []int `json:"blue"`
}

type WvwMatchScore struct {
	// The WvW match id.
	ID string `json:"id"`
	// An object containing the score of the three servers, under the values
	// red, blue, and green.
	Scores RGBStat `json:"scores"`
	// An object containing the victory points of the three servers, under the
	// values red, blue, and green.
	VictoryPoints RGBStat `json:"victory_points"`
	// A list of objects containing detailed information about each of the four
	// maps.
	Maps       []WvwMatchScoreMap `json:"maps"`
	Skirmishes []WvwMatchSkirmish `json:"skirmishes"`
}

type WvwMatchScoreMap struct {
	// The map id
	ID int `json:"id"`
	// The identifier for the map. Can be either RedHome, GreenHome or BlueHome
	// for the borderlands or Center for Eternal Battlegrounds.
	Type string `json:"type"`
	//  An object containing the score of the three servers for only the
	// specified map, under the values red, blue, and green.
	Scores RGBStat `json:"scores"`
}

type WvwMatchStat struct {
	// The WvW match id.
	ID string `json:"id"`
	// An object containing the total deaths of the three servers, under the
	// values red, blue, and green.
	Deaths RGBStat `json:"deaths"`
	// An object containing the total kills of the three servers, under the
	// values red, blue, and green.
	Kills RGBStat `json:"kills"`
	// A list of objects containing detailed information about each of the four
	// maps.
	Maps []WvwMatchStatMap `json:"maps"`
}

type WvwMatchStatMap struct {
	// The map id
	ID int `json:"id"`
	// The identifier for the map. Can be either RedHome, GreenHome or BlueHome
	// for the borderlands or Center for Eternal Battlegrounds.
	Type string `json:"type"`
	// An object containing the total kills of the three servers for only
	// the specified map, under the values red, blue, and green.
	Kills RGBStat `json:"kills"`
	// An object containing the total deaths of the three servers for only
	// the specified map, under the values red, blue, and green.
	Deaths RGBStat `json:"deaths"`
}

type WvwObjective struct {
	// The objective id.
	ID string `json:"id"`
	// The name of the objective.
	Name string `json:"name"`
	// The type of the objective. Possible values include:
	//   Camp
	//   Castle
	//   Keep
	//   Mercenary
	//   Tower
	//   Ruins
	//   Resource
	//   Generic
	//   Spawn
	Type string `json:"type"`
	// The map sector the objective can be found in. (See /v2/continents.)
	SectorID int `json:"sector_id"`
	// The ID of the map that this objective can be found on.
	MapId int `json:"map_id"`
	// The map that this objective can be found on.
	// One of:
	//   GreenHome, BlueHome, RedHome, Center, or EdgeOfTheMists.
	MapType string `json:"map_type"`
	// An array of three numbers representing the X, Y and Z coordinates of
	// the objectives marker on the map.
	Coord []int `json:"coord"`
	// An array of two numbers representing the X and Y coordinates of the
	// sector centroid.
	LabelCoord []int `json:"label_coord"`
	// The icon link
	Marker string `json:"marker"`
	// The chat code for the observed objective.
	ChatLink string `json:"chat_link"`
	// The upgrade id to be resolved against v2/wvw/upgrades.
	UpgradeID int `json:"upgrade_id"`
}

type WvwRank struct {
	// The id of the rank.
	ID string `json:"id"`
	// The given title for the WvW rank.
	Title string `json:"title"`
	// The minimum WvW level required to be at this rank.
	MinRank int `json:"min_rank"`
}

type WvwUpgrade struct {
	// The upgrade id.
	ID    int              `json:"id"`
	Tiers []WvwUpgradeTier `json:"tiers"`
}

type WvwUpgradeTier struct {
	// The name of the upgrade tier.
	Name string `json:"name"`
	// The number of required yaks.
	YaksRequired int                     `json:"yaks_required"`
	Upgrades     []WvwUpgradeTierUpgrade `json:"upgrades"`
}

type WvwUpgradeTierUpgrade struct {
	// The name of the upgrade tier.
	Name string `json:"name"`
	// The given description for this upgrade.
	Description string `json:"description"`
	// The url/image link for the upgrade's icon.
	Icon string `json:"icon"`
}

// This resource returns information about the available abilities in the
// World versus World game mode.
// Return an array of ids for each type of currency.
func (r *Requestor) WvwAbilityIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/wvw/abilities", &pointer)
	return r
}

// This resource returns information about the available abilities in the
// World versus World game mode.
// Return a list of response objects
func (r *Requestor) WvwAbilities(pointer *[]*WvwAbility, ids ...int) *Requestor {
	r.collection("/wvw/abilities", &pointer, ids)
	return r
}

// This resource returns information about the available abilities in the
// World versus World game mode.
// Return an object
func (r *Requestor) WvwAbility(pointer *WvwAbility, id int) *Requestor {
	r.singleton("/wvw/abilities", &pointer, id)
	return r
}

// This resource returns information about the available matches in the
// World versus World game mode.
// Return an array of ids for each type of currency.
func (r *Requestor) WvwMatchIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/wvw/matches", &pointer)
	return r
}

// This resource returns information about the available matches in the
// World versus World game mode.
// Return a list of response objects
func (r *Requestor) WvwMatches(pointer *[]*WvwMatch, ids ...string) *Requestor {
	r.collection("/wvw/matches", &pointer, ids)
	return r
}

// This resource returns information about the available matches in the
// World versus World game mode.
// Return an object
func (r *Requestor) WvwMatch(pointer *WvwMatch, id string) *Requestor {
	r.singleton("/wvw/matches", &pointer, id)
	return r
}

// This resource returns information about the available matches overview in the
// World versus World game mode.
// Return an array of ids for each type of currency.
func (r *Requestor) WvwMatchOverviewIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/wvw/matches/overview", &pointer)
	return r
}

// This resource returns information about the available matches overview in the
// World versus World game mode.
// Return a list of response objects
func (r *Requestor) WvwMatchOverviews(pointer *[]*WvwMatchOverview, ids ...string) *Requestor {
	r.collection("/wvw/matches/overview", &pointer, ids)
	return r
}

// This resource returns information about the available matches overview in the
// World versus World game mode.
// Return an object
func (r *Requestor) WvwMatchOverview(pointer *WvwMatchOverview, id string) *Requestor {
	r.singleton("/wvw/matches/overview", &pointer, id)
	return r
}

// This resource returns information about the available matches scores in the
// World versus World game mode.
// Return an array of ids for each type of currency.
func (r *Requestor) WvwMatchScoreIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/wvw/matches/scores", &pointer)
	return r
}

// This resource returns information about the available matches scores in the
// World versus World game mode.
// Return a list of response objects
func (r *Requestor) WvwMatchScores(pointer *[]*WvwMatchScore, ids ...string) *Requestor {
	r.collection("/wvw/matches/scores", &pointer, ids)
	return r
}

// This resource returns information about the available matches scores in the
// World versus World game mode.
// Return an object
func (r *Requestor) WvwMatchScore(pointer *WvwMatchScore, id string) *Requestor {
	r.singleton("/wvw/matches/scores", &pointer, id)
	return r
}

// This resource returns information about the available matches stats in the
// World versus World game mode.
// Return an array of ids for each type of currency.
func (r *Requestor) WvwMatchStatIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/wvw/matches/stats", &pointer)
	return r
}

// This resource returns information about the available matches stats in the
// World versus World game mode.
// Return a list of response objects
func (r *Requestor) WvwMatchStats(pointer *[]*WvwMatchStat, ids ...string) *Requestor {
	r.collection("/wvw/matches/stats", &pointer, ids)
	return r
}

// This resource returns information about the available matches stats in the
// World versus World game mode.
// Return an object
func (r *Requestor) WvwMatchStat(pointer *WvwMatchStat, id string) *Requestor {
	r.singleton("/wvw/matches/stats", &pointer, id)
	return r
}

// This resource returns details about World vs. World objectives such
// as camps, towers, and keeps.
// Return an array of ids for each type of currency.
func (r *Requestor) WvwObjectiveIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/wvw/objectives", &pointer)
	return r
}

// This resource returns details about World vs. World objectives such
// as camps, towers, and keeps.
// Return a list of response objects
func (r *Requestor) WvwObjectives(pointer *[]*WvwObjective, ids ...string) *Requestor {
	r.collection("/wvw/objectives", &pointer, ids)
	return r
}

// This resource returns details about World vs. World objectives such
// as camps, towers, and keeps.
// Return an object
func (r *Requestor) WvwObjective(pointer *WvwObjective, id string) *Requestor {
	r.singleton("/wvw/objectives", &pointer, id)
	return r
}

// This resource returns information about the available ranks in the
// World versus World game mode.
// Return an array of ids for each type of currency.
func (r *Requestor) WvwRankIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/wvw/ranks", &pointer)
	return r
}

// This resource returns information about the available ranks in the
// World versus World game mode.
// Return a list of response objects
func (r *Requestor) WvwRanks(pointer *[]*WvwRank, ids ...int) *Requestor {
	r.collection("/wvw/ranks", &pointer, ids)
	return r
}

// This resource returns information about the available ranks in the
// World versus World game mode.
// Return an object
func (r *Requestor) WvwRank(pointer *WvwRank, id int) *Requestor {
	r.singleton("/wvw/ranks", &pointer, id)
	return r
}

// This resource returns details about available World vs. World upgrades
// for objectives such as camps, towers, and keeps.
// Return an array of ids for each type of currency.
func (r *Requestor) WvwUpgradeIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/wvw/upgrades", &pointer)
	return r
}

// This resource returns details about available World vs. World upgrades
// for objectives such as camps, towers, and keeps.
// Return a list of response objects
func (r *Requestor) WvwUpgrades(pointer *[]*WvwUpgrade, ids ...int) *Requestor {
	r.collection("/wvw/upgrades", &pointer, ids)
	return r
}

// This resource returns details about available World vs. World upgrades
// for objectives such as camps, towers, and keeps.
// Return an object
func (r *Requestor) WvwUpgrade(pointer *WvwUpgrade, id int) *Requestor {
	r.singleton("/wvw/upgrades", &pointer, id)
	return r
}
