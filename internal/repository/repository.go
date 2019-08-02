package repository

import "github.com/zainul/arkana-kit/internal/entity"

// Account ...
type Account interface {
	SaveUser(e entity.User) error
	UserBy(field string, value interface{}) ([]entity.User, error)
	Update(whereField string, whereVal interface{}, valuesToUpdate map[string]interface{}) error
}
