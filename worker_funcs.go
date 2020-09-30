package query

import (
	"reflect"
	"sort"
	"strconv"
)

const quote = 1
const noquote = 0

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
			v := mapper[key]
			if ix == l {
				qry += "(" + noQuoteStringify(v) + ")"
				break
			}
			qry += "(" + noQuoteStringify(v) + "),"
		}
		return qry
	}

	for ix, key := range keys {
		v := mapper[key]
		if ix == l {
			qry += " " + noQuoteStringify(v)
			break
		}
		qry += " " + noQuoteStringify(v)
	}
	return qry
}

// concactValues mutates all values in mapper to a one string,
// with commas seperating each value.
// obsolete keys will be deleted and the only the first index
// will remain
// mapper only supports pointers to int,uint types and the string data type.
func concactValues(mapper map[int]interface{}) map[int]interface{} {
	if mapper == nil {
		return nil
	}

	qry := ""
	keys := make([]int, 0, len(mapper))
	for k := range mapper {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	l := len(keys) - 1

	for ix, key := range keys {
		v := mapper[key]
		if ix == l {
			qry += quoteStringify(v)
			break
		}
		qry += quoteStringify(v) + ","
	}
	return map[int]interface{}{0: qry}
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
			qry += quoteStringify(v)
			break
		}
		qry += quoteStringify(v) + ","
	}
	qry += ")"
	return qry
}

//addFields adds the values in fields to qry
//if  parenthesize is true prefix isn't added
func addFields(prefix string, parenthesize bool, fields ...interface{}) string {
	if fields == nil {
		return ""
	}

	qry := prefix + " "
	l := len(fields) - 1
	if parenthesize {
		qry += "("
		for i, v := range fields {
			if i == l {
				qry += noQuoteStringify(v)
				break
			}
			qry += noQuoteStringify(v) + ","
		}
		qry += ")"
		return qry
	}

	for i, v := range fields {
		if i == l {
			qry += noQuoteStringify(v)
			break
		}
		qry += noQuoteStringify(v) + ","
	}
	return qry
}

func addFieldsString(prefix string, parenthesize bool, fields ...string) string {
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

func noQuoteStringify(i interface{}) string {
	switch i.(type) {
	case stringer:
		return "'" + i.(stringer).String() + "'"
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Ptr:
		return tostr(v.Elem())
	case reflect.String:
		return v.String()
	default:
		return tostr(v)
	}
}

func tostr(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	default:
		return ""
	}
}

func quoteStringify(i interface{}) string {
	switch i.(type) {
	case stringer:
		return "'" + i.(stringer).String() + "'"
	}

	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Ptr:
		return tostr(v.Elem())
	case reflect.String:
		return "'" + v.String() + "'"
	default:
		return tostr(v)
	}
}
