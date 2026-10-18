package main

import (
	"os"

	"github.com/NurkhanZ/LevelGen/internal/cli"
	"github.com/NurkhanZ/LevelGen/internal/generator"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr, generator.PlatformGenerator{}))
}
