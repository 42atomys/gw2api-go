//go:generate easytags $GOFILE
package gw2api

import (
	"fmt"
	"time"
)

type Character struct {
	// The character's name.
	Name string `json:"name"`
	// The character's race. Possible values:
	//   Asura
	//   Charr
	//   Human
	//   Norn
	//   Sylvari
	Race string `json:"race"`
	// The character's gender. Possible values:
	//   Male
	//   Female
	Gender string `json:"gender"`
	// The character's profession. Possible values:
	//   Elementalist
	//   Engineer
	//   Guardian
	//   Mesmer
	//   Necromancer
	//   Ranger
	//   Revenant
	//   Thief
	//   Warrior
	Profession string `json:"profession"`
	// The character's level.
	Level int `json:"level"`
	// The guild ID of the character's currently represented guild.
	Guild string `json:"guild"`
	// The amount of seconds this character was played.
	Age int `json:"age"`
	// ISO 8601 representation of the character's creation time.
	Created time.Time `json:"created"`
	// The amount of times this character has been defeated.
	Deaths int `json:"deaths"`
	// The currently selected title for the character. References /v2/titles.
	Title int `json:"title"`
}

type CharacterBuildTab struct {
	// The "id" of this tab. (The position at which it resides.)
	Tab int `json:"tab"`
	// Whether or not this is the tab selected on the character currently
	IsActive bool `json:"is_active"`
	// Contains detailed information about the build.
	Build AccountBuildStorage `json:"build"`
}

type CharacterCrafting struct {
	// The name of the discipline. Possible values:
	//   Armorsmith
	//   Artificer
	//   Chef
	//   Huntsman
	//   Jeweler
	//   Leatherworker
	//   Scribe
	//   Tailor
	//   Weaponsmith
	Discipline string `json:"discipline"`
	// The current crafting level for the given discipline and character
	Rating int `json:"rating"`
	// Describes if the given discipline is currently active or not on
	// the character.
	Active bool `json:"active"`
}

type CharacterEquipment struct {
	// The item id, resolvable against /v2/items
	ID int `json:"id"`
	// The equipment slot in which the item is slotted. Possible values:
	//   HelmAquatic
	//   Backpack
	//   Coat
	//   Boots
	//   Gloves
	//   Helm
	//   Leggings
	//   Shoulders
	//   Accessory1
	//   Accessory2
	//   Ring1
	//   Ring2
	//   Amulet
	//   WeaponAquaticA
	//   WeaponAquaticB
	//   WeaponA1
	//   WeaponA2
	//   WeaponB1
	//   WeaponB2
	//   Sickle
	//   Axe
	//   Pick
	Slot string `json:"slot"`
	// returns an array of infusion item ids which can be resolved
	// against /v2/items
	Infusions []int `json:"infusions"`
	// returns an array of upgrade component item ids which can be resolved
	// against /v2/items
	Upgrades []int `json:"upgrades"`
	// Skin id for the given equipment piece. Can be resolved against /v2/skins
	Skin int `json:"skin"`
	// Contains information on the stats chosen if the item offers an
	// option for stats/prefix
	Stats ItemStats `json:"stats"`
	// describes which kind of binding the item has. Possible values:
	//   Character
	//   Account
	Binding string `json:"binding"`
	// The amount of charges remaining on the item.
	Charges int `json:"charges"`
	// Name of the character the item is bound to.
	BoundTo string `json:"bound_to"`
	// Array of selected dyes for the equipment piece. Values default to null
	// if no dye is selected. Colors can be resolved against v2/colors
	Dyes []int `json:"dyes"`
	// describes where this item is stored. Possible values:
	//   Equipped - equipped in the active tab
	//   Armory - equipped in inactive tabs
	//   EquippedFromLegendaryArmory - if the item is stored in the account-wide
	//                                 legendary armory.
	Location string `json:"location"`
	// Identifies which tabs this particular item is reused in.
	Tabs []int `json:"tabs"`
}

type CharacterInventoryBag struct {
	// The bag's item id which can be resolved against /v2/items
	ID int `json:"id"`
	// The amount of slots available with this bag.
	Size int `json:"size"`
	// Contains one object structure per item, object is null if no item
	// is in the given bag slot.
	Inventory []InventoryItem `json:"inventory"`
}

