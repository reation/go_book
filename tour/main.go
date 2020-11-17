package main

import (
	"flag"
	"log"
)

func main(){
	flag.Parse()
	var name string

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	goCmd.StringVar(&name, "name", "go name", "help")

	phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
	phpCmd.StringVar(&name, "n", "php name", "help")

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0]{
	case "go":
		_ = goCmd.Parse(args[1:])
	case "php":
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name:%s", name)
}
