package query

//DeleteBuilder is a builder for DELETE statements
type DeleteBuilder struct {
	query string
}

//NewDeleteBuilder returns a new *DeleteBuilder
func NewDeleteBuilder() *DeleteBuilder {
	return new(DeleteBuilder)
}

//Delete adds a DELETE statment to the builder's query
//table is the database table to delete from
func (d *DeleteBuilder) Delete(table string) *DeleteBuilder {
	d.query = "DELETE FROM " + table
	return d
}

//Where adds a WHERE clause to u's query.
//condition is the desired condition
func (d *DeleteBuilder) Where(condition string) *DeleteBuilder {
	d.query += where(condition)
	return d
}

//WhereWithMap adds a WHERE clause to u's query with fields and conditions
//derived from ixToCond which should map integers(allows for proper ordering)
//to conditions desired to be met.
//You should use consecutive integers starting from zero.
func (d *DeleteBuilder) WhereWithMap(ixToCond map[int]interface{}) *DeleteBuilder {
	d.query += withMap("WHERE", ixToCond, false)
	return d
}

//WhereFieldIn adds a WHERE clause along with an IN operator
func (d *DeleteBuilder) WhereFieldIn(field string, values []string) *DeleteBuilder {
	d.query += whereIn(field, values)
	return d
}

//And is an alternative to WhereWithMap
//
//it adds an AND along with the condition specified
func (d *DeleteBuilder) And(condition string) *DeleteBuilder {
	d.query += and(condition)
	return d
}

//Or is an alternative to WhereWithMap
//
//it adds an OR along with the condition specified
func (d *DeleteBuilder) Or(condition string) *DeleteBuilder {
	d.query += or(condition)
	return d
}

//Clear erases the builder's query
func (d *DeleteBuilder) Clear() {
	d.query = ""
}

func (d *DeleteBuilder) String() string {
	return d.query
}
