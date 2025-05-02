package auth

import "time"

type Token struct {
	Raw       string
	ExpiresAt time.Time
}

type PairToken struct { // TODO: for future
	Access  Token
	Refresh Token
}
