package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "myTaskManager"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}
func (r *TodoItemPostgres) Create(lisId int, item todo.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemquery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemquery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		return 0, err
	}
	createListItemQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) values ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemQuery, lisId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return itemId, tx.Commit()
}
