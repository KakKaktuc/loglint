package main

import (
	"github.com/KakKaktuc/loglint/checker"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(checker.Analyzer)
}
