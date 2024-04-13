package main

import (
	"github.com/aknorsh/go-analysis-sandbox/internal/upperid"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		upperid.UpperIDAnalyzer,
	)
}
