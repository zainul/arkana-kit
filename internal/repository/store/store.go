package store

import (
	"github.com/jinzhu/gorm"
)

// Store is storage
type Store struct {
	DB *gorm.DB
}

// NewStore ...
func NewStore(conn *gorm.DB) *Store {
	return &Store{
		DB: conn,
	}
}
