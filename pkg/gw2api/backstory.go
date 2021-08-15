//go:generate easytags $GOFILE
package gw2api

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type BackstoryAnswer struct {
	// The id of the answer.
	ID string `json:"id"`
	// The title (or name) of the answer.
	Title string `json:"title"`
	// The description of the answer; as displayed in-game when presented as
	// an answer choice to a question during the Biography portion of character
	// creation.
	Description string `json:"description"`
	// The Story Journal entry for the answer; as displayed in-game.
	Journal string `json:"journal"`
	// The id of the Biography question that this answers;
	// resolves against /v2/backstory/questions`.
	Question int `json:"question"`
	// When present, an array of professions that this answer is
	// available as a choice for.
	Professions []string `json:"professions"`
	// When present, an array of races that this answer is available
	// as a choice for.
	Races []string `json:"races"`
}

type BackstoryQuestion struct {
	// The id of the question.
	ID string `json:"id"`
	// The title (or name) of the question.
	Title string `json:"title"`
	// The description of the question; as displayed in-game when
	// presented as a Biography choice during character creation.
	Description string `json:"description"`
	// The list of answers for this question; resolvable
	// against /v2/backstory/answers.
	Answers []string `json:"answers"`
	// The order in which this question is displayed in-game while answering
	// your characters' Biography questions during character creation.
	Order int `json:"order"`
	// When present, an array of professions that this answer is
	// available as a choice for.
	Professions []string `json:"professions"`
	// When present, an array of races that this answer is available
	// as a choice for.
	Races []string `json:"races"`
}

// This resource returns information about the Biography answers that are in the game.
func (r *Requestor) BackstoryAnswer(backstoryAnswer *BackstoryAnswer, id string) *Requestor {
	r.request("/backstory/answers", url.Values{"id": []string{id}}, &backstoryAnswer)
	return r
}

// This resource returns information about the Biography answers that are in the game.
func (r *Requestor) BackstoryAnswers(backstoryAnswers *[]*BackstoryAnswer, ids ...string) *Requestor {
	r.request("/backstory/answers", url.Values{"id": ids}, &backstoryAnswers)
	return r
}

// This resource returns information about the Biography answers that are in the game.
func (r *Requestor) BackstoryAnswersIDs(s []string) *Requestor {
	r.request("/backstory/answers", nil, &s)
	return r
}

// This resource returns information about the Biography questions that are in the game.
func (r *Requestor) BackstoryQuestion(backstoryQuestion *BackstoryQuestion, id int) *Requestor {
	r.request("/backstory/questions", url.Values{"id": []string{fmt.Sprint(id)}}, &backstoryQuestion)
	return r
}

// This resource returns information about the Biography questions that are in the game.
func (r *Requestor) BackstoryQuestions(backstoryQuestions *[]*BackstoryQuestion, ids ...int) *Requestor {
	if len(ids) == 0 {
		r.err = errors.New("at least one id must be given")
		return r
	}

	var urlValues url.Values
	sIds := strings.Trim(strings.Replace(fmt.Sprint(ids), " ", ",", -1), "[]")
	urlValues = url.Values{"ids": strings.Split(sIds, ",")}

	r.request("/backstory/questions", urlValues, &backstoryQuestions)
	return r
}

// This resource returns information about the Biography questions that are in the game.
func (r *Requestor) BackstoryQuestionsIDs(s []int) *Requestor {
	r.request("/backstory/questions", nil, &s)
	return r
}
