package search

import (
	"fmt"
	"regexp"
	"strings"
)

type Vector string

// Clause returns a SQL clause that can be used to search for a string Vector in the given columns.
// For example, Vector("foo bar").Clause("|", "u.first_name", "u.last_name") returns:
//
//	to_tsvector(u.first_name || u.last_name) @@ to_tsquery('foo:* | bar:*');
//
// which can be used in a WHERE clause.
func (v Vector) Clause(operator string, columns ...string) string {
	words := strings.Split(strings.TrimSpace(regexp.MustCompile(`\s+`).ReplaceAllString(string(v), " ")), " ")
	for i := 0; i < len(words); i++ {
		words[i] = words[i] + ":*"
	}
	for i := 0; i < len(columns); i++ {
		columns[i] = columns[i] + "::text"
	}
	return fmt.Sprintf("to_tsvector(%s) @@ to_tsquery('%s')",
		strings.Join(columns, " ||' '|| "),
		strings.Join(words, fmt.Sprintf(" %s ", operator)),
	)
}
