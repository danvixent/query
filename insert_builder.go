package query

//InsertBuilder is a builder for INSERT statements
type InsertBuilder struct {
	query string
}

//NewInsertBuilder returns a new *InsertBuilder
func NewInsertBuilder() *InsertBuilder {
	return new(InsertBuilder)
}

//Insert adds an INSERT statement to the builders'squery
func (i *InsertBuilder) Insert(table string) *InsertBuilder {
	i.query = "INSERT INTO " + table
	return i
}

//Fields adds the fields to be inserted to the builder's query
func (i *InsertBuilder) Fields(fields ...string) *InsertBuilder {
	i.query += addFields("", true, fields...)
	return i
}

//Values adds a set of values for each corresponding column to the builder's query.
//Any value for a string colmun should be wrapped in single quotes.
func (i *InsertBuilder) Values(values string) *InsertBuilder {
	i.query += " VALUES(" + values + ")"
	return i
}

//ValuesFromMap adds multiple value groups derived from ixToValues to the builder'query
//Any value for a string colmun should be wrapped in single quotes.
//Usage example:
// 		ValuesFromMap(map[int]string{
// 				0: "'Mrs','Susan','Jerome','+2319057573110'",
// 				1: "'Mr','George','Thane','+1222922843994'",
// 				2: "'Miss','Jane','Lilly','+2328145379003'",
// 		})
func (i *InsertBuilder) ValuesFromMap(ixToValues map[int]string) *InsertBuilder {
	i.query += withMap("VALUES", ixToValues, true)
	return i
}

//ValuesFromSlice adds multiple value groups derived from values to the builder'query
//Any value for a string colmun should be wrapped in single quotes.
//Usage example:
// ValuesFromSlice([]string{"'Mr','Jamie','Lannister','+13244266775'", "'Miss','Jane','Lenard','+4435356906'"}
func (i *InsertBuilder) ValuesFromSlice(values []string) *InsertBuilder {
	i.query += withSlice("VALUES", values, true)
	return i
}

//ReturnFromInserted selects fields from the temporary inserted table
func (i *InsertBuilder) ReturnFromInserted(fields ...string) *InsertBuilder {
	i.query += " RETURNING" + addFields("", false, fields...)
	return i
}

//ReturnAllFromInserted selects all fields from the temporary inserted table
func (i *InsertBuilder) ReturnAllFromInserted() *InsertBuilder {
	i.query += " RETURNING *"
	return i
}

//Clear erases the builder's query
func (i *InsertBuilder) Clear() {
	i.query = ""
}

func (i *InsertBuilder) String() string {
	return i.query
}
