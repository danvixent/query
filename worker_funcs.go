package query

import (
	"reflect"
	"sort"
	"strconv"
)

//withMap returns a string composed of values from mapper
//if parenthesize is true, each value in mapper is parenthesized before
//adding it to qry, prefix allows for different statements(like WHERE or VALUES) to be specified
func withMap(prefix string, mapper map[int]interface{}, parenthesize bool) string {
	if mapper == nil {
		return ""
	}

	qry := ""
	if prefix == "," {
		qry = prefix
	} else {
		qry = " " + prefix
	}
	keys := make([]int, 0, len(mapper))
	for k := range mapper {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	l := len(keys) - 1

	if parenthesize {
		for ix, key := range keys {
			if ix == l {
				qry += "(" + stringifyNoQuote(mapper[key]) + ")"
				break
			}
			qry += "(" + stringifyNoQuote(mapper[key]) + "),"
		}
		return qry
	}

	for ix, key := range keys {
		if ix == l {
			qry += " " + stringifyNoQuote(mapper[key])
			break
		}
		qry += " " + stringifyNoQuote(mapper[key])
	}
	return qry
}

// values mutates all values in mapper to a one string,
// with commas seperating each value.
// obsolete keys will be deleted and the only the first index
// will remain
// mapper only supports pointers to int,uint types and the string data type.
func values(mapper map[int]interface{}) string {
	if mapper == nil {
		return ""
	}

	qry := ""
	keys := make([]int, 0, len(mapper))
	for k := range mapper {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	l := len(keys) - 1

	for ix, key := range keys {
		if ix == l {
			qry += stringifyQuote(mapper[key])
			break
		}
		qry += stringifyQuote(mapper[key]) + ","
	}
	return qry
}

//whereIn adds a WHERE clause along with IN keyword with values derived from
//the values parameter
func whereIn(field string, values ...interface{}) string {
	if values == nil {
		return ""
	}

	qry := " WHERE " + field + " IN("
	l := len(values) - 1
	for ix, v := range values {
		if ix == l {
			qry += stringifyQuote(v)
			break
		}
		qry += stringifyQuote(v) + ","
	}
	qry += ")"
	return qry
}

func addFields(prefix string, parenthesize bool, fields ...string) string {
	qry := prefix + " "
	if parenthesize {
		qry += "("
		for i, field := range fields {
			if i == len(fields)-1 {
				qry += field
				break
			}
			qry += field + ","
		}
		qry += ")"
		return qry
	}

	for i, field := range fields {
		if i == len(fields)-1 {
			qry += field
			break
		}
		qry += field + ","
	}
	return qry
}

//offset adds an OFFSET clause,
//only if num is a number
func offset(num int) string {
	return " OFFSET " + strconv.Itoa(num)
}

func limit(num int) string {
	return " LIMIT " + strconv.Itoa(num)
}

func where(cond string) string {
	qry := " WHERE " + cond
	return qry
}

func and(cond string) string {
	return " AND " + cond
}

func or(cond string) string {
	return " OR " + cond
}

// stringer allows us to avoid importing fmt
type stringer interface {
	String() string
}

func stringifyNoQuote(i interface{}) string {
	switch i.(type) {
	case string:
		return i.(string)
	case stringer:
		return i.(stringer).String()
	}

	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Ptr:
		return stringifyNoQuote(v.Elem())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	}

	return ""
}

func stringifyQuote(i interface{}) string {
	switch i.(type) {
	case string:
		return "'" + i.(string) + "'"
	case stringer:
		return "'" + i.(stringer).String() + "'"
	}

	v := reflect.ValueOf(i)

	switch v.Kind() {
	case reflect.Ptr:
		return stringifyQuote(v.Elem())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	}

	return ""
}
