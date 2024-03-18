package main

import (
	"code-generator/generator"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	gen := &generator.DefaultGenerator{}

	viper.SetConfigName(fmt.Sprintf(".cg/config.yaml")) // name of config file (without extension)
	viper.SetConfigType("yaml")                         // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")                            // optionally look for config in the working directory
	err := viper.ReadInConfig()                         // Find and read the config file
	if err != nil {                                     // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	err = gen.LoadConfig()
	if err != nil {
		panic(err)
	}

	gen.Generate()
}
