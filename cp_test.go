package gocp

import "testing"

func TestCpBasic(t *testing.T) {
	srcInt := 1
	dstInt := 0
	Cp(srcInt, &dstInt)
	if dstInt != 1 {
		t.Fatalf("dstInt should be 1")
		return
	}

	srcFloat := 2.0
	dstFloat := 0.0
	Cp(srcFloat, &dstFloat)
	if dstFloat != 2.0 {
		t.Fatalf("dstFloat should be 2.0")
		return
	}

	srcString := "src"
	dstString := ""
	Cp(srcString, &dstString)
	if dstString != "src" {
		t.Fatalf("dstString should be src")
		return
	}

	dstBool := false
	Cp(true, &dstBool)
	if !dstBool {
		t.Fatalf("dstBool should be true")
		return
	}

	var srcComplex complex128 = 2
	var dstComplex complex128 = 0
	Cp(srcComplex, &dstComplex)
	if dstComplex != 2 {
		t.Fatalf("dstComplex should be 2")
		return
	}
}

func TestPtrCpBasic(t *testing.T) {
	srcInt := 1
	dstInt := 0
	Cp(&srcInt, &dstInt)
	if dstInt != 1 {
		t.Fatalf("dstInt should be 1")
		return
	}

	srcFloat := 2.0
	dstFloat := 0.0
	Cp(&srcFloat, &dstFloat)
	if dstFloat != 2.0 {
		t.Fatalf("dstFloat should be 2.0")
		return
	}

	srcString := "src"
	dstString := ""
	Cp(&srcString, &dstString)
	if dstString != "src" {
		t.Fatalf("dstString should be src")
		return
	}

	srcBool := true
	dstBool := false
	Cp(&srcBool, &dstBool)
	if !dstBool {
		t.Fatalf("dstBool should be true")
		return
	}

	var srcComplex complex128 = 2
	var dstComplex complex128 = 0
	Cp(&srcComplex, &dstComplex)
	if dstComplex != 2 {
		t.Fatalf("dstComplex should be 2")
		return
	}
}

func TestNil(t *testing.T) {
	stringSrc := "stringSrc"
	intDst := 2
	Cp(&stringSrc, &intDst)
	if intDst != 2 {
		t.Fatalf("intDst should be 2")
		return
	}

	dst := 20
	Cp(nil, &dst)
	if dst != 20 {
		t.Fatalf("dst should be 20")
	}

	Cp(20, nil)
	Cp(nil, nil)
}
