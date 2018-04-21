package models

import (
	"github.com/qor/media"
	"github.com/qor/media/oss"
)

type ContentImageStorage struct{ oss.OSS }

func (ContentImageStorage) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 50, Height: 50},
		"middle": {Width: 120, Height: 120},
		"big":    {Width: 320, Height: 320},
	}
}
