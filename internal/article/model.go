package article

import "time"

// TODO: прикрутить комментарии

type Article struct {
	Id                int       `json:"id"`
	Title             string    `json:"title"`
	Authors           []string  `json:"authors"`
	Flow              string    `json:"flow"`
	CreationTime      time.Time `json:"creationTime"`
	LevelOfComplexity string    `json:"level_of_complexity"`
	Tags              []string  `json:"tags"`
	Hubs              []string  `json:"hubs"`
	Rating            int       `json:"rating"`
	Content           string    `json:"content"`
}
