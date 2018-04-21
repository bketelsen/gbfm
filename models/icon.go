package models

import (
	"github.com/qor/media"
	"github.com/qor/media/oss"
)

type IconStorage struct{ oss.OSS }

func (IconStorage) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 70, Height: 70},
		"middle": {Width: 120, Height: 120},
		"big":    {Width: 320, Height: 320},
	}
}
