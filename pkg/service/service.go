package service

import (
	"todo_app"
	"todo_app/pkg/repository"

)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetById(userID, listID int) (todo.TodoList, error)
	Delete(userID, listID int) error
	Update(userID, listID int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userID, listID int, item todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetById(userID, itemID int) (todo.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input todo.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
		TodoList: newTodoListService(repos.TodoList),
		TodoItem: newTodoItemService(repos.TodoItem, repos.TodoList),
	}
}