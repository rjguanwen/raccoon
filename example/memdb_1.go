// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  zhengbin  2020/9/1 11:43
// @Update  姓名（需要改）  2020/9/1 11:43
package main

import (
	"fmt"
	"github.com/hashicorp/go-memdb"
)

// Create a sample struct
type Person struct {
	Email string
	Name  string
	Age   int
}

func main() {

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"person": &memdb.TableSchema{
				Name: "person",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
				},
			},
		},
	}

	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

	// Create a write transaction
	txn := db.Txn(true)

	// Insert a new person
	p := &Person{"joe@aol.com", "Joe", 30}
	if err := txn.Insert("person", p); err != nil {
		panic(err)
	}

	// Commit the transaction
	txn.Commit()

	// Create read-only transaction
	txn = db.Txn(false)
	defer txn.Abort()

	// Lookup by email
	raw, err := txn.First("person", "id", "joe@aol.com")
	if err != nil {
		panic(err)
	}

	// Say hi!
	fmt.Printf("Hello %s!", raw.(*Person).Name)
}