type CharacterSkills struct {
	// contains the information on each slotted utility for PvE
	Pve CharacterSkill `json:"pve"`
	// contains the information on each slotted utility for PvP
	Pvp CharacterSkill `json:"pvp"`
	// contains the information on each slotted utility for WvW
	Wvw CharacterSkill `json:"wvw"`
}

type CharacterSkill struct {
	// contains the skill id for the heal skill, resolvable against /v2/skills.
	Heal int `json:"heal"`
	// Each integer corresponds to a skill id for the equipped utilities,
	// resolvable against /v2/skills.
	Utilities []int `json:"utilities"`
	// Contains the skill id for the elite skill, resolvable against /v2/skills.
	Elite int `json:"elite"`
	// (Revenant only) - each string corresponds to a Revenant legend,
	// resolvable against /v2/legends.
	Legens []int `json:"legens"`
}

type CharacterSpecializations struct {
	// contains the information on each slotted specialization and trait for PvE
	Pve CharacterSpecialization `json:"pve"`
	// contains the information on each slotted specialization and trait for PvP
	Pvp CharacterSpecialization `json:"pvp"`
	// contains the information on each slotted specialization and trait for WvW
	Wvw CharacterSpecialization `json:"wvw"`
}

type CharacterSpecialization struct {
	// Specialization id, can be resolved against /v2/specializations.
	ID int `json:"id"`
	// Returns ids for each selected trait, can be resolved against /v2/traits.
	Traits []int `json:"traits"`
}

type CharacterTraining struct {
	// Skill tree id, can be compared against the training section
	// for each /v2/professions.
	ID int `json:"id"`
	// Shows how many hero points have been spent in this tree
	Spent int `json:"spent"`
	// States whether or not the tree is fully trained.
	Done int `json:"done"`
}

type CharacterSAB struct {
	// Contains objects describing which worlds, and in which difficult,
	// have been cleared
	Zones []CharacterSABZone `json:"zones"`
	// Contains objects describing the unlocks on the given character.
	// (list of possible values visible on
	// https://github.com/arenanet/api-cdi/blob/master/v2/characters/sab.js
	Unlocks []CharacterSABUnlock `json:"unlocks"`
	// Contains the objects of unlocked songs on the character
	Songs []CharacterSABSong `json:"songs"`
}

type CharacterSABZone struct {
	// An internal ID that uniquely identifies a zone by world, zone, and mode.
	ID int `json:"id"`
	// The mode used when completing this zone. One of:
	//   infantile, normal, tribulation.
	Mode string `json:"mode"`
	// The world number
	World int `json:"world"`
	// The zone number
	Zone int `json:"zone"`
}

type CharacterSABUnlock struct {
	// The id of the unlock
	ID int `json:"id"`
	// The name of the upgrade
	// Valid values include:
	//   chain_stick, slingshot, whip, mini_bomb, candle, torch, wooden_whistle,
	//   digger, nice_scoop, glove_of_wisdom, bauble_purse, bauble_tote_bag,
	//   moto_breath, moto_finger, health_vessel_1, health_vessel_2,
	//   medium_health_potion
	Name string `json:"name"`
}

type CharacterSABSong struct {
	// The id of the song
	ID int `json:"id"`
	// An unlocalized string describing the song. Valid values include:
	//   secret_song, gatekeeper_lullaby, shatter_serenade
	Name string `json:"name"`
}

type CharacterExtra struct {
	// Contains information on each trained wvw ability
	WvwAbilities CharacterExtraWvwAbilities `json:"wvw_abilities"`
	// Contains information on character's pvp equipment setup.
	EquipmentPvp CharacterExtraEquipmentPvp `json:"equipment_pvp"`
	// Returns character flags. Possible values:
	//   Beta - Beta character for testing period of add-ons
	Flags []string `json:"flags"`
}

type CharacterExtraEquipmentPvp struct {
	// Id for the equipped pvp amulet, can be resolved against /v2/pvp/amulets.
	Amulet int `json:"amulet"`
	// Id for the equipped pvp rune, can be resolved against /v2/items.
	Rune int `json:"rune"`
	// Returns the id for all equipped pvp sigils.
	// Can be resolved against /v2/items.
	Sigils []int `json:"sigils"`
}

type CharacterExtraWvwAbilities struct {
	// Ability id, can be resolved against /v2/wvw/abilities
	ID int `json:"id"`
	// Current rank for the given ability.
	Rank int `json:"rank"`
}

