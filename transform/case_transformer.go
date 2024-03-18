package transform

import "strings"

type CaseTransformer struct {
}

func (*CaseTransformer) Title(s string) string {
	return strings.Title(s)
}

func (*CaseTransformer) CamelCase(s string) string {
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

func (*CaseTransformer) SnakeCase(s string) string {
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
