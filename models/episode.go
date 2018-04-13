package models

// Episode represents an episode
type Episode struct {
	coreModel
	slugger

	Title       string   `json:"title" db:"title"`
	Description string   `json:"description" db:"description"`
	Thumbnail   string   `json:"thumbnail_url" db:"thumbnail_url"`
	EmbedCode   string   `json:"embed_code" db:"embed_code"`
	Body        string   `json:"body" db:"body"`
	Pro         bool     `json:"pro" db:"pro"`
	Repo        string   `json:"repo" db:"repo"`
	Topics      []Topic  `json:"topics" db:"topics" many_to_many:"episodes_topics"`
	Authors     []Author `json:"authors" db:"authors" many_to_many:"episodes_authors"`
}

func init() {
	registry["episode"] = func() (interface{}, interface{}) {
		return new(Episode), new([]Episode)
	}
}
