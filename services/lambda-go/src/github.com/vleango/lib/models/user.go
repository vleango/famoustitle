package models

import (
	"fmt"
	"time"
)

type User struct {
	ID             string            `json:"id"`
	FirstName      string            `json:"first_name"`
	LastName       string            `json:"last_name"`
	Email          string            `json:"email"`
	PasswordDigest string            `json:"password_digest"`
	Admin          bool              `json:"admin"`
	Articles       map[string]string `json:"articles"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
}

func (u *User) FullName() string {
	return fmt.Sprintf("%v %v", u.FirstName, u.LastName)
}
