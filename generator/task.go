package generator

type Task struct {
	Name        string
	Template    string
	Output      string
	Enable      bool
	Variables   map[string]interface{}
	FilePostfix string
}
