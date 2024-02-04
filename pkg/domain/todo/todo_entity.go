package todo

import "errors"

type TodoEntity struct {
	todoId string
	userId string
	todo   string
}

func NewTodoEntity(todoId string, userId string, todo string) (*TodoEntity, error) {
	if todoId == "" || userId == "" || todo == "" {
		return nil, errors.New("必須項目が設定されていません")
	}
	return &TodoEntity{
		todoId: todoId,
		userId: userId,
		todo:   todo,
	}, nil
}

// getter
func (e *TodoEntity) GetTodoId() string {
	return e.todoId
}
func (e *TodoEntity) GetUserId() string {
	return e.userId
}
func (e *TodoEntity) GetTodo() string {
	return e.todo
}
