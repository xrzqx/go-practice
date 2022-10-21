package handler

import (
	"github.com/xrzqx/go-practice/go-docker-todo-list/db"
	"net/http"
)

func InitRoutes(mysql *db.Mysql) *http.ServeMux {
	todoHandler := &todoHandler{
		mysql:  mysql,
		static: &db.Static{},
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/static", todoHandler.GetStatic)
	mux.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization, cache-control")

		switch r.Method {
		case http.MethodOptions:
			w.Write([]byte("allowed"))
		case http.MethodGet:
			todoHandler.getAllTodo(w, r)
		case http.MethodPost:
			todoHandler.insertTodo(w, r)
		case http.MethodPut:
			todoHandler.updateTodo(w, r)
		case http.MethodDelete:
			todoHandler.deleteTodo(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	})

	return mux
}
