//go:generate easytags $GOFILE
package gw2api

type Story struct {
	// The id of the story.
	ID int `json:"id"`
	// The id for the story season; resolvable against v2/stories/seasons.
	Season int `json:"season"`
	// The name of the story.
	Name string `json:"name"`
	// The description of the story.
	Description string `json:"description"`
	// The (in-game, not real-world) date of the story.
	Timeline string `json:"timeline"`
	// The minimum level required for a character to begin this story.
	Level int `json:"level"`
	// The order in which this story is displayed in the Story Journal.
	Order int `json:"order"`
	// An array of chapter objects providing details about the chapters for
	// this story
	Chapters []StoryChapter `json:"chapters"`
	// When present, provides a list of races that are eligible to participate
	// in this story.
	Races []string `json:"races"`
	// When present, provides a list of additional requirements for a character
	// to participate in this story.
	Flags []string `json:"flags"`
}

type StoryChapter struct {
	// The name of the chapter.
	Name string `json:"name"`
}

type StorySeason struct {
	// The id of the season.
	ID string `json:"id"`
	// The name of the season.
	Name string `json:"name"`
	// The order in which this season is displayed in the Story Journal.
	Order int `json:"order"`
	// An array of story ids for the stories that belong to this season.
	Stories []int `json:"stories"`
}

// This resource returns information about the stories that are in the game.
// Return an array of ids for each stories.
func (r *Requestor) StoryIDs(pointer *[]int) *Requestor {
	r.collectionIDs("/stories", &pointer)
	return r
}

// This resource returns information about the stories that are in the game.
// Return a list of response objects
func (r *Requestor) Stories(pointer *[]*Story, ids ...int) *Requestor {
	r.collection("/stories", &pointer, ids)
	return r
}

// This resource returns information about the stories that are in the game.
// Return an object
func (r *Requestor) Story(pointer *Story, id int) *Requestor {
	r.singleton("/stories", &pointer, id)
	return r
}

// This resource returns information about the stories that are in the game.
// Return an array of ids for each story season.
func (r *Requestor) StorySeasonIDs(pointer *[]string) *Requestor {
	r.collectionIDs("/stories/seasons", &pointer)
	return r
}

// This resource returns information about the stories that are in the game.
// Return a list of response objects
func (r *Requestor) StorySeasons(pointer *[]*StorySeason, ids ...string) *Requestor {
	r.collection("/stories/seasons", &pointer, ids)
	return r
}

// This resource returns information about the stories that are in the game.
// Return an object
func (r *Requestor) StorySeason(pointer *StorySeason, id string) *Requestor {
	r.singleton("/stories/seasons", &pointer, id)
	return r
}
