package igen

import (
	"code-generator/generator"
)

type Generator interface {
	Generate() string
	LoadConfig() error
	CurrentTask() *generator.Task
}
