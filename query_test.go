package query

import (
	"math/rand"
	"testing"
)

func TestUpdateBuilder_Update(t *testing.T) {
	tests := []struct {
		name string
		want string
		exec func() string
	}{
		{
			"update1",
			"UPDATE Stock.Product SET ProductName='Powersuper Battery' WHERE ProductID=99",
			NewUpdateBuilder().Update("Stock.Product").Set("ProductName='Powersuper Battery'").Where("ProductID=99").String,
		},
		{
			"update2",
			"UPDATE Stock.ProductBarCodes SET Barcode='4353532242' WHERE BarcodeID=2",
			NewUpdateBuilder().Update("Stock.ProductBarCodes").Set("Barcode='4353532242'").Where("BarcodeID=2").String,
		},
		{
			"update3",
			"UPDATE Stock.Product SET ProductName='Pakgen Bulbs' WHERE CategoryID=3 OR BarcodeID=22",
			NewUpdateBuilder().Update("Stock.Product").Set("ProductName='Pakgen Bulbs'").
				WhereWithMap(map[int]interface{}{
					0: "CategoryID=3 OR",
					1: "BarcodeID=22",
				}).String,
		},
		{
			"update4",
			"UPDATE Person.Contact SET FirstName='Daniel' LastName='Jamie' WHERE ContactID=1",
			NewUpdateBuilder().Update("Person.Contact").SetFromMap(map[int]interface{}{
				0: "FirstName='Daniel'",
				1: "LastName='Jamie'",
			}).Where("ContactID=1").String,
		},
		{
			"update5",
			"UPDATE Sales.OrderDetail SET OrderDetailID=33 WHERE ProductID=3 AND Quantity=400 OR UnitPrice=300",
			NewUpdateBuilder().Update("Sales.OrderDetail").Set("OrderDetailID=33").Where("ProductID=3").And("Quantity=400").Or("UnitPrice=300").String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exec(); got != tt.want {
				t.Errorf("got = {%v} \n want = {%v}", got, tt.want)
			}
		})
	}
}

func TestSelectBuilder_Select(t *testing.T) {
	tests := []struct {
		name string
		want string
		exec func() string
	}{
		{
			"select1",
			"SELECT * FROM Person.Address",
			NewSelectBuilder().SelectAll("Person.Address").String,
		},
		{
			"select2",
			"SELECT ContactID,Title,FirstName,LastName,PhoneNumber FROM Person.Contact WHERE ContactID>100 ORDER BY FirstName",
			NewSelectBuilder().Select("ContactID", "Title", "FirstName", "LastName", "PhoneNumber").From("Person.Contact").
				Where("ContactID>100").OrderBy("FirstName").String,
		},
		{
			"select3",
			"SELECT OrderID,StoreID,OrderDate,DueDate,TotalAmountDue,PaymentDate,PaymentMethodID FROM Sales.OrderHeader WHERE OrderID>2 AND StoreID<=100 AND DueDate!=11/02/2020",
			NewSelectBuilder().Select("OrderID", "StoreID", "OrderDate", "DueDate", "TotalAmountDue", "PaymentDate", "PaymentMethodID").
				From("Sales.OrderHeader").WhereWithMap(map[int]interface{}{
				0: "OrderID>2 AND",
				1: "StoreID<=100 AND",
				2: "DueDate!=11/02/2020",
			}).String,
		},
		{
			"select4",
			"SELECT * FROM Person.Contact WHERE ContactID>3 AND AddressID=33 OR FirstName='Kelly'",
			NewSelectBuilder().SelectAll("Person.Contact").Where("ContactID>3").And("AddressID=33").Or("FirstName='Kelly'").String,
		},
		{
			"select5",
			"SELECT * FROM Stock.Product WHERE ProductID IN(2,44,22,11,42,53)",
			NewSelectBuilder().SelectAll("Stock.Product").WhereFieldIn("ProductID", 2, 44, 22, 11, 42, 53).String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exec(); got != tt.want {
				t.Errorf("got = {%v} \n want = {%v}", got, tt.want)
			}
		})
	}
}

