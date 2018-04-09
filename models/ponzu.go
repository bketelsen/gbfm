/*
* CODE GENERATED AUTOMATICALLY WITH github.com/bketelsen/ponzigen
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package models

import (
	"github.com/bketelsen/ponzi"
	"github.com/gophersnacks/content/content"
	"github.com/pkg/errors"
	"time"
)

var BaseURL string

type AuthorListResult struct {
	Data []content.Author `json:"data"`
}
type EpisodeListResult struct {
	Data []content.Episode `json:"data"`
}
type GifmListResult struct {
	Data []content.Gifm `json:"data"`
}
type GuideListResult struct {
	Data []content.Guide `json:"data"`
}
type SeriesListResult struct {
	Data []content.Series `json:"data"`
}
type SnackListResult struct {
	Data []content.Snack `json:"data"`
}

var authorCache *ponzi.Cache
var episodeCache *ponzi.Cache
var gifmCache *ponzi.Cache
var guideCache *ponzi.Cache
var seriesCache *ponzi.Cache
var snackCache *ponzi.Cache

func initAuthorCache() {
	if authorCache == nil {
		authorCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initEpisodeCache() {
	if episodeCache == nil {
		episodeCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initGifmCache() {
	if gifmCache == nil {
		gifmCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initGuideCache() {
	if guideCache == nil {
		guideCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initSeriesCache() {
	if seriesCache == nil {
		seriesCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}
func initSnackCache() {
	if snackCache == nil {
		snackCache = ponzi.New(BaseURL, 1*time.Minute, 30*time.Second)
	}
}

func GetAuthor(id int) (content.Author, error) {
	initAuthorCache()
	var sp AuthorListResult
	err := authorCache.Get(id, "Author", &sp)
	if err != nil {
		return content.Author{}, err
	}
	if len(sp.Data) == 0 {
		return content.Author{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetEpisode(id int) (content.Episode, error) {
	initEpisodeCache()
	var sp EpisodeListResult
	err := episodeCache.Get(id, "Episode", &sp)
	if err != nil {
		return content.Episode{}, err
	}
	if len(sp.Data) == 0 {
		return content.Episode{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetGifm(id int) (content.Gifm, error) {
	initGifmCache()
	var sp GifmListResult
	err := gifmCache.Get(id, "Gifm", &sp)
	if err != nil {
		return content.Gifm{}, err
	}
	if len(sp.Data) == 0 {
		return content.Gifm{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetGuide(id int) (content.Guide, error) {
	initGuideCache()
	var sp GuideListResult
	err := guideCache.Get(id, "Guide", &sp)
	if err != nil {
		return content.Guide{}, err
	}
	if len(sp.Data) == 0 {
		return content.Guide{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetSeries(id int) (content.Series, error) {
	initSeriesCache()
	var sp SeriesListResult
	err := seriesCache.Get(id, "Series", &sp)
	if err != nil {
		return content.Series{}, err
	}
	if len(sp.Data) == 0 {
		return content.Series{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetSnack(id int) (content.Snack, error) {
	initSnackCache()
	var sp SnackListResult
	err := snackCache.Get(id, "Snack", &sp)
	if err != nil {
		return content.Snack{}, err
	}
	if len(sp.Data) == 0 {
		return content.Snack{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}

func GetAuthorBySlug(slug string) (content.Author, error) {
	initAuthorCache()
	var sp AuthorListResult
	err := authorCache.GetBySlug(slug, "Author", &sp)
	if err != nil {
		return content.Author{}, err
	}
	if len(sp.Data) == 0 {
		return content.Author{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetEpisodeBySlug(slug string) (content.Episode, error) {
	initEpisodeCache()
	var sp EpisodeListResult
	err := episodeCache.GetBySlug(slug, "Episode", &sp)
	if err != nil {
		return content.Episode{}, err
	}
	if len(sp.Data) == 0 {
		return content.Episode{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetGifmBySlug(slug string) (content.Gifm, error) {
	initGifmCache()
	var sp GifmListResult
	err := gifmCache.GetBySlug(slug, "Gifm", &sp)
	if err != nil {
		return content.Gifm{}, err
	}
	if len(sp.Data) == 0 {
		return content.Gifm{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetGuideBySlug(slug string) (content.Guide, error) {
	initGuideCache()
	var sp GuideListResult
	err := guideCache.GetBySlug(slug, "Guide", &sp)
	if err != nil {
		return content.Guide{}, err
	}
	if len(sp.Data) == 0 {
		return content.Guide{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetSeriesBySlug(slug string) (content.Series, error) {
	initSeriesCache()
	var sp SeriesListResult
	err := seriesCache.GetBySlug(slug, "Series", &sp)
	if err != nil {
		return content.Series{}, err
	}
	if len(sp.Data) == 0 {
		return content.Series{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}
func GetSnackBySlug(slug string) (content.Snack, error) {
	initSnackCache()
	var sp SnackListResult
	err := snackCache.GetBySlug(slug, "Snack", &sp)
	if err != nil {
		return content.Snack{}, err
	}
	if len(sp.Data) == 0 {
		return content.Snack{}, errors.New("Not Found")
	}
	return sp.Data[0], err

}

func GetAuthorList() ([]content.Author, error) {
	initAuthorCache()
	var sp AuthorListResult
	err := authorCache.GetAll("Author", &sp)
	if err != nil {
		return []content.Author{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Author{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetEpisodeList() ([]content.Episode, error) {
	initEpisodeCache()
	var sp EpisodeListResult
	err := episodeCache.GetAll("Episode", &sp)
	if err != nil {
		return []content.Episode{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Episode{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetGifmList() ([]content.Gifm, error) {
	initGifmCache()
	var sp GifmListResult
	err := gifmCache.GetAll("Gifm", &sp)
	if err != nil {
		return []content.Gifm{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Gifm{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetGuideList() ([]content.Guide, error) {
	initGuideCache()
	var sp GuideListResult
	err := guideCache.GetAll("Guide", &sp)
	if err != nil {
		return []content.Guide{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Guide{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetSeriesList() ([]content.Series, error) {
	initSeriesCache()
	var sp SeriesListResult
	err := seriesCache.GetAll("Series", &sp)
	if err != nil {
		return []content.Series{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Series{}, errors.New("Not Found")
	}
	return sp.Data, err

}
func GetSnackList() ([]content.Snack, error) {
	initSnackCache()
	var sp SnackListResult
	err := snackCache.GetAll("Snack", &sp)
	if err != nil {
		return []content.Snack{}, err
	}
	if len(sp.Data) == 0 {
		return []content.Snack{}, errors.New("Not Found")
	}
	return sp.Data, err

}
