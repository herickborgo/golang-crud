package user

import "rest-api/domains/shared"

// User model
type User struct {
	shared.Model
	Name     string `json:"name:omitempty"`
	Email    string `json:"email:omitempty"`
	Password string `json:"password:omitempty"`
}
