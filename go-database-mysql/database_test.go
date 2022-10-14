package go_database_mysql

import (
	"database/sql"
	"testing"
	_ "github.com/go-sql-driver/mysql"
	// "fmt"
)

func TestOpenConnection(t *testing.T)  {
	// fmt.Printf("Helloo world")
	db, err := sql.Open("mysql","wslubuntu:wslubuntu@tcp(localhost:3306)/go_test_user")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}