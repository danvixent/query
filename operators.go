package query

// Eq equates a f to v
func Eq(f string, v interface{}) string {
	return f + "=" + quoteStringify(v)
}

// Eq add != in-between f and v
func NEq(f string, v interface{}) string {
	return f + "!=" + quoteStringify(v)
}

// G add > in-between f & v
func G(f string, v interface{}) string {
	return f + ">" + quoteStringify(v)
}

// L adds < in-between f & v
func L(f string, v interface{}) string {
	return f + "<" + quoteStringify(v)
}

// GEq adds >= in-between f & v
func GEq(f string, v interface{}) string {
	return f + ">=" + quoteStringify(v)
}

// LEq adds <= in-between f & v
func LEq(f string, v interface{}) string {
	return f + "<=" + quoteStringify(v)
}

// SubQry equates f to a subquery
func SubQry(f string, v interface{ String() string }) string {
	return f + "=(" + v.String() + ")"
}

// Or prepends OR to v
// this is only intended use with WhereWithMap
func Or(v string) string {
	return "OR " + v
}

// IsNull adds " IS NULL" to v and returns the resutl
func IsNull(v string) string {
	return v + " IS NULL"
}

// IsNotNull adds " IS NOT NULL" to v and returns the resutl
func IsNotNull(v string) string {
	return v + " IS NOT NULL"
}

// And prepends AND to v
// this is only intended use with WhereWithMap
func And(v string) string {
	return "AND " + v
}
