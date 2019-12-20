package utils

import (
	"strings"
	"text/template"
)

func ReworkID(id string) string {
	id = template.URLQueryEscaper(id)
	id = strings.ReplaceAll(id, "+", "%20")
	return id
}