func TestJoinBuilder_Join(t *testing.T) {
	tests := []struct {
		name string
		want string
		exec func() string
	}{
		{
			"join1",
			"SELECT soh.OrderID,ss.StoreName,soh.OrderDate,soh.TotalAmountDue,soh.DeliveryDate,soh.PaymentDate,mpm.PaymentMethod FROM Sales.OrderHeader AS soh JOIN Sales.Store AS ss ON soh.StoreID=ss.StoreID JOIN Management.PaymentMethods AS mpm ON soh.PaymentMethodID=mpm.PaymentMethodID WHERE soh.TotalAmountDue>10000 AND soh.OrderID>22 AND mpm.PaymentMethod=Cash ORDER BY soh.OrderID",
			NewJoinBuilder().
				Select("soh.OrderID", "ss.StoreName", "soh.OrderDate", "soh.TotalAmountDue", "soh.DeliveryDate", "soh.PaymentDate", "mpm.PaymentMethod").
				From("Sales.OrderHeader").As("soh").Join("Sales.Store").As("ss").On("soh.StoreID", "ss.StoreID").
				Join("Management.PaymentMethods").As("mpm").On("soh.PaymentMethodID", "mpm.PaymentMethodID").
				WhereWithMap(map[int]interface{}{
					0: "soh.TotalAmountDue>10000 AND",
					1: "soh.OrderID>22 AND",
					2: "mpm.PaymentMethod=Cash",
				}).OrderBy("soh.OrderID").String,
		},
		{
			"join2",
			"SELECT pc.FirstName,pc.LastName,pc.PhoneNumber,pc.Email FROM Purchasing.Supplier AS ps JOIN Person.Contact AS pc ON ps.ContactID=pc.ContactID WHERE pc.FirstName='Boluwatife' AND pc.LastName='Oyeniran' ORDER BY pc.FirstName",
			NewJoinBuilder().
				Select("pc.FirstName", "pc.LastName", "pc.PhoneNumber", "pc.Email").From("Purchasing.Supplier").As("ps").
				Join("Person.Contact").As("pc").On("ps.ContactID", "pc.ContactID").Where("pc.FirstName='Boluwatife'").And("pc.LastName='Oyeniran'").OrderBy("pc.FirstName").String,
		},
		{
			"join3",
			"SELECT * FROM Sales.OrderDetail AS sod JOIN Sales.OrderHeader AS soh ON sod.OrderID=soh.OrderID GROUP BY soh.OrderID",
			NewJoinBuilder().FromSelectBuilder(NewSelectBuilder()).SelectAll("Sales.OrderDetail").As("sod").Join("Sales.OrderHeader").As("soh").On("sod.OrderID", "soh.OrderID").GroupBy("soh.OrderID").String,
		},
		{
			"join4",
			"SELECT * FROM Sales.OrderDetail WHERE OrderID>100 OR TotalAmountDue>90000",
			NewJoinBuilder().SelectAll("Sales.OrderDetail").Where("OrderID>100").Or("TotalAmountDue>90000").String,
		},
		{
			"join5",
			"SELECT * FROM Sales.OrderDetail WHERE OrderID IN(32,76,33,44)",
			NewJoinBuilder().SelectAll("Sales.OrderDetail").WhereFieldIn("OrderID", "32", "76", "33", "44").String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exec(); got != tt.want {
				t.Errorf("got = {%v} \n want = {%v}", got, tt.want)
			}
		})
	}
}

func TestInsertBuilder_Insert(t *testing.T) {
	tests := []struct {
		name string
		want string
		exec func() string
	}{
		{
			"insert1",
			"INSERT INTO Person.Contact (Title,FirstName,LastName,PhoneNumber) VALUES('Mrs','Susan','Jerome','+2319057573110'),('Mr','George','Thane','+1222922843994')",
			NewInsertBuilder().Insert("Person.Contact").Fields("Title", "FirstName", "LastName", "PhoneNumber").
				ValuesFromMap(map[int]interface{}{
					0: "'Mrs'",
					1: "'Susan'",
					2: "'Jerome'",
					3: "'+2319057573110'",
				}).ValuesSet(map[int]interface{}{
				0: "'Mr'",
				1: "'George'",
				2: "'Thane'",
				3: "'+1222922843994'",
			}).String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exec(); got != tt.want {
				t.Errorf("got = {%v} \n want = {%v}", got, tt.want)
			}
		})
	}
}

