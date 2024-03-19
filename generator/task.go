package generator

type Task struct {
	Name       string
	Template   string
	SourceType string
	Table      string
	Output     string
	Options    string
	Enable     bool
	Variables  map[string]interface{}
}
