//go:generate easytags $GOFILE
package gw2api

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type Achievement struct {
	// The achievement id.
	ID int `json:"id"`
	// The achievement icon.
	Icon string `json:"icon"`
	// The achievement name.
	Name string `json:"name"`
	// The achievement description.
	Description string `json:"description"`
	// The achievement requirement as listed in-game.
	Requirement string `json:"requirement"`
	// The achievement description prior to unlocking it.
	LockedText string `json:"locked_text"`
	// The achievement type. Possible values:
	//   `Default` - A default achievement.
	//   `ItemSet` - Achievement is linked to Collections
	Type string `json:"type"`
	// Achievement categories. Possible values:
	//   `Pvp`                  - can only get progress in PvP or WvW
	//   `CategoryDisplay`      - is a meta achievement
	//   `MoveToTop`            - affects in-game UI collation
	//   `IgnoreNearlyComplete` - doesnt appear in the "nearly complete" UI
	//   `Repeatable`           - can be repeated multiple times
	//   `Hidden`               - hidden achievement; must fulfil unlock
	//                            requirements before making progress or
	//                            showing in the hero panel
	//   `RequiresUnlock`       - must fulfil unlock requirements before making
	//                            progress but will show in the hero panel
	//                            before unlocking
	//   `RepairOnLogin`        - unknown
	//   `Daily`                - Flags an achievement as resetting daily.
	//   `Weekly`               - Flags an achievement as resetting weekly.
	//   `Monthly`              - Flags an achievement as resetting monthly.
	//   `Permanent`            - Flags an achievement as progress never reseting.
	Flags []string `json:"flags"`
	// Describes the achievement's tiers.
	Tiers []AchievementTier `json:"tiers"`
	// Contains an array of achievement ids required to progress the given
	// achievement.
	Prereqisites []int `json:"prereqisites"`
	// Describes the rewards given for the achievement.
	Rewards  AchievementReward `json:"rewards"`
	Bits     AchievementBit    `json:"bits"`
	PointCap int               `json:"point_cap"`
}

type AchievementTier struct {
	// The number of "things" (achievement-specific) that must be
	// completed to achieve this tier.
	Count int `json:"count"`
	// The amount of AP awarded for completing this tier.
	Points int `json:"points"`
}

type AchievementReward struct {
	// The type of reward. Additional fields appear for different values of `type`
	Type string `json:"type"`
	// Possible values for Count :
	//   `type == "Coins"` - The number of Coins to be rewarded.
	//   `type == "Item"` - The number of id to be rewarded.
	Count int `json:"count"`
	// Possible values for ID :
	//   `type == "Item"`    - The item ID to be rewarded.
	//   `type == "Mastery"` - The mastery point ID to be rewarded.
	//   `type == "Title"`   - The title id.
	ID int `json:"id"`
	// Possible values for Region :
	//   `type == "Mastery"` - The region the Mastery Point applies to.
	//                         Either `Tyria`, `Maguuma`, `Desert` or `Tundra`.
	Region string `json:"region"`
}

type AchievementBit struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type AchievementsCategory struct {
	// The category's ID.
	ID int `json:"id"`
	// The category name.
	Name string `json:"name"`
	// The category description.
	Description string `json:"description"`
	// A number describing where to sort this category among other the other
	// categories in its group. Lowest numbers go first, highest numbers go last.
	Order int `json:"order"`
	// A URL to an image for the icon of the category.
	Icon string `json:"icon"`
	// An array containing a number of achievement IDs that this category
	// contains. (See /v2/achievements.)
	Achievements []int `json:"achievements"`
}

type AchievementsDailyStructure struct {
	Pve      []AchievementsDaily
	Pvp      []AchievementsDaily
	Wvw      []AchievementsDaily
	Fractals []AchievementsDaily
	Special  []AchievementsDaily
}

type AchievementsDaily struct {
	// The achievement id.
	ID int
	// Describes the level requirement for the daily to appear.
	Level AchievementsDailyLevel
	// Describes if a daily requires a Guild Wars 2 campaign or not.
	RequiredAccess AchievementsDailyRequiredAccess
}

type AchievementsDailyLevel struct {
	// The minimum level. Any character below this level will not
	// see this daily achievement.
	Min int
	// The maximum level. Any character above this level will not
	// see this daily achievement.
	Max int
}

type AchievementsDailyRequiredAccess struct {
	// A Guild Wars 2 campaign.
	// Can either be `GuildWars2`, `HeartOfThorns` or `PathOfFire`.
	Product string
	// The condition if a account can or cannot see this daily achievement.
	// Can either be `HasAccess` or `NoAccess`.
	Condition string
}

