# gw2api-go

A Go Wrapper to your GuildWars 2 Go Project

### What's is GW2-API 

At the start GW2-API as created to be a wrapper of the GuildWars2 Official API.
After few months of use in private projects, I decide to make it public and rewrite the library.

GW2-API is a complete and up-to-date wrapper in Golang.

### Getting Start / How to use ?

To use Gw2APIGO, it's really simple. All requests and API context is managed by a `gw2api.Requestor`

```go
  // Create a new gw2api Requestor
  r := gw2api.NewRequestor()

  // Initialize api object 
  var world gw2api.World

  // Perform a request with our defined api object
  r.World(&world, 2101)

  // Show magic
  log.Printf("Awesome world: %s", world.Name)
```


The best pratice is ti check if the request pipeline encountered an error 
```go
  // Create a new gw2api Requestor
  if err := r.World(&world, 2101).Err(); err != nil {
    panic(err.Error())
  }
```


In some endpoints you can need a translation, you can do a request with a specific
lang with `.Lang(gw2api.Lang)`
```go
  r.Lang(gw2api.LangFR).Title(&title, 1)
```


On endpoints who needs API Key and permission, you can give to requestor an APIKey,
the requestor will fetch permission of APIKey to prevent any call to the API if your API
Key don't have needed permission
```go
  r.Auth(apiKey).Account(&account)
```

In some advanced case, you can edit the timeout of the requestor too with `.Timeout(time.Duration)`
```go
  r.Timeout(5 * time.Second).Title(&title, 1)
```


When you try to call an authenticated endpoints without APIKey, the Requestor will return you
an error accessible on `.Err()` method. You can check the Err wirh `errors` package.
Its the same process when your API Key dont have the required scope
```go
  err := r.Account(&account).Err()
  if errors.Is(err, gw2api.ErrRequireAuthentication) { // => true
    log.Printf("Give an api key")
  }
```

| Err  | Description     |
|-----|------------------|
| ErrTooManyRequest | too many request: 429. You have reach API limitations |
| ErrRequireAuthentication | API endpoint needs authentication |
| ErrMissingScope | missing scope permissions for this endpoint |



### TODO

- [x] Core wrapper
- [x] Language support
- [x] API Token Permissions
- [ ] Route active support
- [x] Subtoken support
- [ ] v2/
  - [x] account
    - [x] account/achievements
    - [x] account/bank
    - [x] account/buildstorage
    - [x] account/dailycrafting
    - [x] account/dungeons
    - [x] account/dyes
    - [x] account/emotes
    - [x] account/finishers
    - [x] account/gliders
    - [x] account/home
      - [x] account/home/cats
      - [x] account/home/nodes
    - [x] account/inventory
    - [x] account/legendaryarmory
    - [x] account/luck
    - [-] account/mail (API not active)
    - [x] account/mailcarriers
    - [x] account/mapchests
    - [x] account/masteries
    - [x] account/mastery/points
    - [x] account/materials
    - [x] account/minis
    - [x] account/mounts
      - [x] account/mounts/skins
      - [x] account/mounts/types
    - [x] account/novelties
    - [x] account/outfits
    - [x] account/pvp/heroes
    - [x] account/raids
    - [x] account/recipes
    - [x] account/skins
    - [x] account/titles
    - [x] account/wallet
    - [x] account/worldbosses
  - [x] achievements
    - [x] achievements/categories
    - [x] achievements/daily
    - [x] achievements/daily/tomorrow
    - [x] achievements/groups
  - [-] adventures (API not active)
    - [-] adventures/:id/leaderboards (API not active)
      - [-] adventures/:id/leaderboards/:board/:region (API not active)
  - [x] backstory/answers
  - [x] backstory/questions
  - [x] build
  - [x] characters
    - [x] (characters all and single)
    - [x] characters/:id/backstory
    - [x] characters/:id/buildtabs
    - [x] characters/:id/buildtabs/active
    - [x] characters/:id/core
    - [x] characters/:id/crafting
    - [-] characters/:id/dungeons (API not active)
    - [x] characters/:id/equipment
    - [x] characters/:id/equipmenttabs
    - [x] characters/:id/equipmenttabs/active
    - [x] characters/:id/heropoints
    - [x] characters/:id/inventory
    - [x] characters/:id/quests
    - [x] characters/:id/recipes
    - [x] characters/:id/sab
    - [x] characters/:id/skills
    - [x] characters/:id/specializations
    - [x] characters/:id/training
  - [x] colors
  - [x] commerce/delivery
  - [x] commerce/exchange
  - [x] commerce/listings
  - [x] commerce/prices
  - [x] commerce/transactions
    - [ ] Pagination System
  - [x] continents
  - [x] createsubtoken
  - [x] currencies
  - [x] dailycrafting
  - [x] dungeons
  - [x] emblem
  - [x] emotes
  - [-] events (API not active)
  - [-] events-state (API not active)
  - [x] files
  - [x] finishers
  - [-] gemstore/catalog (API not active)
  - [x] gliders
  - [x] guild/:id
    - [x] guild/:id/log
    - [x] guild/:id/members
    - [x] guild/:id/ranks
    - [x] guild/:id/stash
    - [x] guild/:id/storage
    - [ ] guild/:id/teams - needs /v2/pvp/stats.
    - [x] guild/:id/treasury
    - [x] guild/:id/upgrades
  - [x] guild/permissions
  - [x] guild/search
  - [x] guild/upgrades
  - [ ] home
    - [ ] home/cats
    - [ ] home/nodes
  - [ ] items
  - [ ] itemstats
  - [ ] legendaryarmory
  - [ ] legends
  - [ ] mailcarriers
  - [ ] mapchests
  - [ ] maps
  - [ ] masteries
  - [ ] materials
  - [ ] minis
  - [ ] mounts
    - [ ] mounts/skins
    - [ ] mounts/types
  - [ ] novelties
  - [ ] outfits
  - [ ] pets
  - [ ] professions
  - [ ] pvp
    - [ ] pvp/amulets
    - [ ] pvp/games
    - [ ] pvp/heroes
    - [ ] pvp/ranks
    - [ ] pvp/rewardtracks
    - [ ] pvp/runes
    - [ ] pvp/seasons
      - [ ] pvp/seasons/:id/leaderboards
        - [ ] pvp/seasons/:id/leaderboards/:board/:region
    - [ ] pvp/sigils
    - [ ] pvp/standings
    - [ ] pvp/stats
  - [ ] quaggans
  - [ ] quests
  - [ ] races
  - [ ] raids
  - [ ] recipes
    - [ ] recipes/search
  - [ ] skills
  - [ ] skins
  - [ ] specializations
  - [x] stories
    - [x] stories/seasons
  - [x] titles
  - [x] tokeninfo
  - [ ] traits
  - [-] vendors (API not active)
  - [x] worldbosses
  - [x] worlds
  - [x] wvw/abilities
  - [x] wvw/matches
  - [x] wvw/matches/overview
  - [x] wvw/matches/scores
  - [x] wvw/matches/stats
  - [ ] wvw/matches/stats/:id/guilds/:guild_id (no doc)
  - [ ] wvw/matches/stats/:id/teams/:team/top/kdr (no doc)
  - [ ] wvw/matches/stats/:id/teams/:team/top/kills (no doc)
  - [x] wvw/objectives
  - [x] wvw/ranks
  - [-] wvw/rewardtracks (API not active)
  - [x] wvw/upgrades