package models

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/gophersnacks/content/content"
)

type FullEpisode struct {
	Episode    *content.Episode
	AuthorList []*content.Author
}

func GetFullEpisodeBySlug(id string) (*FullEpisode, error) {
	var fe FullEpisode
	e, err := GetEpisodeBySlug(id)
	if err != nil {
		return &fe, err
	}
	fe = FullEpisode{
		Episode:    &e,
		AuthorList: make([]*content.Author, len(e.Authors)),
	}
	for i, id := range AuthorIDsForEpisode(e) {
		a, err := GetAuthor(id)
		if err != nil {
			return &fe, err
		}
		fe.AuthorList[i] = &a
	}
	fmt.Println(fe.AuthorList)
	return &fe, nil
}

func AuthorIDsForEpisode(m content.Episode) []int {
	var authors []int
	for _, s := range m.Authors {
		i, err := getID(s)
		if err == nil {
			authors = append(authors, i)
		}
	}
	return authors
}

func getID(s string) (int, error) {
	//?type=Module&id=4
	u, err := url.Parse(s)
	if err != nil {
		return 0, err
	}
	vals := u.Query()
	ii, ok := vals["id"]
	if !ok {
		return 0, err
	}
	return strconv.Atoi(ii[0])
}
