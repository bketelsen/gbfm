package grifts

import (
	"fmt"
	"strings"

	"github.com/gophersnacks/gbfm/models"
	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
)

var _ = grift.Namespace("db", func() {
	grift.Add("listusers", func(c *grift.Context) error {
		u := []models.User{}
		err := models.GORM.Find(&u)
		if err != nil {
			fmt.Println(err)
		}
		for _, user := range u {
			fmt.Println(user)
		}
		return nil
	})

	grift.Add("listepisodes", func(c *grift.Context) error {
		ee := []models.Episode{}
		err := models.GORM.Find(&ee)
		if err != nil {
			fmt.Println(err)
		}
		for _, e := range ee {
			fmt.Println(e)
		}
		return nil
	})
	grift.Add("episode", func(c *grift.Context) error {
		e := models.Episode{}
		e.Body = "Episode Body Text"
		e.Pro = false
		e.Repo = "https://github.com/gophersnacks/episode15"
		e.Slug = "secret-episode"
		e.Title = "Secret Episode"
		err := e.Create(models.GORM)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	})
	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		// Add DB seeding stuff here
		// Add DB seeding stuff here
		brian := models.User{}
		brian.Email = "brian@gophersnacks.com"
		brian.PasswordHash = "gopher"
		brian.PasswordConfirmation = "gopher"
		brian.Admin = true
		if err := brian.Create(models.GORM); err != nil {
			if strings.TrimSpace(err.Error()) != "" {
				return errors.WithStack(err)
			}
		}

		ashley := models.User{}
		ashley.Email = "ashley@gophersnacks.com"
		ashley.PasswordHash = "gopher"
		ashley.PasswordConfirmation = "gopher"
		ashley.Admin = true
		if err := ashley.Create(models.GORM); err != nil {
			if strings.TrimSpace(err.Error()) != "" {
				return errors.WithStack(err)
			}
		}
		aaron := models.User{}
		aaron.Email = "aaron@gophersnacks.com"
		aaron.PasswordHash = "gopher"
		aaron.PasswordConfirmation = "gopher"
		aaron.Admin = true
		if err := aaron.Create(models.GORM); err != nil {
			if strings.TrimSpace(err.Error()) != "" {
				return errors.WithStack(err)
			}
		}
		abrian := models.Author{}
		abrian.Name = "Brian Ketelsen"
		abrian.Description = "Brian is a guy from Tampa"
		abrian.Photo = "/assets/images/authors/brian.jpg"
		abrian.Slug = "brian"
		if err := abrian.Create(models.GORM); err != nil {
			return errors.WithStack(err)
		}

		aashley := models.Author{}
		aashley.Name = "Ashley McNamara"
		aashley.Description = "Ashley is a person from Austin"
		aashley.Photo = "/assets/images/authors/brian.jpg"
		aashley.Slug = "ashley"
		if err := aashley.Create(models.GORM); err != nil {
			return errors.WithStack(err)
		}

		aaaron := models.Author{}
		aaaron.Name = "Aaron Schlesinger"
		aaaron.Description = "Aaron is a guy from Tampa"
		aaaron.Photo = "/assets/images/authors/brian.jpg"
		aaaron.Slug = "aaron"
		if err := aaaron.Create(models.GORM); err != nil {
			return errors.WithStack(err)
		}

		return nil
	})
})
