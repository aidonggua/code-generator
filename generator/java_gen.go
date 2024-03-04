package generator

import (
	"fmt"
	"github.com/spf13/viper"
)

type JavaGenerator struct {
	PackageName string
	UseLombok   bool
}

func (g *JavaGenerator) Generate() string {
	fmt.Printf("%v\n", g)
	return ""
}

func (g *JavaGenerator) LoadConfig() error {
	g.PackageName = viper.GetString("project.package-name")
	g.UseLombok = viper.GetBool("project.use-lombok")
	return nil
}
