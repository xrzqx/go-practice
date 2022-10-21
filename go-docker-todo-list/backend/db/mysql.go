package db

import (
	"database/sql"
	"fmt"
	"github.com/xrzqx/go-practice/go-docker-todo-list/schema"
	"os"
)

type Mysql struct {
	DB *sql.DB
}

func (p *Mysql) GetAll() ([]schema.Todo, error) {
	query := `
		SELECT *
		FROM todo
		ORDER BY id;
	`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	todoList := []schema.Todo{}
	for rows.Next() {
		var t schema.Todo
		if err := rows.Scan(&t.ID, &t.Note, &t.Done); err != nil {
			return nil, err
		}
		todoList = append(todoList, t)
	}
	return todoList, nil
}

func (p *Mysql) Insert(todo *schema.Todo) (int, error) {
	query := `
		INSERT INTO todo (note, done)
		VALUES($1, $2);
	`
	rows, err := p.DB.Query(query, todo.Note, convertBoolToBit(todo.Done))
	if err != nil {
		return -1, err
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (p *Mysql) Update(todo *schema.Todo) error {
	query := `
		UPDATE todo
		SET note = $2, done = $3
		WHERE id = $1;
	`

	rows, err := p.DB.Query(query, todo.ID, todo.Note, convertBoolToBit(todo.Done))
	if err != nil {
		return err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return err
		}
	}
	return nil
}

func (p *Mysql) Delete(id int) error {
	query := `
		DELETE FROM todo
		WHERE id = $1;
	`

	if _, err := p.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func (p *Mysql) Close() {
	p.DB.Close()
}

func ConnectMysql() (*Mysql, error) {
	//connStr, err := loadMysqlConfig()
	//if err != nil {
	//	return nil, err
	//}

	//db, err := sql.Open("mysql", connStr)
	db, err := sql.Open("mysql", "root:rahasia@tcp(localhost:5432)/demo_xrzqx_todo")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &Mysql{db}, nil
}

func loadMysqlConfig() (string, error) {
	if os.Getenv("DB_HOST") == "" {
		return "", fmt.Errorf("Environment variable DB_HOST must be set")
	}
	if os.Getenv("DB_PORT") == "" {
		return "", fmt.Errorf("Environment variable DB_PORT must be set")
	}
	if os.Getenv("DB_USER") == "" {
		return "", fmt.Errorf("Environment variable DB_USER must be set")
	}
	if os.Getenv("DB_DATABASE") == "" {
		return "", fmt.Errorf("Environment variable DB_DATABASE must be set")
	}
	if os.Getenv("DB_PASSWORD") == "" {
		return "", fmt.Errorf("Environment variable DB_PASSWORD must be set")
	}
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
	return connStr, nil
}

func convertBoolToBit(val bool) int {
	if val {
		return 1
	}
	return 0
}
