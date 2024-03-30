package helper

import "strings"

type CaseHelper struct {
}

func (*CaseHelper) CamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool { return r == ' ' || r == '_' })
	result := ""
	for i, word := range words {
		if i == 0 {
			result += strings.ToLower(word)
		} else {
			result += strings.Title(word)
		}
	}

	return result
}

func (*CaseHelper) SnakeCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool { return r == ' ' || r == '_' })
	result := ""
	for i, word := range words {
		if i == 0 {
			result += strings.ToLower(word)
		} else {
			result += "_" + strings.ToLower(word)
		}
	}

	return result
}

func (c *CaseHelper) KebabCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool { return r == ' ' || r == '_' })
	result := ""
	for i, word := range words {
		if i == 0 {
			result += strings.ToLower(word)
		} else {
			result += "-" + strings.ToLower(word)
		}
	}

	return result
}
