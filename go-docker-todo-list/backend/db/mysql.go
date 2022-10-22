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
	query := `INSERT INTO 
		todo(note, done) 
		VALUES(?,?);
		`
	statement, err := p.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer statement.Close()
	res, err := statement.Exec(todo.Note, todo.Done)
	if err != nil {
		return -1, err
	}
	lid, err := res.LastInsertId()
	return int(lid), nil
}

func (p *Mysql) Update(todo *schema.Todo) error {
	query := `
		UPDATE todo
		SET note = ?, done = ?
		WHERE id = ?;
	`

	rows, err := p.DB.Query(query, todo.Note, todo.Done, todo.ID)
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
		WHERE id = ?;
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
	connStr, err := loadMysqlConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", connStr)
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

