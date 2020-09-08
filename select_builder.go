package query

//SelectBuilder is bulider for select statement
type SelectBuilder struct {
	query string
}

//NewSelectBuilder returns a pointer to a new SelectBuilder
func NewSelectBuilder() *SelectBuilder {
	return new(SelectBuilder)
}

//Select adds a select statement to the builder's query
func (s *SelectBuilder) Select(fields ...string) *SelectBuilder {
	s.query = addFields("SELECT", false, toInterface(fields...)...)
	return s
}

//SelectAll adds a SELECT * statement to the builder's query
func (s *SelectBuilder) SelectAll(table string) *SelectBuilder {
	s.query = "SELECT * FROM " + table
	return s
}

//From adds the table to select from to the builder's query,
//Select MUST be called prior to From, on the
//same *SelectBuilder.
func (s *SelectBuilder) From(table string) *SelectBuilder {
	s.query += " FROM " + table
	return s
}

//Where adds a WHERE clause to the builder's query
//From MUST be called prior to Where, on the
//same *SelectBuilder.
func (s *SelectBuilder) Where(condition string) *SelectBuilder {
	s.query += where(condition)
	return s
}

//WhereWithMap adds a WHERE clause to the builder's query with fields and conditions
//derived from fieldToCond, which should map fields to conditions desired to
//be met.
//You should use consecutive integers starting from zero.
//
//Usage example:
//	WhereWithMap(map[int]string{
//			0: "CategoryID=3 OR",
//			1: "BarcodeID=22",
//	})
func (s *SelectBuilder) WhereWithMap(ixToCond map[int]interface{}) *SelectBuilder {
	s.query += withMap("WHERE", ixToCond, false)
	return s
}

//WhereFieldIn adds a WHERE clause along with an IN operator
func (s *SelectBuilder) WhereFieldIn(field string, values ...interface{}) *SelectBuilder {
	s.query += whereIn(field, values...)
	return s
}

//And is an alternative to WhereWithMap
//
//it adds an AND along with the condition specified
func (s *SelectBuilder) And(condition string) *SelectBuilder {
	s.query += and(condition)
	return s
}

//Offset  adds AN OFFSET clause to the query
func (s *SelectBuilder) Offset(num int) *SelectBuilder {
	s.query += offset(num)
	return s
}

//Limit adds a LIMIT clause to the query
func (s *SelectBuilder) Limit(num int) *SelectBuilder {
	s.query += limit(num)
	return s
}

//Or is an alternative to WhereWithMap
//
//it adds an OR along with the condition specified
func (s *SelectBuilder) Or(condition string) *SelectBuilder {
	s.query += or(condition)
	return s
}

//OrderBy adds an ORDER BY clause to the builder's query
func (s *SelectBuilder) OrderBy(field string) *SelectBuilder {
	s.query += " ORDER BY " + field
	return s
}

//GroupBy adds a GROUP BY clause the builder's query
func (s *SelectBuilder) GroupBy(field string) *SelectBuilder {
	s.query += " GROUP BY " + field
	return s
}

//Asc adds ASC for ordering
func (s *SelectBuilder) Asc() *SelectBuilder {
	s.query += " ASC"
	return s
}

//Desc adds DESC for ordering
func (s *SelectBuilder) Desc() *SelectBuilder {
	s.query += " DESC"
	return s
}

//Distinct adds a DISTINCT clause the builder's query
func (s *SelectBuilder) Distinct(fields ...string) *SelectBuilder {
	s.query += addFields("DISTINCT", false, toInterface(fields...)...)
	return s
}

//Clear erases the builder's query
func (s *SelectBuilder) Clear() {
	s.query = ""
}

func (s *SelectBuilder) String() string {
	return s.query
}
