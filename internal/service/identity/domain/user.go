package domain

// User represents a user in the system.
type User struct {
	Model
	Username string
	Password []byte
}

// Users represents a collection of users.
type Users interface{}
