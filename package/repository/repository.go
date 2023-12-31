package repository

import (
	"github.com/jmoiron/sqlx"
	"myTaskManager/models"
	"myTaskManager/package/repository/postgres"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Delete(userId, listId int) error
	Update(userId int, listId int, input models.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item models.TodoItem) (int, error)
	GetAll(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, listId int) (models.TodoItem, error)
	Delete(userId, ItemId int) error
	Update(userId, itemId int, input models.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepositoty(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		TodoList:      postgres.NewTodoListPostgres(db),
		TodoItem:      postgres.NewTodoItemPostgres(db),
	}
}
