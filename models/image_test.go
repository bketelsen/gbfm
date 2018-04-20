package models

func (ms ModelSuite) Test_Image() {
	r, db := ms.Require(), ms.DB
	_, err := EmptyFromRegistry("image")
	r.NoError(err)
	sample, err := SampleFromRegistry("image")
	r.NoError(err)
	r.NoError(db.Create(sample))
	r.NotEmpty(sample.GetID())
}
