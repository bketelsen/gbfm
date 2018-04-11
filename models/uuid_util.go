package models

import (
	buuid "github.com/gobuffalo/uuid"
	suuid "github.com/satori/go.uuid"
)

// dirty hack to switch between satori and buffalo UUIDs
func buffaloUUIDToSatori(u buuid.UUID) suuid.UUID {
	return suuid.Must(suuid.FromString(u.String()))
}
