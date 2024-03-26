package igen

import "code-generator/task"

type Generator interface {
	Generate() string
	LoadConfig() error
	CurrentTask() *task.Task
}
