package models

import (
	"bubble/dao"
	"database/sql"
	"fmt"
)


type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status int  `json:"status"`
}
/*
	Todo这个Model的增删改查操作都放在这里
*/

// CreateATodo 创建todo
func CreateATodoSqlite(todo *Todo) (err error){

	newId := 0
	rows, err := dao.DBSqlite.Query("SELECT max(id) FROM todoinfo")
	for rows.Next() {
		err = rows.Scan(&newId)
	}

	stmt, err := dao.DBSqlite.Prepare("INSERT INTO todoinfo(id, title, status) values(?,?,?)")
	if err != nil {
		return err
	}

	fmt.Println("插入数据：", todo, "id: ", newId)
	_, err = stmt.Exec(newId + 1, todo.Title, todo.Status)
	if err != nil {
		fmt.Println("插入数据出错：", err)
		return err
	}

	return nil
}

func GetAllTodoSqlite() (todoList []*Todo, err error){
	rows, err := dao.DBSqlite.Query("SELECT id, title, status FROM todoinfo")
	if err != nil {
		return
	}

	var (
		gid int
		gtitle string
		gstatus int
	)
	for rows.Next() {
		err = rows.Scan(&gid, &gtitle, &gstatus)
		if err != nil {
			return
		}
		newTodo := new(Todo)
		newTodo.ID = gid
		newTodo.Title = gtitle
		newTodo.Status = gstatus
		todoList = append(todoList, newTodo)
	}

	rows.Close()
	return
}

func GetATodoSqlite(id string)(todo *Todo, err error){
	todo = new(Todo)
	var row *sql.Rows
	row, err = dao.DBSqlite.Query("SELECT id, title, status FROM todoinfo WHERE id = ?", id)
	if err != nil {
		return
	}
	if row != nil {
		row.Scan(&todo.ID, &todo.Title, &todo.Status)
	}
	return
}

func UpdateATodoSqlite(id string)(err error){
	fmt.Println("sqlite更新状态：", id)
	var stmt *sql.Stmt
	stmt, err = dao.DBSqlite.Prepare("update todoinfo set status = (status+1) % 2 where id = ?")
	if err != nil {
		return
	}
	_, err = stmt.Exec(id)
	return
}

func DeleteATodoSqlite(id string)(err error){
	var stmt *sql.Stmt
	stmt, err = dao.DBSqlite.Prepare("DELETE FROM todoinfo where id = ?")
	if err != nil {
		return
	}
	_, err = stmt.Exec(id)
	return
}

