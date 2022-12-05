package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")
var ErrDuplicateMail = errors.New("Duplicate email ")
var ErrInvalidCredential = errors.New("Invalid credentials")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID              int
	Name            string
	Email           string
	Hashed_Password []byte
	Created         time.Time
}
