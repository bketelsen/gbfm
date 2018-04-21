package snacks

import "testing"

type slugs struct {
	id       uint
	slug     string
	combined string
}

func TestSlug(t *testing.T) {
	ss := []slugs{
		{
			id:       1,
			slug:     "test",
			combined: "1-test",
		},
		{
			id:       5150,
			slug:     "combined-with-5151-a-number",
			combined: "5150-combined-with-5151-a-number",
		},
	}
	for _, s := range ss {

		id, err := idFromSlug(s.combined)
		if err != nil {
			t.Error("Got Error", err)
		}
		if id != s.id {
			t.Errorf("got %v, expected %v", id, s.id)
		}
	}
}
