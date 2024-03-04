package main

import (
	"code-generator/generator"
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	javaGen := &generator.JavaGenerator{}

	viper.SetConfigName(fmt.Sprintf("cg.yaml")) // name of config file (without extension)
	viper.SetConfigType("yaml")                 // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")                    // optionally look for config in the working directory
	err := viper.ReadInConfig()                 // Find and read the config file
	if err != nil {                             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Printf("project name: %s\n", viper.GetString("name"))
	fmt.Printf("language: %s\n", viper.GetString("language"))

	err = javaGen.LoadConfig()
	if err != nil {
		panic(err)
	}

	javaGen.Generate()
}
