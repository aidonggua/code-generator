package main

import (
	"code-generator/generator"
	initialize "code-generator/init"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	gen := &generator.DefaultGenerator{}

	args := os.Args
	if len(args) < 2 || len(args) > 2 || args[1] == "help" || (args[1] != "init" && args[1] != "run" && args[1] != "version") {
		fmt.Println("Usage: cg <init> | <run> | <version> | <help>")
		return
	}

	if args[1] == "version" {
		fmt.Println("v0.2.1")
		return
	}

	if args[1] == "init" {
		initializer := &initialize.Initializer{}
		initializer.Init()
		return
	}

	if args[1] == "run" {
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
		return
	}
}
