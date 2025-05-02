package domain

import "time"

type Entity struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
