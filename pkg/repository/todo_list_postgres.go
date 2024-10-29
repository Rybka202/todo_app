package repository

import (
	"fmt"
	"strings"
	"todo_app"

	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func newTodoListPostgres(db *sqlx.DB) *TodoListPostgres{
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userID int, list todo.TodoList) (int, error){
	tx, err := r.db.Begin()
	if err != nil{
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil{
		tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	if _, err := tx.Exec(createUsersListsQuery, userID, id); err != nil{
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userID int) ([]todo.TodoList, error){
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT %[1]s.id, %[1]s.title, %[1]s.description FROM %[1]s INNER JOIN %[2]s on %[1]s.id = %[2]s.list_id WHERE %[2]s.user_id = $1", 
		todoListsTable, usersListsTable)
	if err := r.db.Select(&lists, query, userID); err != nil{
		return []todo.TodoList{}, fmt.Errorf("error select all lists: %v", err)
	}
	return lists, nil
}

func (r *TodoListPostgres) GetById(userID int, listID int) (todo.TodoList, error){
	var list todo.TodoList
	query := fmt.Sprintf(`SELECT %[1]s.id, %[1]s.title, %[1]s.description FROM %[1]s INNER JOIN %[2]s on %[1]s.id = %[2]s.list_id
		WHERE %[2]s.user_id = $1 AND %[1]s.id = $2`, todoListsTable, usersListsTable)
	if err := r.db.Get(&list, query, userID, listID); err != nil{
		return todo.TodoList{}, fmt.Errorf("error select list by id: %v", err)
	}
	return list, nil
}

func (r *TodoListPostgres) Delete(userID int, listID int) error{
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id=ul.list_id AND ul.user_id=$1 AND tl.id=$2",
						todoListsTable, usersListsTable)
	if _, err := r.db.Exec(query, userID, listID); err != nil{
		return err
	}
	return nil
}

func (r *TodoListPostgres) Update(userID, listID int, input todo.UpdateListInput) error{
	setValues := make([]string, 0)
	args := make([]any, 0)
	argID := 1

	if input.Title != nil{
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}

	if input.Description != nil{
		setValues = append(setValues, fmt.Sprintf("description=$%d", argID))
		args = append(args, *input.Description)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id=ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
						todoListsTable, setQuery, usersListsTable, argID, argID+1)
	
	args = append(args, listID, userID)
	if _, err := r.db.Exec(query, args...); err != nil{
		return err
	}

	return nil
}

