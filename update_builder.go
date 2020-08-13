package query

//UpdateBuilder is a builder for UPDATE statements
type UpdateBuilder struct {
	query string
}

//NewUpdateBuilder returns a new *UpdateBuilder
func NewUpdateBuilder() *UpdateBuilder {
	return new(UpdateBuilder)
}

//Update adds an UPDATE statemens to the builder's query
//table represents the database table to update.
func (u *UpdateBuilder) Update(table string) *UpdateBuilder {
	u.query = "UPDATE " + table
	return u
}

//Set adds a field and its new value to the builder's query
//Update must be called prior to Set.
func (u *UpdateBuilder) Set(field string) *UpdateBuilder {
	u.query += " SET " + field
	return u
}

//SetFromMap adds a WHERE clause to the builder's query with fields and new values
//derived from ixToField which should map integers(allows for proper ordering)
//to fields equating their new values.
//
//You should use consecutive integers starting from zero.
//
//Usage example:
//	SetFromMap(map[int]string{
//			0: "FirstName='Daniel'",
//			 1: "LastName='Jamie'",
//	})
//
// Note that string values in ixToValues beginning with '(' won't be quoted
// by this method, as they will be assumed to be subqueries.
func (u *UpdateBuilder) SetFromMap(ixToField map[int]interface{}) *UpdateBuilder {
	u.query += withMap("SET", ixToField, false)
	return u
}

//Where adds a WHERE clause to the builder's query.
//condition is the desired condition
func (u *UpdateBuilder) Where(condition string) *UpdateBuilder {
	u.query += where(condition)
	return u
}

//WhereWithMap adds a WHERE clause to the builder's query with fields and conditions
//derived from ixToCond which should map integers(allows for proper ordering)
//to conditions desired to be met.
//You should use consecutive integers starting from zero.
//Usage example:
//	WhereWithMap(map[int]string{
//			0: "CategoryID=3 OR",
//			1: "BarcodeID=22",
//	})
func (u *UpdateBuilder) WhereWithMap(ixToCond map[int]interface{}) *UpdateBuilder {
	u.query += withMap("WHERE", ixToCond, false)
	return u
}

//ReturnFromInserted selects fields from the temporary inserted table
func (u *UpdateBuilder) ReturnFromInserted(fields ...string) *UpdateBuilder {
	u.query += " RETURNING" + addFields("", false, toInterface(fields...)...)
	return u
}

//ReturnAllFromInserted selects all fields from the temporary inserted table
func (u *UpdateBuilder) ReturnAllFromInserted() *UpdateBuilder {
	u.query += " RETURNING *"
	return u
}

//And is an alternative to WhereWithMap
//
//it adds an AND along with the condition specified
func (u *UpdateBuilder) And(condition string) *UpdateBuilder {
	u.query += and(condition)
	return u
}

//Or is an alternative to WhereWithMap
//
//it adds an OR along with the condition specified
func (u *UpdateBuilder) Or(condition string) *UpdateBuilder {
	u.query += or(condition)
	return u
}

//Clear erases the builder's query
func (u *UpdateBuilder) Clear() {
	u.query = ""
}

func (u *UpdateBuilder) String() string {
	return u.query
}
