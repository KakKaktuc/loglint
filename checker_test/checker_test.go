package checker_test

import (
	"testing"

	"loglint/checker"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLogLint(t *testing.T) {
	testdata := analysistest.TestData() // ищет папку testdata
	analysistest.Run(t, testdata, checker.Analyzer, "src")
}
