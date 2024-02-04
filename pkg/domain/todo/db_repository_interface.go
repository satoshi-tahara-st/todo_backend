package todo

type ITodoRepository interface {
	Fetch(todoId string) ([]TodoEntity, error)
}
