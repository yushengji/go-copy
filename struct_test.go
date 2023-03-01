package gocp

import (
	"testing"
	"time"
)

func TestSimpleStruct(t *testing.T) {
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
	stru(&entity{o: as}, &entity{o: &bs})
	if bs.AA != 2 {
		t.Fatalf("bs AA should 2")
		return
	}

	if bs.BB != "as" {
		t.Fatalf("bs AA should as")
		return
	}
}

func TestObjectStruct(t *testing.T) {
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
	as := a{Slice: []int{1, 2, 3}, Map: map[string]int{"a": 1, "b": 2}, T: now}
	bs := b{}
	stru(&entity{o: as}, &entity{o: &bs})

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

func TestStructStruct(t *testing.T) {
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
	stru(&entity{o: as}, &entity{o: &bs})
	if bs.C.X != 1.1 {
		t.Fatalf("bs c should 1.1")
		return
	}
}

func TestPointerStruct(t *testing.T) {
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
	stru(&entity{o: as}, &entity{o: &bs})
	if bs.C.X != 1.1 {
		t.Fatalf("bs c should 1.1")
		return
	}
}
