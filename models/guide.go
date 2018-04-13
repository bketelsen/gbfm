package models

// Guide is a guide
type Guide struct {
	coreModel
	slugger
	Title        string   `json:"title" db:"title"`
	Description  string   `json:"description" db:"description"`
	ThumbnailURL string   `json:"thumbnail_url" db:"thumbnail_url"`
	EmbedCode    string   `json:"embed_code" db:"embed_code"`
	Body         string   `json:"body" db:"body"`
	Pro          bool     `json:"pro" db:"pro"`
	Topics       []Topic  `json:"topics" db:"topics" many_to_many:"guides_topics"`
	Authors      []Author `json:"authors" db:"authors" many_to_many:"guides_authors"`
}

func init() {
	registry["guide"] = func() (interface{}, interface{}) {
		return new(Guide), new([]Guide)
	}
}
