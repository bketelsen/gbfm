package models

func (ms *ModelSuite) TestRoundTripEpisode() {
	r, db := ms.Require(), ms.DB

	ep := &Episode{
		Topics: []Topic{
			Topic{Slug: "abc"},
		},
		Authors: []Author{
			Author{Slug: "def"},
		},
	}
	r.NoError(db.Eager().Create(ep))

	retEp := &Episode{}
	r.NoError(db.Eager().Find(retEp, ep.ID))
	r.Equal(ep.ID.String(), retEp.ID.String())
	r.Equal(ep.Authors, retEp.Authors)
	r.Equal(ep.Topics, retEp.Topics)
}
