package gormstore

import (
	"gorm.io/gorm"

	"github.com/Dominux/clean_architecture_blog/internal/store"
)


type Store struct {
	db *gorm.DB
	postRepository *PostRepository
}

func New(db *gorm.DB) *Store {
	return &Store{db: db}
}

func (s *Store) Post() store.PostRepository {
	if s.postRepository == nil {
		s.postRepository = &PostRepository{
			store: s,
		}
	}
	return s.postRepository
}
