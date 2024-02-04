package di

import (
	"github.com/satoshi-tahara-st/todo_backend/pkg/application"
	"github.com/satoshi-tahara-st/todo_backend/pkg/handlers"
	db "github.com/satoshi-tahara-st/todo_backend/pkg/infra/db/todo"
	service "github.com/satoshi-tahara-st/todo_backend/pkg/service/todo"
	"gorm.io/gorm"
)

func TodoHandler(dbs map[string]map[string]*gorm.DB) handlers.TodoHandler {
	refDb := dbs[application.SigninigKey][application.ConnectionRef]
	todoRepository := db.NewTodoRepositoryImpl(refDb)
	todoService := service.NewTodoService(todoRepository)

	return handlers.NewTodoHandler(todoService)
}
