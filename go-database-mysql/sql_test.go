package go_database_mysql

import (
	"context"
	"testing"
	"fmt"
)

func TestExecSql(t *testing.T)  {
	db:= GetConnection()
	defer db.Close()

	ctx := context.Background()

	sql_script := "INSERT INTO User(Name) VALUES('Test')"
	_, err := db.ExecContext(ctx, sql_script)
	if err != nil{
		panic(err)
	}

	fmt.Println("Inserted new User")

}