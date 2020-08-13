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
	i.query += addFields("", true, toInterface(fields...)...)
	return i
}

//Values adds a set of values for each corresponding column to the builder's query.
//Any value for a string colmun should be wrapped in single quotes.
func (i *InsertBuilder) Values(values ...interface{}) *InsertBuilder {
	i.query += addFields("VALUES", true, values...)
	return i
}

//ValuesFromMap adds multiple value groups derived from ixToValues to the builder'query
//Any value for a string colmun should be wrapped in single quotes.
//Usage example:
// 		ValuesFromMap(map[int]string{
// 				0: "'Mrs'",
// 				1: "'Susan'",
// 				2: "'Jerome'",
//				3: "'+2319057573110'"
// 		})
//
// Note that string values in ixToValues beginning with '(' won't be quoted
// by this method,as they will be assumed to be subqueries.
func (i *InsertBuilder) ValuesFromMap(ixToValues map[int]interface{}) *InsertBuilder {
	i.query += withMap("VALUES", concactValues(ixToValues), true)
	return i
}

// ValuesSet adds another value set without adding the VALUES keyword
//
// Note that string values in ixToValues beginning with '(' won't be quoted
// by this method, as they will be assumed to be subqueries.
func (i *InsertBuilder) ValuesSet(ixToValues map[int]interface{}) *InsertBuilder {
	i.query += withMap(",", concactValues(ixToValues), true)
	return i
}

//ReturnFromInserted selects fields from the temporary inserted table
func (i *InsertBuilder) ReturnFromInserted(fields ...interface{}) *InsertBuilder {
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
