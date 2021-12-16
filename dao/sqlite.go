package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

//打开数据库，如果不存在，则创建
var DBSqlite *sql.DB

// 初始化，如果不存在则创建表
func InitSqlite() (err error ){
	DBSqlite, err = sql.Open("sqlite3", "./bubble.db")
	if err != nil {
		return
	}
	fmt.Println("初始化成功")

	sql_table := `
    CREATE TABLE IF NOT EXISTS todoinfo(
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title VARCHAR(64) NULL,
        status INTEGER
    );
    `

	_, err = DBSqlite.Exec(sql_table)
	return err
}

func CloseSqlite() {
	DBSqlite.Close()
}
