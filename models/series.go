package models

type Series struct {
	coreModel
	slugger
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Thumbnail   string   `json:"thumbnail"`
	Body        string   `json:"body"`
	Pro         bool     `json:"pro"`
	Keywords    []string `json:"keywords"`
	Episodes    []string `json:"episodes"`
}

func init() {
	registry["series"] = func() (interface{}, interface{}) {
		return new(Series), new([]Series)
	}
}
