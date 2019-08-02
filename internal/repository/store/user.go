package store

import (
	"strings"

	"github.com/zainul/arkana-kit/internal/entity"
)

// SaveUser is insert user
func (s *Store) SaveUser(e entity.User) error {
	return s.DB.Create(&e).Error
}

// UserBy get user by dynamic param as column
func (s *Store) UserBy(field string, value interface{}) ([]entity.User, error) {
	result := make([]entity.User, 0)
	err := s.DB.Where(map[string]interface{}{field: value}).Find(&result).Error
	return result, err
}

// Update single condition
func (s *Store) Update(field string, val interface{}, valuesToUpdate map[string]interface{}) error {
	query := "UPDATE \"user\" SET "
	vals := make([]interface{}, 0)
	queryField := make([]string, 0)

	for key, valmap := range valuesToUpdate {
		vals = append(vals, valmap)

		queryField = append(queryField, key+"=?")
	}

	query = query + strings.Join(queryField, " ,") + " WHERE " + field + "= ?"
	vals = append(vals, val)

	return s.DB.Exec(query, vals...).Error
}
