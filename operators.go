package query

// Eq equates a f to v
func Eq(f string, v interface{}) string {
	return f + "=" + stringify(v, true)
}

// G add > in-between f & v
func G(f string, v interface{}) string {
	return f + ">" + stringify(v, true)
}

// L adds < in-between f & v
func L(f string, v interface{}) string {
	return f + "<" + stringify(v, true)
}

// GEq adds >= in-between f & v
func GEq(f string, v interface{}) string {
	return f + ">=" + stringify(v, true)
}

// LEq adds <= in-between f & v
func LEq(f string, v interface{}) string {
	return f + "<=" + stringify(v, true)
}

// SubQry equates f to a subquery
func SubQry(f string, v interface{ String() string }) string {
	return f + "=(" + v.String() + ")"
}

// Or prepends OR to v
// this is only intended use with WhereWithMap
func Or(v string) string {
	return "OR" + v
}

// And prepends AND to v
// this is only intended use with WhereWithMap
func And(v string) string {
	return "AND" + v
}
