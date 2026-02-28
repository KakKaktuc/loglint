package main

import (
	"loglint/checker"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(checker.Analyzer)
}
