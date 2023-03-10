package gocp

import (
	"fmt"
	"time"
)

type A struct {
	A int
	B time.Time
	C time.Time
}

type B struct {
	A  int
	BB time.Time `gocp-name:"B"`
	C  time.Time `gocp:"time"`
}

func main() {
	RegisterFieldPlugin(NamePlugin, TimeStringPlugin)
	aa := A{1, time.Now(), time.Now()}
	var bb B
	Cp(aa, &bb)
	fmt.Println(bb)
}
