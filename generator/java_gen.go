package generator

import (
	"fmt"
	"github.com/spf13/viper"
)

type JavaGenerator struct {
}

func (g *JavaGenerator) Generate() string {
	fmt.Println(viper.GetString("java.package-name"))
	fmt.Println(viper.GetString("java.use-lombok"))
	return ""
}

func (g *JavaGenerator) LoadConfig(project string) error {
	viper.SetConfigName(fmt.Sprintf("cg.%s.yaml", project)) // name of config file (without extension)
	viper.SetConfigType("yaml")                             // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")                                // optionally look for config in the working directory
	err := viper.ReadInConfig()                             // Find and read the config file
	if err != nil {                                         // Handle errors reading the config file
		return fmt.Errorf("Fatal error config file: %s \n", err)
	}
	return nil
}
