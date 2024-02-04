package todo

import (
	"fmt"

	domain "github.com/satoshi-tahara-st/todo_backend/pkg/domain/todo"
	"gorm.io/gorm"
)

type Todo struct {
	TodoId       string
	UserId       string
	Todo         string
	DeletedTime  string
	RegisterTime string
	UpdateTime   string
}

type todoRepositoryImpl struct {
	refDb *gorm.DB
}

// ITodoRepositoryの実装クラス
// データベースコネクションを外側から渡します
func NewTodoRepositoryImpl(refDb *gorm.DB) domain.ITodoRepository {
	return &todoRepositoryImpl{refDb: refDb}
}

func (r *todoRepositoryImpl) Fetch(todoId string) ([]domain.TodoEntity, error) {
	var todos []Todo
	var records *gorm.DB
	records = r.refDb.Table("todo_t").Select("todo_t.*").Scan(&todos)
	if records.Error != nil {
		return nil, fmt.Errorf("Fetch: レコードの取得に失敗:\n %w", records.Error)
	}
	return toEntities(todos)
}

func toEntities(todos []Todo) ([]domain.TodoEntity, error) {
	results := []domain.TodoEntity{}
	for _, todo := range todos {
		result, err := domain.NewTodoEntity(todo.TodoId, todo.UserId, todo.Todo)
		if err != nil {
			return nil, fmt.Errorf("Fetch: todoエンティティーの変換に失敗:\n %w", err)
		}
		results = append(results, *result)
	}
	return results, nil
}
