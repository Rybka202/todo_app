package repository

import (
	"fmt"
	"strings"
	"todo_app"

	"github.com/jmoiron/sqlx"
) 

type TodoItemPostgres struct {
	db *sqlx.DB
}

func newTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres{
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listID int, item todo.TodoItem) (int, error){
	tx, err := r.db.Begin()
	if err != nil{
		return 0, err
	}

	var itemID int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)

	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	if err = row.Scan(&itemID); err != nil{
		tx.Rollback()
		return 0, err
	}

	createListsItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	if _, err = tx.Exec(createListsItemsQuery, listID, itemID); err != nil{
		tx.Rollback()
		return 0, err
	}

	return itemID, tx.Commit()
}

func (r *TodoItemPostgres) GetAll(userID, listID int) ([]todo.TodoItem, error){
	var items []todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on ti.id = li.item_id 
						INNER JOIN %s ul on li.list_id = ul.list_id WHERE ul.user_id = $1 AND li.list_id = $2`, 
						todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Select(&items, query, userID, listID); err != nil{
		return []todo.TodoItem{}, fmt.Errorf("error select all items: %v", err)
	}
	return items, nil
}

func (r *TodoItemPostgres) GetById(userID, itemID int) (todo.TodoItem, error){
	var item todo.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti INNER JOIN %s li on ti.id = li.item_id 
						INNER JOIN %s ul on li.list_id = ul.list_id WHERE ul.user_id = $1 AND ti.id = $2`, 
						todoItemsTable, listsItemsTable, usersListsTable)
	if err := r.db.Get(&item, query, userID, itemID); err != nil{
		return todo.TodoItem{}, fmt.Errorf("error select item by id: %v", err)
	}
	return item, nil
}

func (r *TodoItemPostgres) Delete(userID, itemID int) error{
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul WHERE ti.id=li.item_id AND li.list_id=ul.list_id 
						AND ul.user_id=$1 AND ti.id=$2`, todoItemsTable, listsItemsTable, usersListsTable)
	
	if _, err := r.db.Exec(query, userID, itemID); err != nil{
		return fmt.Errorf("error delete item: %v", err)
	}

	return nil
}

func (r *TodoItemPostgres) Update(userID, itemID int, input todo.UpdateItemInput) error{
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

	if input.Done != nil{
		setValues = append(setValues, fmt.Sprintf("done=$%d", argID))
		args = append(args, *input.Done)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul WHERE ti.id=li.item_id AND li.list_id=ul.list_id 
						AND ti.id=$%d AND ul.user_id=$%d`,
						todoItemsTable, setQuery, listsItemsTable, usersListsTable, argID, argID+1)
	
	args = append(args, itemID, userID)
	if _, err := r.db.Exec(query, args...); err != nil{
		return err
	}

	return nil
}