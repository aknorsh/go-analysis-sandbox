package main

import (
	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/aknorsh/go-analysis-sandbox/internal/noliteral"
	"github.com/aknorsh/go-analysis-sandbox/internal/upperid"
)

func main() {
	multichecker.Main(
		upperid.Analyzer,
		noliteral.Analyzer,
	)
}
