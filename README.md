 # go-cp
Go-cp can copy most variable values of go

## Features
- Copy struct to struct
- Copy a pointer value to its actual type
- Copy struct slice to struct slice
- Copy struct array to struct array
- Copy primitive type value
- Elegant high scalability

## Quick Start
```go
package main

import (
	"fmt"
	"github.com/yushengji/go-copy"
)

type A struct {
	A int
	B int
}

type B struct {
	A int
	C int
}

func main() {
	aa := A{A: 1, B: 2}
	var bb B
	gocp.Cp(aa, &bb)
	// bb.A=1, bb.C=0
	fmt.Println(bb)
}

```
Let's see what happened:
- bb.A: bb.A has the same type and property names as aa.A
- bb.C: bb.C copy failed because no suitable mapping could be found in A

gocp can match the name and type in the two structures, and the copy will be successful after the match is successful.

## Plugin
Plugin is an extension mechanism provided by gocp, which are mainly divided into the following two types:
1. **TypePlugin**: Replace existing copy mechanism
2. **FieldPlugin**: Extend the existing struct attribute copy mechanism

### TypePlugin
`TypePlugin` is used to extend the copy mechanism of different Kinds. Frame currently supported Kinds are: `Int` and other basic types,
`Map`, `pointer`, `Struct`, `Array`, `Slice`.If the existing copy mechanism cannot meet the actual needs, You can implement the `TypePlugin` 
interface by yourself, and use `RegisterTypePlugin` registration to make it take effect before copying.

Implementation reference:
1. `gocp.structCopier`
2. `gocp.sliceCopier`
3. `gocp.arrayCopier`

### FieldPlugin
FieldPlugin is used for a mechanism for copying properties of extended structures. Use this plugin to change the default behavior of
the framework when copying struct. 

The framework has the following property plugins built in:
1. `gocp.namePlugin`: Unified target structure attribute name when suitable for structure copy
2. `gocp.datePlugin`: Applicable to the mutual conversion between structure attribute date(like 2023-03-11) and string
3. `gocp.timePlugin`: Applicable to the mutual conversion between structure attribute time(like 2023-03-11 17:10:00) and string

Usage example:
``` go
package main

import (
	"fmt"
	"github.com/yushengji/go-copy"
	"time"
)

type A struct {
	B time.Time
	C time.Time
	D time.Time
}

type B struct {
	BB time.Time `gocp-name:"B"`
	C  string    `gocp:"date"`
	D  string    `gocp:"time"`
}

func main() {
	gocp.RegisterFieldPlugin(gocp.NamePlugin, gocp.DateStringPlugin, gocp.TimeStringPlugin)
	aa := A{time.Now(), time.Now(), time.Now()}
	var bb B
	gocp.Cp(aa, &bb)
	fmt.Println(bb)
}
```
When you need to modify the default name mapping rules, you can use `gocp-name` to specify that it is mapped
to the specified attribute name.

When mapping between time and string is required, date/time can be selected 
for mapping according to actual scenarios.

## LICENSE
The go-copy is released under version 2.0 of the Apache License.