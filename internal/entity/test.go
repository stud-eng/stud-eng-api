package entity

import "time"

type Test struct {
	ID       uint32
	Mail     string
	Name     string
	Password string
	UpdateAt time.Time
	CreateAt time.Time
	DeleteAt time.Time
}
