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

	// verify that authors were saved and queried properly
	r.Equal(len(ep.Authors), len(retEp.Authors))
	for i, expectedAuthor := range ep.Authors {
		r.Equal(expectedAuthor.ID.String(), ep.Authors[i].ID.String())
	}

	// verify that topics were saved and queries properly
	r.Equal(len(ep.Topics), len(retEp.Topics))
	for i, expectedTopic := range ep.Topics {
		r.Equal(expectedTopic.ID.String(), ep.Topics[i].ID.String())
	}
}