func TestDeleteBuilder_Delete(t *testing.T) {
	tests := []struct {
		name string
		want string
		exec func() string
	}{
		{
			"delete1",
			"DELETE FROM Stock.Product WHERE ProductID=20",
			NewDeleteBuilder().Delete("Stock.Product").Where("ProductID=20").String,
		},
		{
			"delete2",
			"DELETE FROM Stock.ProductPrice WHERE ProductID>2 AND PacketUnitPrice>=2000 OR CartonUnitPrice>40000",
			NewDeleteBuilder().Delete("Stock.ProductPrice").WhereWithMap(map[int]interface{}{
				0: "ProductID>2 AND",
				1: "PacketUnitPrice>=2000 OR",
				2: "CartonUnitPrice>40000",
			}).String,
		},
		{
			"delete3",
			"DELETE FROM Sales.OrderDetail WHERE OrderID>100 OR TotalAmountDue>90000 AND DueDate=10/11/2020",
			NewDeleteBuilder().Delete("Sales.OrderDetail").Where("OrderID>100").Or("TotalAmountDue>90000").And("DueDate=10/11/2020").String,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.exec(); got != tt.want {
				t.Errorf("got = {%v} \n want = {%v}", got, tt.want)
			}
		})
	}
}

func BenchmarkUpdateBuilder_Update(b *testing.B) {

}

func BenchmarkSelectBuilder_Select(b *testing.B) {
	tables := []string{
		"Person.Address",
		"Person.Contact",
		"Sales.Order",
		"Sales.OrderDetails",
		"Sales.Store",
	}

	fields := [][]string{
		{"OrderID", "DueDate", "TotalAmountDue", "PaymentDate", "PaymentMethodID", "DateModified"},
		{"ContactID", "Title", "FirstName", "LastName", "PhoneNumber", "OrderDate"},
		{"OrderID", "Title", "DueDate", "PhoneNumber", "PaymentDate", "StoreID"},
		{"ContactID", "DueDate", "DateAdded", "Quantity", "UnitID", "Unit"},
		{"SupplierID", "ContactID", "DateAdded", "DateModified", "TotalAmount", "SupplyDate"},
	}

	type m = map[int]interface{}
	where := []m{
		{
			0: "OrderID>2",
			1: "StoreID<=100",
			2: "DueDate!=11/02/2020",
			3: "FirstName='Gary'",
		},
		{
			0: "ContactID>222",
			1: "Title='Mrs'",
			2: "OrderDate>=01/11/2003",
			3: "TotalAmount>=21213",
		},
		{
			0: "ContactID>222",
			1: "TotalAmountDue=113333031",
			2: "LastName='Homie'",
			3: "TotalAmount>=100091039",
		},
		{
			0: "DetailID=522",
			1: "DueDate=11/02/1009",
			2: "OrderDate>=01/11/2003",
			3: "TotalAmount>=100091039",
		},
		{
			0: "ContactID>222",
			1: "DueDate=11/33/33031",
			2: "OrderDate>=01/11/2003",
			3: "TotalAmount>=100091039",
		},
	}

	and := []string{
		"ContactID>222",
		"OrderDate>=01/11/2003",
		"TotalAmount>=100091039",
		"StoreID<=100",
		"DetailID=522",
	}

	var result *SelectBuilder
	for i := 0; i < b.N; i++ {
		r1 := rand.Intn(5)
		r2 := rand.Intn(5)
		result = NewSelectBuilder().Select(fields[r1]...).From(tables[r1]).WhereWithMap(where[r1]).
			And(and[r2]).Or(and[r2]).Distinct(fields[r2]...).Asc().Desc().GroupBy(and[r1]).SelectAll(tables[r2]).
			WhereFieldIn(fields[r1][r2], toInterface(fields[r2]...)...)
	}
	result.Clear()
}

func BenchmarkJoinBuilder_Join(b *testing.B) {

}

func BenchmarkInsertBuilder_Insert(b *testing.B) {

}

func BenchmarkDeleteBuilder_Delete(b *testing.B) {

}
