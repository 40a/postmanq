package main

import (
	"github.com/AdOnWeb/postmanq"
	"flag"
	"fmt"
)


func main() {
	var file string
	flag.StringVar(&file, "f", postmanq.ExampleConfigYaml, "configuration yaml file")
	flag.Parse()

	app := postmanq.NewPostApplication()
	if app.IsValidConfigFilename(file) {
		app.SetConfigFilename(file)
		app.Run()
	} else {
		fmt.Printf("Usage: postmanq -f %s\n", postmanq.ExampleConfigYaml)
		flag.VisitAll(postmanq.PrintUsage)
	}
}
