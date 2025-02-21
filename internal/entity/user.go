package entity

type User struct {
	ID           int64  `json:"id,omitempty"`
	Login        string `json:"login"`
	Password     string `json:"password"`
	PasswordHash string `json:"password_hash,omitempty"`
	PasswordSalt string `json:"password_salt,omitempty"`
	Secret       string `json:"secret,omitempty"`
}

//����� �� ��������, �� � ���� ���� ���� �������� (� ���� ���� 1�)
