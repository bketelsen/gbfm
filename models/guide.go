package models

// Guide is a guide
type Guide struct {
	coreModel
	slugger
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Thumbnail   string  `json:"thumbnail"`
	EmbedCode   string  `json:"embed_code"`
	Body        string  `json:"body"`
	Pro         bool    `json:"pro"`
	Topics      []Topic `json:"topics" db:"topics" has_many:"topics"`
	Author      Author  `json:"author" db:"author" has_one:"author"`
}
