package gocp

import "reflect"

type fieldCacheK struct {
	srcT, dstT reflect.Type
}

var filterResCache = make(map[fieldCacheK][]int)

// filterFields filter same fields return index
func filterFields(srcT, dstT reflect.Type) []int {
	k := fieldCacheK{
		srcT: srcT,
		dstT: dstT,
	}
	ret, ok := filterResCache[k]
	if ok {
		return ret
	}

	num := srcT.NumField()
	for i := 0; i < num; i++ {
		srcField := srcT.Field(i)
		dstField := dstT.Field(i)

		if !srcField.IsExported() || !dstField.IsExported() {
			continue
		}

		if srcField.Name != dstField.Name {
			continue
		}

		if srcField.Type != dstField.Type {
			continue
		}

		ret = append(ret, i)
	}
	filterResCache[k] = ret
	return ret
}
