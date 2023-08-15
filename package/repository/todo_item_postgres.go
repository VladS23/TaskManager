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
func (r *TodoItemPostgres) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	var items []todo.TodoItem
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl INER JOIN %s li on li.item_id=tl.id"+
		"INNER JOIN %S ul on ul.list_id=li.list_id"+
		"WHERE li.list_id=$1 AND ul.user_id=$2",
		todoItemsTable, listsItemsTable, userListTable)
	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}
func (r *TodoItemPostgres) GetById(userId, itemId int) (todo.TodoItem, error) {
	var item todo.TodoItem
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description, tl.done FROM %s tl INER JOIN %s li on li.item_id=tl.id"+
		"INNER JOIN %S ul on ul.list_id=li.list_id"+
		"WHERE tl.id=$1 AND ul.user_id=$2",
		todoItemsTable, listsItemsTable, userListTable)
	if err := r.db.Get(&item, query, itemId, userId); err != nil {
		return item, err
	}
	return item, nil
}
func (r *TodoItemPostgres) Delete(userId, ItemId int) error {
	query := fmt.Sprintf("DELETE FROM %s li USING %s li, %s ul WHERE tl.id=li.item_id AND li.list_id=ul.list_id AND ul.user.id=$1 AND tl.id=$2", todoItemsTable, listsItemsTable, userListTable)
	_, err := r.db.Exec(query, userId, ItemId)
	return err
}
