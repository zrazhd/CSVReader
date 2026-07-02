package internal

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	path := "../cmd/data.csv"
	for i := 0; i < b.N; i++ {
		Run(path)
	}
}
