package entity

import "time"

type (
	Test struct {
		ID        uint32
		Mail      string
		Name      string
		Password  string
		UpdatedAt time.Time
		CreatedAt time.Time
		DeletedAt time.Time
	}

	// Tests struct {
	// 	Tests []*Test
	// }
)
