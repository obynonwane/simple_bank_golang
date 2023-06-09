package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

/* Database connection parameters*/
const (
	dbDriver = "postgres"
 	dbSource = "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable"
)

/* *Queries (is a pointer) is used to access the variable stored at the memory location
where Queries is pointed to, however testQueries is a type of *Queries
*/
var testQueries *Queries

/*This the function that runs before any other test is ran 
 It is used for Initilization e.g setting up and connecting to DB
 */
func TestMain(m *testing.M) {
	conn, err :=sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to db",err)
	}

	/*Creating a new pointer of type *Queries */
	testQueries = New(conn)

	os.Exit(m.Run())
}