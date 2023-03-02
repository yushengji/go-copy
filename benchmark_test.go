package gocp

import (
	"testing"
)

type benchmarkSimpleA struct {
	AA int
	BB string
}

type benchmarkSimpleB struct {
	AA int
	BB string
}

var benchmarkSimpleAA = benchmarkSimpleA{AA: 1, BB: "xxx"}

func BenchmarkTraditionSimpleStruct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		commonSimpleCp()
	}
}

func BenchmarkCpSimpleStruct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goSimpleCp()
	}
}

func BenchmarkTraditionSimpleStructBatch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		commonSimpleCpBatch()
	}
}

func BenchmarkCpSimpleStructBatch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		goSimpleCpBatch()
	}
}

func commonSimpleCp() {
	_ = benchmarkSimpleB{
		AA: benchmarkSimpleAA.AA,
		BB: benchmarkSimpleAA.BB,
	}
}

func goSimpleCp() {
	bb := new(benchmarkSimpleB)
	Cp(benchmarkSimpleAA, bb)
}

func commonSimpleCpBatch() {
	for i := 0; i < 2000; i++ {
		_ = benchmarkSimpleB{
			AA: benchmarkSimpleAA.AA,
			BB: benchmarkSimpleAA.BB,
		}
	}
}

func goSimpleCpBatch() {
	for i := 0; i < 2000; i++ {
		bb := new(benchmarkSimpleB)
		Cp(benchmarkSimpleAA, bb)
	}
}
