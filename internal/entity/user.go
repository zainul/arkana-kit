package entity

import "time"

type User struct {
	ID             string     `json:"id" gorm:"column:id;"`
	FirstName      string     `json:"first_name" gorm:"column:first_name;"`
	LastName       string     `json:"last_name" gorm:"column:last_name;"`
	Email          string     `json:"email" gorm:"column:email;"`
	EncryptedPass  string     `json:"encrypted_pass" gorm:"column:encrypted_pass;"`
	PhoneNumber    string     `json:"phone_number" gorm:"column:phone_number;"`
	AccountType    int        `json:"account_type" gorm:"column:account_type;"`
	InvitationCode string     `json:"invitation_code" gorm:"column:invitation_code;"`
	ActivationCode string     `json:"activation_code" gorm:"column:activation_code;"`
	ActivateAt     *time.Time `json:"activate_at" gorm:"column:activate_at;"`
	ModifiedAt     *time.Time `json:"modified_at" gorm:"column:modified_at;"`
	Username       string     `json:"username" gorm:"column:username;"`
}

// TableName only for gorm purpose
func (e *User) TableName() string {
	return "user"
}