type CharacterSummary struct {
	Character
	// This resource returns information about the backstory of a character
	// attached to a specific account.
	Backstory []string `json:"backstory"`
	// An array containing an entry for each crafting discipline the character
	// has unlocked
	Crafting []CharacterCrafting `json:"crafting"`
	// An array containing an entry for each piece of equipment currently on
	// the selected character
	Equipment []CharacterEquipment `json:"equipment"`
	// Contains one object structure per bag in the character's inventory
	Bags []CharacterInventoryBag `json:"bags"`
	// Contains the pve, pvp, and wvw objects for the current utilities equipped
	Skills CharacterSkills `json:"skills"`
	// contains the pve, pvp, and wvw objects for the current specializations
	// and traits equipped
	Specializations CharacterSpecializations `json:"specializations"`
	// contains objects for each skill tree trained
	Training CharacterTraining `json:"training"`
	CharacterExtra
}

type CharacterBackstorySummary struct {
	// This resource returns information about the backstory of a character
	// attached to a specific account.
	Backstory []string `json:"backstory"`
}

type CharacterCraftingSummary struct {
	// This resource returns information about the backstory of a character
	// attached to a specific account.
	Crafting []CharacterCrafting `json:"crafting"`
}

type CharacterEquimentSummary struct {
	// An array containing an entry for each piece of equipment currently
	// on the selected character
	Equiment []CharacterEquipment `json:"equiment"`
}

type CharacterInventorySummary struct {
	// Contains one object structure per bag in the character's inventory
	Bags []CharacterInventoryBag `json:"bags"`
}

type CharacterSkillSummary struct {
	// Contains the pve, pvp, and wvw objects for the current utilities equipped
	Skills CharacterSkills `json:"skills"`
}

type CharacterReceipesSummary struct {
	// the ID of a recipe that can be resolved against /v2/recipes.
	// The returned recipes describe what the particular character can craft,
	// with respect to their individual crafting professions and levels.
	Recipes []int `json:"recipes"`
}

type CharacterSpecializationsSummary struct {
	// Contains the pve, pvp, and wvw objects for the current specializations
	// and traits equipped
	Specializations CharacterSpecializations `json:"specializations"`
}

type CharacterTrainingSummary struct {
	// Contains objects for each skill tree trained
	Training CharacterTraining `json:"training"`
}

type CharacterEquipmentTab struct {
	// The "id" of this tab. (The position at which it resides.)
	Tab int `json:"tab"`
	// The name given to the equipment combination.
	Name string `json:"name"`
	// Whether or not this is the tab selected on the character currently.
	IsActive bool `json:"is_active"`
	// Contains an object for each equiped piece of equipment
	Equipment CharacterEquipment `json:"equipment"`
	// Contains the following key-value pairs
	EquipmentPvp CharacterExtraEquipmentPvp `json:"equipment_pvp"`
}

// This resource returns information about characters attached to a
// specific account.
// It will return an array of characters name.
func (r *Requestor) CharactersName(pointer *[]string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter).
		collectionIDs("/characters", &pointer)
	return r
}

// This resource returns information about characters attached to a
// specific account.
// It will return a characters summary specific by name.
func (r *Requestor) Character(pointer *CharacterSummary, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter).
		singleton("/characters", &pointer, name)
	return r
}

// This resource returns information about characters attached to a
// specific account.
// It will return an array of characters summary.
func (r *Requestor) Characters(pointer *[]*CharacterSummary, ids ...string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter).
		collection("/characters", &pointer, ids)
	return r
}

// An object containing an array of strings representing backstory answer IDs
// pertaining to the questions answered during character creation.
// References /v2/backstory/answers.
func (r *Requestor) CharacterBackstory(pointer *CharacterBackstorySummary, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter).
		request(fmt.Sprintf("/characters/%s/backstory", name), nil, &pointer)
	return r
}

// This resource returns core information about a character attached
// to a specific account.
func (r *Requestor) CharacterCore(pointer *Character, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter).
		request(fmt.Sprintf("/characters/%s/core", name), nil, &pointer)
	return r
}

// This resource returns core information about a character attached
// to a specific account.
func (r *Requestor) CharacterCrafting(pointer *CharacterCraftingSummary, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter).
		request(fmt.Sprintf("/characters/%s/crafting", name), nil, &pointer)
	return r
}

