package contract

// RegisterUser  ...
type RegisterUser struct {
	FirstName      string `json:"first_name" validate:"string,non_empty=true=01E6"`
	LastName       string `json:"last_name" validate:"string,non_empty=true=01E7"`
	Email          string `json:"email" validate:"email,non_empty=true=01E8,validemail=true=01E9"`
	Password       string `json:"password" validate:"string,non_empty=true=01E10"`
	AccountType    string `json:"account_type" validate:"string,non_empty=true=01E11"`
	PhoneNumber    string `json:"phone" validate:"string,non_empty=true=01E12"`
	InvitationCode string `json:"invitation_code"`
}

// ActivateUser activate user
type ActivateUser struct {
	Code string `json:"code" validate:"string,non_empty=true=01E16"`
}
