package gocp

import "testing"

type benchmarkSimpleA struct {
	AA int
	BB string
}

type benchmarkSimpleB struct {
	AA int
	BB string
}

var benchmarkSimpleAA = benchmarkSimpleA{AA: 1, BB: "xxx"}

func BenchmarkTraditionStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		commonSimpleCp()
	}
}

func BenchmarkCpStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goSimpleCp()
	}
}

func commonSimpleCp() benchmarkSimpleB {
	return benchmarkSimpleB{
		AA: benchmarkSimpleAA.AA,
		BB: benchmarkSimpleAA.BB,
	}
}

func goSimpleCp() benchmarkSimpleB {
	bb := new(benchmarkSimpleB)
	Cp(benchmarkSimpleAA, bb)
	return *bb
}
