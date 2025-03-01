package entity

type User struct {
	ID           int64  `json:"id,omitempty" example:"1"`
	Login        string `json:"login" example:"login"`
	Password     string `json:"password" example:"password"`
	PasswordHash string `json:"password_hash,omitempty" example:""` //!!!!!!!!!!!!!!!!!!
	PasswordSalt string `json:"password_salt,omitempty" example:""` //!!!!!!!!!!!!!!!!!!
}
