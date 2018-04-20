package models_test

import (
	"fmt"

	"github.com/gophersnacks/gbfm/models"
)

func (ms *ModelSuite) Test_Snack_Associations() {
	r, db := ms.Require(), ms.DB

	topic := &models.Topic{
		Name: "My Topic",
		Slug: "my-topic",
	}

	r.NoError(db.Create(topic))

	retTop := &models.Topic{}
	r.NoError(db.Eager().Find(retTop, topic.ID))
	r.Equal(topic.ID.String(), retTop.ID.String())

	snack := &models.Snack{}
	snack.Title = "My snack"

	r.NoError(db.Create(snack))

	snack.Topics = append(snack.Topics, *retTop)

	r.NoError(db.Eager().Save(snack))
	retSnack := &models.Snack{}
	r.NoError(db.Eager().Find(retSnack, snack.ID))
	fmt.Println(retSnack, retSnack.Topics)

}
