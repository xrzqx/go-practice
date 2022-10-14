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

func TestQueryCtx(t *testing.T)  {
	db:= GetConnection()
	defer db.Close()

	ctx := context.Background()

	// sql_script := "SELECT idUser, Name FROM User"
	sql_script := "SELECT * FROM User"
	rows, err := db.QueryContext(ctx, sql_script)
	
	if err != nil{
		panic(err)
	}

	defer rows.Close()

	for rows.Next(){
		var idUser, Name string
		err = rows.Scan(&idUser, &Name)
		if err != nil{
			panic(err)
		}
		fmt.Println("id: ", idUser)
		fmt.Println("Name: ", Name)
	}
	

}