package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xrzqx/go-practice/go-docker-todo-list/db"
	"github.com/xrzqx/go-practice/go-docker-todo-list/handler"
)

func main() {
	var mysql *db.Mysql
	var err error

	mysql, err = db.ConnectMysql()
	if err != nil {
		panic(err)
	}
	if mysql == nil {
		panic("mysql is unreachable")
	}

	mux := handler.InitRoutes(mysql)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
