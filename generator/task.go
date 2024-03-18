package generator

type Task struct {
	Name       string
	Template   string
	Type       string
	SubPackage string
	Source     string
	Output     string
	Options    string
	Enable     bool
}
