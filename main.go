package main

import "code-generator/generator"

func main() {
	javaGen := &generator.JavaGenerator{}
	err := javaGen.LoadConfig("sample")
	if err != nil {
		panic(err)
	}
	javaGen.Generate()
}
