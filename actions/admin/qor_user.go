package admin

// qorCurrentUser implements qor.CurentUser
//
// I didn't see anything in the qor suite of packages that actually implements this.
// I didn't look that deep because this is pretty easy to implement
type qorCurrentUser struct {
	displayName string
}

func (qu qorCurrentUser) DisplayName() string {
	return qu.displayName
}
