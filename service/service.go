package service

import (
	"bookmark-man/models"
	"encoding/json"
	"fmt"

	scribble "github.com/risjacksonlee/scribble-with-updates"
)

type Service struct {
	db *scribble.Driver
}

func New() *Service {

	db, err := scribble.New(".", nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	return &Service{
		db: db,
	}
}

func (s *Service) GetBookmarksForUser(usr string) ([]models.Bookmark, error) {
	collection := fmt.Sprintf("%s_bookmarks", usr)
	jbs, err := s.db.ReadAll(collection)
	if err != nil {
		return []models.Bookmark{}, err
	}

	var bms []models.Bookmark
	for _, f := range jbs {
		b := models.Bookmark{}
		if err := json.Unmarshal([]byte(f), &b); err != nil {
			return []models.Bookmark{}, err
		}
		bms = append(bms, b)
	}

	return bms, nil
}

func (s *Service) Get(usr string, id string) (models.Bookmark, error) {
	coll := fmt.Sprintf("%s_bookmarks", usr)

	r := models.Bookmark{}
	if err := s.db.Read(coll, id, &r); err != nil {
		return models.Bookmark{}, err
	}

	return r, nil
}

func (s Service) Add(usr string, b models.Bookmark) error {
	coll := fmt.Sprintf("%s_bookmarks", usr)

	if err := s.db.Write(coll, b.Name, b); err != nil {
		return err
	}

	return nil
}

func (s Service) Update(usr string, b models.Bookmark) (bool, error) {
	coll := fmt.Sprintf("%s_bookmarks", usr)

	if err := s.db.Write(coll, b.Id, b); err != nil {
		return false, err
	}

	return true, nil
}

func (s Service) Delete(usr string, id string) error {
	coll := fmt.Sprintf("%s_bookmarks", usr)

	err := s.db.Delete(coll, id)

	return err
}
