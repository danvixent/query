package query

import (
	"sort"
	"strconv"
)

const quote = 1
const noquote = 0

//withMap returns a string composed of values from mapper
//if parenthesize is true, each value in mapper is parenthesized before
//adding it to qry, prefix allows for different statements(like WHERE or VALUES) to be specified
func withMap(prefix string, mapper map[int]interface{}, parenthesize bool) string {
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
				qry += "(" + stringify(v, false) + ")"
				break
			}
			qry += "(" + stringify(v, false) + "),"
		}
		return qry
	}

	for ix, key := range keys {
		v := mapper[key]
		if ix == l {
			qry += " " + stringify(v, false)
			break
		}
		qry += " " + stringify(v, false)
	}
	return qry
}

// concactValues mutates all values in mapper to a one string,
// with commas seperating each value.
// obsolete keys will be deleted and the only the first index
// will remain
// mapper only supports pointers to int,uint types and the string data type.
func concactValues(mapper map[int]interface{}) map[int]interface{} {
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
			qry += stringify(v, true)
			break
		}
		qry += stringify(v, true) + ","
	}
	return map[int]interface{}{0: qry}
}

//whereIn adds a WHERE clause along with IN keyword with values derived from
//the values parameter
func whereIn(field string, values ...interface{}) string {
	qry := " WHERE " + field + " IN("
	l := len(values) - 1
	for ix, v := range values {
		if ix == l {
			qry += stringify(v, true)
			break
		}
		qry += stringify(v, true) + ","
	}
	qry += ")"
	return qry
}

//addFields adds the values in fields to qry
//if  parenthesize is true prefix isn't added
func addFields(prefix string, parenthesize bool, fields ...interface{}) string {
	l := len(fields) - 1
	if parenthesize {
		qry := " ("
		for i, v := range fields {
			if i == l {
				qry += stringify(v, false)
				break
			}
			qry += stringify(v, false) + ","
		}
		qry += ")"
		return qry
	}

	qry := prefix + " "
	for i, v := range fields {
		if i == l {
			qry += stringify(v, false)
			break
		}
		qry += stringify(v, false) + ","
	}
	return qry
}

//offset adds an OFFSET clause,
//only if num is a number
func offset(num *int32) *string {
	qry := " OFFSET " + strconv.FormatInt(int64(*num), 10)
	return &qry
}

func limit(num *int32) *string {
	qry := " LIMIT " + strconv.FormatInt(int64(*num), 10)
	return &qry
}

func where(cond string) string {
	qry := " WHERE " + cond
	return qry
}

func and(cond string) string {
	qry := " AND " + cond
	return qry
}

func or(cond string) string {
	qry := " OR " + cond
	return qry
}

func toInterface(values ...string) []interface{} {
	v := make([]interface{}, 0, len(values))
	for i := range values {
		v = append(v, &values[i])
	}
	return v
}

type Stringer interface {
	String() string
}

//stringify converts any *int,*uint type to its string equivalent
//if a non-pointer type is passed, an empty string is returned
func stringify(i interface{}, quote bool) string {
	switch i.(type) {
	case *int32:
		return strconv.FormatInt(int64(*i.(*int32)), 10)
	case *string:
		s := i.(*string)
		if quote && !((*s)[0] == '(') {
			return "'" + *s + "'"
		}
		return *s
	case string:
		s := i.(string)
		if quote && !(s[0] == '(') {
			return "'" + s + "'"
		}
		return s
	case Stringer:
		return i.(Stringer).String()
	case *int:
		return strconv.Itoa(*i.(*int))
	case int:
		return strconv.Itoa(i.(int))
	case *int64:
		return strconv.FormatInt(*i.(*int64), 10)
	case *int8:
		return strconv.Itoa(int(*i.(*int8)))
	case *int16:
		return strconv.Itoa(int(*i.(*int16)))
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case int8:
		return strconv.Itoa(int(i.(int8)))
	case int16:
		return strconv.Itoa(int(i.(int16)))
	case int32:
		return strconv.FormatInt(int64(i.(int32)), 10)
	case *uint:
		return strconv.FormatUint(uint64(*i.(*uint)), 10)
	case *uint8:
		return strconv.FormatUint(uint64(*i.(*uint8)), 10)
	case *uint16:
		return strconv.FormatUint(uint64(*i.(*uint16)), 10)
	case *uint32:
		return strconv.FormatUint(uint64(*i.(*uint32)), 10)
	case *uint64:
		return strconv.FormatUint(uint64(*i.(*uint64)), 10)
	case uint:
		return strconv.FormatUint(uint64(i.(uint)), 10)
	case uint8:
		return strconv.FormatUint(uint64(i.(uint8)), 10)
	case uint16:
		return strconv.FormatUint(uint64(i.(uint16)), 10)
	case uint32:
		return strconv.FormatUint(uint64(i.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(uint64(i.(uint64)), 10)
	}
	return ""
}
