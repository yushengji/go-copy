package gocp

import (
	"testing"
	"time"
)

func TestCpBasic(t *testing.T) {
	srcInt := 1
	dstInt := 0
	Cp(srcInt, &dstInt)
	if dstInt != 1 {
		t.Fatalf("dstInt not eq srcInt")
		return
	}

	srcFloat := 2.0
	dstFloat := 0.0
	Cp(srcFloat, &dstFloat)
	if dstFloat != 2.0 {
		t.Fatalf("dstFloat not eq srcFloat")
		return
	}

	srcString := "src"
	dstString := ""
	Cp(srcString, &dstString)
	if dstString != "src" {
		t.Fatalf("dstString not eq srcString")
		return
	}

	dstBool := false
	Cp(true, &dstBool)
	if !dstBool {
		t.Fatalf("dstBool not eq true")
		return
	}

	var srcComplex complex128 = 2
	var dstComplex complex128 = 0
	Cp(srcComplex, &dstComplex)
	if dstComplex != 2 {
		t.Fatalf("dstComplex not eq srcComplex")
		return
	}
}

func TestPtrCpBasic(t *testing.T) {
	srcInt := 1
	dstInt := 0
	Cp(&srcInt, &dstInt)
	if dstInt != 1 {
		t.Fatalf("dstInt not eq srcInt")
		return
	}

	srcFloat := 2.0
	dstFloat := 0.0
	Cp(&srcFloat, &dstFloat)
	if dstFloat != 2.0 {
		t.Fatalf("dstFloat not eq 2.0")
		return
	}

	srcString := "src"
	dstString := ""
	Cp(&srcString, &dstString)
	if dstString != "src" {
		t.Fatalf("dstString not eq src")
		return
	}

	srcBool := true
	dstBool := false
	Cp(&srcBool, &dstBool)
	if !dstBool {
		t.Fatalf("dstBool not eq true")
		return
	}

	var srcComplex complex128 = 2
	var dstComplex complex128 = 0
	Cp(&srcComplex, &dstComplex)
	if dstComplex != 2 {
		t.Fatalf("dstComplex not eq srcComplex")
		return
	}
}

func TestNil(t *testing.T) {
	stringSrc := "stringSrc"
	intDst := 2
	Cp(&stringSrc, &intDst)
	if intDst != 2 {
		t.Fatalf("intDst not eq stringSrc")
		return
	}

	dst := 20
	Cp(nil, &dst)
	if dst != 20 {
		t.Fatalf("dst not eq 20")
	}

	Cp(20, nil)
	Cp(nil, nil)
}

func TestCpSimpleStruct(t *testing.T) {
	type a struct {
		AA int
		BB string
	}

	type b struct {
		AA int
		BB string
	}

	as := a{AA: 2, BB: "as"}
	bs := b{}
	Cp(as, &bs)
	if bs.AA != 2 {
		t.Fatalf("bs AA not eq as.AA")
		return
	}

	if bs.BB != "as" {
		t.Fatalf("bs AA not eq as.BB")
		return
	}
}

func TestCpObjectStruct(t *testing.T) {
	type a struct {
		Slice []int
		Map   map[string]int
		T     time.Time
	}

	type b struct {
		Slice []int
		Map   map[string]int
		T     time.Time
	}

	now := time.Now()
	as := a{Slice: []int{1, 2, 3}, Map: map[string]int{"A": 1, "b": 2}, T: now}
	bs := b{}
	Cp(as, &bs)

	if len(as.Slice) != len(bs.Slice) {
		t.Fatalf("bs.Slice len not eq as.Slice")
		return
	}

	for i := 0; i < len(as.Slice); i++ {
		if as.Slice[i] != bs.Slice[i] {
			t.Fatalf("bs[%d] should %d", i, as.Slice[i])
			return
		}
	}

	for k, v := range as.Map {
		if bs.Map[k] != v {
			t.Fatalf("bs v not eq as v")
			return
		}
	}

	if as.T != bs.T {
		t.Fatalf("as t not eq bs t")
		return
	}
}

func TestCpStructStruct(t *testing.T) {
	type c struct {
		X float64
	}

	type a struct {
		C c
	}

	type b struct {
		C c
	}

	as := a{C: c{X: 1.1}}
	bs := b{}
	Cp(as, &bs)
	if bs.C.X != 1.1 {
		t.Fatalf("bs c not eq as")
		return
	}
}

func TestCpPointerStruct(t *testing.T) {
	type c struct {
		X float64
	}

	type a struct {
		C *c
	}

	type b struct {
		C *c
	}

	as := a{C: &c{X: 1.1}}
	bs := b{}
	Cp(as, &bs)
	if bs.C.X != 1.1 {
		t.Fatalf("bs c not eq as")
		return
	}
}

func TestCpArray(t *testing.T) {
	type a struct {
		A int
	}

	type b struct {
		A int
	}

	as := [3]a{{1}, {2}, {3}}
	var bs [3]b
	Cp(as, &bs)
	for i := range as {
		if as[i].A != bs[i].A {
			t.Fatalf("as not eq bs")
			return
		}
	}

	type c struct {
		A int
		B float64
	}

	var cs [5]c
	Cp(as, &cs)
	for i := range as {
		if as[i].A != cs[i].A {
			t.Fatalf("as not eq cs")
			return
		}
	}
}
