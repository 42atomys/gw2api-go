package gw2api

type TokenPermission uint

const (
	_                                      = iota
	TokenPermissionAccount TokenPermission = iota
	TokenPermissionBuilds
	TokenPermissionCharacter
	TokenPermissionGuilds
	TokenPermissionInventory
	TokenPermissionProgression
	TokenPermissionPvP
	TokenPermissionTradingpost
	TokenPermissionUnlocks
	TokenPermissionWallet
	TokenPermissionSize
)

var (
	permissionsMapping = map[string]TokenPermission{
		"account":     TokenPermissionAccount,
		"builds":      TokenPermissionBuilds,
		"characters":  TokenPermissionCharacter,
		"guilds":      TokenPermissionGuilds,
		"inventories": TokenPermissionInventory,
		"progression": TokenPermissionProgression,
		"pvp":         TokenPermissionPvP,
		"tradingpost": TokenPermissionTradingpost,
		"unlocks":     TokenPermissionUnlocks,
		"wallet":      TokenPermissionWallet,
	}
)

func getBitwise(n, pos uint) bool {
	return (n>>pos)&1 == 1
}

func setBitwise(n *uint, pos uint) {
	*n |= 1 << pos
}
