package generator

type Task struct {
	Name       string
	Template   string
	FileType   string
	Prefix     string
	Postfix    string
	Enable     bool
	Properties map[string]interface{}
}
