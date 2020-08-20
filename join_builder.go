package query

//JoinBuilder is a qury builder for JOIN clauses
type JoinBuilder struct {
	s *SelectBuilder
}

//NewJoinBuilder returns a new JoinBuilder
func NewJoinBuilder() *JoinBuilder {
	j := new(JoinBuilder)
	j.s = NewSelectBuilder()
	return j
}

//Join adds a JOIN clause to the builder's query
//table represents the name of the table
//to join to.
func (j *JoinBuilder) Join(table string) *JoinBuilder {
	j.s.query += " JOIN " + table
	return j
}

// Using adds a using clause to the builder's query
func (j *JoinBuilder) Using(fields ...string) *JoinBuilder {
	j.s.query += addFields("USING", false, toInterface(fields...)...)
	return j
}

//On adds the matching colmuns in joined tables.
func (j *JoinBuilder) On(column1 string, column2 string) *JoinBuilder {
	j.s.query += " ON " + column1 + "=" + column2
	return j
}

//As sets an alias for a table,
//Alternatively the alias could be set beside the table name while
//adding the table to the builder's query
func (j *JoinBuilder) As(alias string) *JoinBuilder {
	j.s.query += " AS " + alias
	return j
}

//FromSelectBuilder sets j's internal *SelectBuilder to s
func (j *JoinBuilder) FromSelectBuilder(s *SelectBuilder) *JoinBuilder {
	j.s = s
	return j
}

//Select adds a select statement to the builder's query
func (j *JoinBuilder) Select(fields ...string) *JoinBuilder {
	j.s.Select(fields...)
	return j
}

//SelectAll adds a SELECT * FROM statement to the builder's query
//table is the database to select all from.
func (j *JoinBuilder) SelectAll(table string) *JoinBuilder {
	j.s.SelectAll(table)
	return j
}

//From adds the table to select from to the builder's query,
//Select MUST be called prior to From, on the
//same *JoinBuilder.
func (j *JoinBuilder) From(table string) *JoinBuilder {
	j.s.From(table)
	return j
}

//Where adds a WHERE clause to the builder's query
//From MUST be called prior to Where, on the
//same *JoinBuilder.
//
//Examples : j.Where("id=2"),  j.Where("name='Danny'")
//
//String values MUST be quoted with single-quotes
func (j *JoinBuilder) Where(condition string) *JoinBuilder {
	j.s.Where(condition)
	return j
}

//WhereWithMap adds a WHERE clause to the builder's query with  conditions
//derived from ixToCond, which should map integers(allows for proper ordering)
//to conditions desired to be met.
//You should use consecutive integers starting from zero.
//Usage example:
//	WhereWithMap(map[int]string{
//			0: "CategoryID=3 OR",
//			1: "BarcodeID=22",
//	})
func (j *JoinBuilder) WhereWithMap(ixToCond map[int]interface{}) *JoinBuilder {
	j.s.WhereWithMap(ixToCond)
	return j
}

//WhereFieldIn adds a WHERE clause along with an IN operator
func (j *JoinBuilder) WhereFieldIn(field string, values ...interface{}) *JoinBuilder {
	j.s.WhereFieldIn(field, values...)
	return j
}

//And is an alternative to WhereWithMap
//
//it adds an AND along with the condition specified
func (j *JoinBuilder) And(condition string) *JoinBuilder {
	j.s.And(condition)
	return j
}

//Offset adds AN OFFSET clause to the query
func (j *JoinBuilder) Offset(num *int32) *JoinBuilder {
	j.s.Offset(num)
	return j
}

//Limit adds a LIMIT clause to the query
func (j *JoinBuilder) Limit(num *int32) *JoinBuilder {
	j.s.Limit(num)
	return j
}

//Distinct adds a DISTINCT clause the builder's query
func (j *JoinBuilder) Distinct(fields ...string) *JoinBuilder {
	j.s.Distinct(fields...)
	return j
}

//Or is an alternative to WhereWithMap
//
//it adds an OR along with the condition specified
func (j *JoinBuilder) Or(condition string) *JoinBuilder {
	j.s.Or(condition)
	return j
}

//OrderBy adds an ORDER BY clause to the builder's query
func (j *JoinBuilder) OrderBy(field string) *JoinBuilder {
	j.s.OrderBy(field)
	return j
}

//Asc adds ASC for ordering
func (j *JoinBuilder) Asc() *JoinBuilder {
	j.s.Asc()
	return j
}

//Desc adds DESC for ordering
func (j *JoinBuilder) Desc() *JoinBuilder {
	j.s.Desc()
	return j
}

//GroupBy adds a GROUP BY clause to the builder's query
func (j *JoinBuilder) GroupBy(field string) *JoinBuilder {
	j.s.GroupBy(field)
	return j
}

//Clear erases the builder's query
func (j *JoinBuilder) Clear() {
	j.s.Clear()
}

func (j *JoinBuilder) String() string {
	return j.s.String()
}