// This resource returns information about the equipment on a
// character attached to a specific account.
func (r *Requestor) CharacterEquipment(pointer *CharacterEquimentSummary, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds, TokenPermissionInventory).
		request(fmt.Sprintf("/characters/%s/equipment", name), nil, &pointer)
	return r
}

// This resource returns information about the hero points obtained by
// a character attached to a specific account.
func (r *Requestor) CharacterHeroPoints(pointer *[]string, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionProgression).
		request(fmt.Sprintf("/characters/%s/heropoints", name), nil, &pointer)
	return r
}

// This resource returns information about the hero points obtained by
// a character attached to a specific account.
func (r *Requestor) CharacterInventory(pointer *CharacterInventorySummary, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionInventory).
		request(fmt.Sprintf("/characters/%s/inventory", name), nil, &pointer)
	return r
}

// This resource returns information about the quests selected that by a
// character attached to a specific account.
func (r *Requestor) CharacterQuests(pointer *[]int, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionProgression).
		request(fmt.Sprintf("/characters/%s/quests", name), nil, &pointer)
	return r
}

// This resource returns information about recipes that the given
// character can use.
func (r *Requestor) CharacterRecipes(pointer *CharacterReceipesSummary, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionInventory).
		request(fmt.Sprintf("/characters/%s/recipes", name), nil, &pointer)
	return r
}

// This resource returns information about Super Adventure Box on a
// character attached to a specific account.
func (r *Requestor) CharacterSAB(pointer *CharacterSAB, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter).
		request(fmt.Sprintf("/characters/%s/sab", name), nil, &pointer)
	return r
}

// This resource returns information about the skills equipped on
// a character attached to a specific account.
func (r *Requestor) CharacterSkills(pointer *CharacterSkillSummary, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds).
		request(fmt.Sprintf("/characters/%s/skills", name), nil, &pointer)
	return r
}

// This resource returns information about the specializations equipped on a
// character attached to a specific account.
func (r *Requestor) CharacterSpecializations(pointer *CharacterSpecializationsSummary, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds).
		request(fmt.Sprintf("/characters/%s/specializations", name), nil, &pointer)
	return r
}

// This resource returns information about the training of a character
// attached to a specific account.
func (r *Requestor) CharacterTraining(pointer *CharacterTrainingSummary, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds).
		request(fmt.Sprintf("/characters/%s/training", name), nil, &pointer)
	return r
}

// This resource returns information about an accounts build template tabs.
// Request a list of all available tabs.
func (r *Requestor) CharacterBuildTabsIDs(pointer *[]int, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds).
		request(fmt.Sprintf("/characters/%s/buildtabs", name), nil, &pointer)
	return r
}

// This resource returns information about an accounts build template tabs.
// Request information about the specified tab only.
func (r *Requestor) CharacterBuildTab(pointer *CharacterBuildTab, name string, id int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds).
		request(fmt.Sprintf("/characters/%s/buildtabs/%d", name, id), nil, &pointer)
	return r
}

// This resource returns information about an accounts build template tabs.
// Request information about the currently selected tab only.
func (r *Requestor) CharacterActiveBuildTab(pointer *CharacterBuildTab, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds).
		request(fmt.Sprintf("/characters/%s/buildtabs/active", name), nil, &pointer)
	return r
}

// This resource returns information about an accounts equipment template tabs.
// Request a list of all available tabs.
func (r *Requestor) CharacterEquipmentTabsIDs(pointer *[]int, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds).
		request(fmt.Sprintf("/characters/%s/equipmenttabs", name), nil, &pointer)
	return r
}

// This resource returns information about an accounts equipment template tabs.
// Request information about the specified tab only.
func (r *Requestor) CharacterEquipmentTab(pointer *CharacterEquipmentTab, name string, id int) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds).
		request(fmt.Sprintf("/characters/%s/equipmenttabs/%d", name, id), nil, &pointer)
	return r
}

// This resource returns information about an accounts equipment template tabs.
// Request information about the currently selected tab only.
func (r *Requestor) CharacterActiveEquipmentTab(pointer *CharacterEquipmentTab, name string) *Requestor {
	r.
		needPerms(TokenPermissionAccount, TokenPermissionCharacter, TokenPermissionBuilds).
		request(fmt.Sprintf("/characters/%s/equipmenttabs/active", name), nil, &pointer)
	return r
}
