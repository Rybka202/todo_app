package service

import (
	"todo_app"
	"todo_app/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func newTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService{
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userID, listID int, item todo.TodoItem) (int, error){
	if _, err := s.listRepo.GetById(userID, listID); err != nil{
		return 0, err
	}

	return s.repo.Create(listID, item)
}

func (s *TodoItemService) GetAll(userID, listID int) ([]todo.TodoItem, error){
	return s.repo.GetAll(userID, listID)
}

func (s *TodoItemService) GetById(userID, itemID int) (todo.TodoItem, error){
	return s.repo.GetById(userID, itemID)
}

func (s *TodoItemService) Delete(userID, itemID int) error{
	return s.repo.Delete(userID, itemID)
}

func (s *TodoItemService) Update(userID, itemID int, input todo.UpdateItemInput) error{
	if err := input.Validate(); err != nil{
		return err
	}
	return s.repo.Update(userID, itemID, input)
}