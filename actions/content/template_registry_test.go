package content

func (as *ActionSuite) TestGetTemplateNames() {
	r := as.Require()
	for name, expectedTN := range templateRegistry {
		actualTN, err := getTemplateNames(name)
		r.NoError(err)
		r.Equal(*expectedTN, *actualTN)
		actualTN, err = getTemplateNames(name + "s")
		r.NoError(err)
		r.Equal(*expectedTN, *actualTN)
	}
}
