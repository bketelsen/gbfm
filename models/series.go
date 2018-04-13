package models

type Series struct {
	coreModel
	slugger
	Title        string   `json:"title" db:"title"`
	Description  string   `json:"description" db:"description"`
	ThumbnailURL string   `json:"thumbnail_url" db:"thumbnail_url`
	Body         string   `json:"body" db:"body"`
	Pro          bool     `json:"pro" db:"pro"`
	Topics       []Topic  `json:"topics" db:"topics" many_to_many:"series_topics"`
	Authors      []Author `json:"authors" db:"authors" many_to_many:"series_authors"`
}

func init() {
	registry["series"] = func() (interface{}, interface{}) {
		return new(Series), new([]Series)
	}
}
