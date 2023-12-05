package model

import (
	"time"
)

type Note struct {
	Id       int64
	Title    string
	Text     string
	Author   string
	CreateAt time.Time
	UpdateAt *time.Time
}
