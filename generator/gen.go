package generator

type Generator interface {
	Generate() string
	LoadConfig(project string) error
}
