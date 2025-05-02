package middleware

import "strings"

type uri string

func (u uri) IsSame(u2 uri) bool {
	r1s := strings.Split(string(u), "/")
	r2s := strings.Split(string(u2), "/")

	if len(r1s) != len(r2s) {
		return false
	}

	for i := range r1s {
		if strings.HasPrefix(r1s[i], ":") || strings.HasPrefix(r2s[i], ":") {
			continue
		}

		if r1s[i] != r2s[i] {
			return false
		}
	}
	return true
}
