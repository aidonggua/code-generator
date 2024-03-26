package task

import "golang.org/x/exp/slices"

type Task struct {
	Name       string
	Template   string
	SourceType string
	Table      string
	Output     string
	Options    string
	Enable     bool
	Variables  map[string]interface{}
	Imports    []string
}

func (t *Task) AddImport(s string) {
	if !slices.Contains(t.Imports, s) {
		t.Imports = append(t.Imports, s)
	}
}
