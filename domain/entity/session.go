package entity

import "time"

type Session struct {
	ID        int
	UserID    int
	Token     string
	ExpiresAt time.Time
}
