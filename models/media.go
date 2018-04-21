package models

import (
	"github.com/qor/media"
	"github.com/qor/media/oss"
)

type ContentImageStorage struct{ oss.OSS }

func (ContentImageStorage) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small":  {Width: 200, Height: 105},
		"middle": {Width: 600, Height: 315},
		"big":    {Width: 1200, Height: 630},
	}
}
