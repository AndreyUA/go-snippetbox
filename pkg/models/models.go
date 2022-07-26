package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("snippet not found!")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