type AchievementsGroup struct {
	// The group's GUID.
	ID string
	// The group name.
	Name string
	// The group description.
	Description string
	// A number describing where to sort this group among other groups.
	// Lowest numbers go first, highest numbers go last.
	Order int
	// An array containing a number of category IDs that this group contains.
	// @see `/v2/achievements/categories`.
	Categories []int
}

// This resource returns all achievements in the game,
// including localized names and icons.
// A list of all ids is returned.
func (r *Requestor) AchievementIDs(s []int) *Requestor {
	r.request("/achievements", nil, &s)
	return r
}

// This resource returns all achievements in the game,
// When multiple ids are requested using the ids
// parameter, a list of response objects is returned.
func (r *Requestor) Achievements(achievements *[]*Achievement, ids ...int) *Requestor {
	if len(ids) == 0 {
		r.err = errors.New("at least one id must be given")
		return r
	}

	var urlValues url.Values
	sIds := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	urlValues = url.Values{"ids": strings.Split(sIds, ",")}

	r.request("/achievements", urlValues, &achievements)
	return r
}

// This resource returns an achievement in the game by her ID,
func (r *Requestor) Achievement(achievement *Achievement, id int) *Requestor {
	r.request("/achievements", url.Values{"id": []string{fmt.Sprint(id)}}, &achievement)
	return r
}

// This resource returns all achievements categories in the game,
// A list of all ids is returned.
func (r *Requestor) AchievementsCategoryIDs(s []int) *Requestor {
	r.request("/achievements/categories", nil, &s)
	return r
}

// This resource returns all achievements categories in the game,
// When multiple ids are requested using the ids
// parameter, a list of response objects is returned.
func (r *Requestor) AchievementsCategories(achievementsCategory *[]*AchievementsCategory, ids ...int) *Requestor {
	if len(ids) == 0 {
		r.err = errors.New("at least one id must be given")
		return r
	}

	var urlValues url.Values
	sIds := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	urlValues = url.Values{"ids": strings.Split(sIds, ",")}

	r.request("/achievements/categories", urlValues, &achievementsCategory)
	return r
}

// This resource returns an achievement category in the game by her ID,
func (r *Requestor) AchievementsCategory(achievementsCategory *AchievementsCategory, id int) *Requestor {
	r.request("/achievements/categories", url.Values{"id": []string{fmt.Sprint(id)}}, &achievementsCategory)
	return r
}

// This resource returns the current set of daily achievements.
// When requested, the endpoint will return an object containing
// 5 sub-objects: `pve`, `pvp`, `wvw`, `fractals`, and `special`.
// Each sub-object contains an array of objects describing the daily
// achievements for that category. The special sub-object is for any current
// temporary content like festival dailies
func (r *Requestor) AchievementsDaily(achievementsDaily *AchievementsDailyStructure) *Requestor {
	r.request("/achievements/daily", nil, &achievementsDaily)
	return r
}

// This resource returns the next set of daily achievements.
// When requested, the endpoint will return an object containing
// 5 sub-objects: `pve`, `pvp`, `wvw`, `fractals`, and `special`.
// Each sub-object contains an array of objects describing the daily
// achievements for that category. The special sub-object is for any current
// temporary content like festival dailies
func (r *Requestor) AchievementsDailyTomorrow(achievementsDaily *AchievementsDailyStructure) *Requestor {
	r.request("/achievements/daily/tomorrow", nil, &achievementsDaily)
	return r
}

// This resource returns all the top-level groups for achievements.
// A list of all ids is returned.
func (r *Requestor) AchievementsGroupIDs(s []string) *Requestor {
	r.request("/achievements/groups", nil, &s)
	return r
}

// This resource returns all the top-level groups for achievements.
// When multiple ids are requested using the ids
// parameter, a list of response objects is returned.
func (r *Requestor) AchievementsGroups(achievementsGroup *[]*AchievementsGroup, ids ...string) *Requestor {
	if len(ids) == 0 {
		r.err = errors.New("at least one id must be given")
		return r
	}

	r.request("/achievements/groups", url.Values{"ids": ids}, &achievementsGroup)
	return r
}

// This resource returns a top level group for achievements in the game by her ID,
func (r *Requestor) AchievementsGroup(achievementsGroup *AchievementsGroup, id string) *Requestor {
	r.request("/achievements/groups", url.Values{"id": []string{id}}, &achievementsGroup)
	return r
}
