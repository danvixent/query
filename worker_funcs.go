package query

import (
	"sort"
	"strconv"
)

//withMap returns a string composed of values from mapper
//if parenthesize is true, each value in mapper is parenthesized before
//adding it to qry, prefix allows for different statements(like WHERE or VALUES) to be specified
func withMap(prefix string, mapper map[int]string, parenthesize bool) string {
	qry := " " + prefix
	keys := make([]int, 0, len(mapper))
	for k := range mapper {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	l := len(keys) - 1

	if parenthesize {
		for ix, key := range keys {
			value := mapper[key]
			if ix == l {
				qry += "(" + value + ")"
				break
			}
			qry += "(" + value + "),"
		}
		return qry
	}

	for ix, key := range keys {
		value := mapper[key]
		if ix == l {
			qry += " " + value
			break
		}
		qry += " " + value
	}
	return qry
}

//withSlice is like withMap but with slices of strings
func withSlice(prefix string, values []string, parenthesize bool) string {
	l := len(values) - 1
	qry := " " + prefix

	if parenthesize {
		for ix, value := range values {
			if ix == l {
				qry += "(" + value + ")"
				break
			}
			qry += "(" + value + "),"
		}
		return qry
	}

	for ix, value := range values {
		if ix == l {
			qry += " " + value
			break
		}
		qry += " " + value
	}
	return qry
}

//whereIn adds a WHERE clause along with IN keyword with values derived from
//the values parameter
func whereIn(field string, values []string) string {
	qry := " WHERE " + field + " IN("
	l := len(values) - 1
	for ix, value := range values {
		if ix == l {
			qry += value
			break
		}
		qry += value + ","
	}
	qry += ")"
	return qry
}

//addFields adds the values in fields to qry
//if  parenthesize is true prefix isn't added
func addFields(prefix string, parenthesize bool, fields ...string) string {
	if parenthesize {
		qry := " ("
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

	qry := prefix + " "
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
