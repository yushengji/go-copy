package gocp

func ptr(src, dst *entity) {
	srcElem := src.tpe().Elem()
	dstElem := dst.tpe().Elem()

	if allValType(srcElem.Kind()) && allValType(dstElem.Kind()) {
		setVal(src, dst)
		return
	}
}
