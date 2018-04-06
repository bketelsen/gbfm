package grifts

import (
	"strings"

	"github.com/gophersnacks/gbfm/models"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		// Add DB seeding stuff here
		brian := models.User{}
		brian.Email = "brian@gophersnacks.com"
		brian.Password = "gopher"
		brian.PasswordConfirmation = "gopher"
		brian.Admin = true
		if err, _ := brian.Create(models.DB); err != nil {
			if strings.TrimSpace(err.Error()) != "" {
				return errors.WithStack(err)
			}
		}

		ashley := models.User{}
		ashley.Email = "ashley@gophersnacks.com"
		ashley.Password = "gopher"
		ashley.PasswordConfirmation = "gopher"
		ashley.Admin = true
		if err, _ := ashley.Create(models.DB); err != nil {
			if strings.TrimSpace(err.Error()) != "" {
				return errors.WithStack(err)
			}
		}
		aaron := models.User{}
		aaron.Email = "aaron@gophersnacks.com"
		aaron.Password = "gopher"
		aaron.PasswordConfirmation = "gopher"
		aaron.Admin = true
		if err, _ := aaron.Create(models.DB); err != nil {
			if strings.TrimSpace(err.Error()) != "" {
				return errors.WithStack(err)
			}
		}
		abrian := models.Author{}
		abrian.Name = "Brian Ketelsen"
		abrian.Description = "Brian is a guy from Tampa"
		abrian.PhotoUrl = "/assets/images/authors/brian.jpg"
		abrian.Slug = "brian"
		if err := models.DB.Create(&abrian); err != nil {
			return errors.WithStack(err)
		}

		aashley := models.Author{}
		aashley.Name = "Ashley McNamara"
		aashley.Description = "Ashley is a person from Austin"
		aashley.PhotoUrl = "/assets/images/authors/brian.jpg"
		aashley.Slug = "ashley"
		if err := models.DB.Create(&aashley); err != nil {
			return errors.WithStack(err)
		}

		aaaron := models.Author{}
		aaaron.Name = "Aaron Schlesinger"
		aaaron.Description = "Aaron is a guy from Tampa"
		aaaron.PhotoUrl = "/assets/images/authors/brian.jpg"
		aaaron.Slug = "aaron"
		if err := models.DB.Create(&aaaron); err != nil {
			return errors.WithStack(err)
		}
		snack := models.Snack{}
		snack.AuthorID = brian.ID
		snack.Description = "Create an isomorphic web app with buffalo and gopherjs"
		snack.Title = "Isomorphic Buffalo"
		snack.EmbedCode = `<script src="https://fast.wistia.com/embed/medias/m3g28tikkf.jsonp" async></script><script src="https://fast.wistia.com/assets/external/E-v1.js" async></script><div class="wistia_responsive_padding" style="padding:62.5% 0 0 0;position:relative;"><div class="wistia_responsive_wrapper" style="height:100%;left:0;position:absolute;top:0;width:100%;"><div class="wistia_embed wistia_async_m3g28tikkf videoFoam=true" style="height:100%;width:100%">&nbsp;</div></div></div>`
		snack.ThumbnailURL = "https://embed-ssl.wistia.com/deliveries/3f13e7cad1cc547d181b567602a5c1dc2681709c.jpg?image_crop_resized=200x120"
		snack.Slug = "iso-buffalo"

		if err := models.DB.Create(&snack); err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
	grift.Add("episode", func(c *grift.Context) error {

		episode := models.Episode{}
		episode.Description = "Create an isomorphic web app with buffalo and gopherjs"
		episode.Title = "Isomorphic Buffalo"
		episode.EmbedCode = `<script src="https://fast.wistia.com/embed/medias/m3g28tikkf.jsonp" async></script><script src="https://fast.wistia.com/assets/external/E-v1.js" async></script><div class="wistia_responsive_padding" style="padding:62.5% 0 0 0;position:relative;"><div class="wistia_responsive_wrapper" style="height:100%;left:0;position:absolute;top:0;width:100%;"><div class="wistia_embed wistia_async_m3g28tikkf videoFoam=true" style="height:100%;width:100%">&nbsp;</div></div></div>`
		episode.ThumbnailURL = "https://embed-ssl.wistia.com/deliveries/3f13e7cad1cc547d181b567602a5c1dc2681709c.jpg?image_crop_resized=200x120"
		episode.Slug = "iso-buffalo"
		if err := models.DB.Create(&episode); err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
})
