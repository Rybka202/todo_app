package service

import(
	"todo_app"
	"todo_app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func newTodoListService(repo repository.TodoList) *TodoListService{
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list todo.TodoList) (int, error){
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error){
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetById(userID, listID int) (todo.TodoList, error){
	return s.repo.GetById(userID, listID)
}

func (s *TodoListService) Delete(userID, listID int) error{
	return s.repo.Delete(userID, listID)
}

func (s *TodoListService) Update(userID, listID int, input todo.UpdateListInput) error{
	if err := input.Validate(); err != nil{
		return err
	}
	return s.repo.Update(userID, listID, input)
}