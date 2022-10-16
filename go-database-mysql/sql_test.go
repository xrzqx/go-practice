package go_database_mysql

import (
	"context"
	"testing"
	"fmt"
	"sync"
	"time"
	"database/sql"
)

func RequestExecAsyc(group *sync.WaitGroup){
	defer group.Done()
	group.Add(1)
	db:= GetConnection()
	defer db.Close()

	ctx := context.Background()
	sql_script := "INSERT INTO User(Name) VALUES('Test')"
	_, err := db.ExecContext(ctx, sql_script)
	if err != nil{
		panic(err)
	}
	time.Sleep(1 * time.Millisecond)
	
}

func RequestPrepareStatementAsyc(group *sync.WaitGroup, db *sql.DB){
	defer group.Done()
	group.Add(1)

	ctx := context.Background()
	sql_script := "INSERT INTO User(Name) VALUES('Test')"
	_, err := db.ExecContext(ctx, sql_script)
	if err != nil{
		panic(err)
	}
	time.Sleep(1 * time.Millisecond)
	
}

// func TestExecSql(t *testing.T)  {
// 	db:= GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	sql_script := "INSERT INTO User(Name) VALUES('Test')"
// 	_, err := db.ExecContext(ctx, sql_script)
// 	if err != nil{
// 		panic(err)
// 	}

// 	fmt.Println("Inserted new User")

// }

func TestNumExecSql(t *testing.T)  {
	group := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		go RequestExecAsyc(group)
	}
	group.Wait()
	fmt.Println("Exec Async Done")

}

// func TestPrepareStatementSql(t *testing.T)  {
// 	db:= GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	sql_script := "INSERT INTO User(Name) VALUES(?)"
// 	// _, err := db.ExecContext(ctx, sql_script)
// 	statement, err := db.PrepareContext(ctx,sql_script)
// 	if err != nil{
// 		panic(err)
// 	}
// 	defer statement.Close()
// 	// fmt.Println("Inserted new User")

// 	for i := 0; i < 200; i++ {
// 		Name := "Test"
// 		_, err := statement.ExecContext(ctx, Name)
// 		if err != nil{
// 			panic(err)
// 		}
// 	}
// 	fmt.Println("Prepare Statement Done")

// }

func TestNumPrepareStatementSql(t *testing.T)  {
	db:= GetConnection()
	defer db.Close()
	group := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		go RequestPrepareStatementAsyc(group,db)
	}
	group.Wait()
	fmt.Println("Prepare Statement Async Done")

}



// func TestQueryCtx(t *testing.T)  {
// 	db:= GetConnection()
// 	defer db.Close()

// 	ctx := context.Background()

// 	// sql_script := "SELECT idUser, Name FROM User"
// 	sql_script := "SELECT * FROM User"
// 	rows, err := db.QueryContext(ctx, sql_script)
	
// 	if err != nil{
// 		panic(err)
// 	}

// 	defer rows.Close()

// 	for rows.Next(){
// 		var idUser, Name string
// 		err = rows.Scan(&idUser, &Name)
// 		if err != nil{
// 			panic(err)
// 		}
// 		fmt.Println("id: ", idUser)
// 		fmt.Println("Name: ", Name)
// 	}
	

// }