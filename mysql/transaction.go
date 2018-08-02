package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		fmt.Errorf("connectErr is %s", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Errorf("connectErr is %s", err)
	}

	// 事务开始
	tx, err := db.Begin()
	if err != nil {
		fmt.Errorf("beginErr is %s", err)
	}

	stmt, err := tx.Prepare("insert into together_travel_log(uid,ttid,info,ctime) value (?,?,?,?)")
	if err != nil {
		fmt.Errorf("sql is %s", err)
	}

	defer stmt.Close()

	result, err := stmt.Exec(123, 123, "test go mysql", "2018-01-01 12:00:00")
	if err != nil {
		fmt.Errorf("insertErr is %s", err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		fmt.Errorf("insertErr is %s", err)
	}
	fmt.Println(insertId)


	//err = tx.Commit()
	//if err != nil {
	//	fmt.Errorf("commitErr is %s", err)
	//}

	err = tx.Rollback()
	if err != nil {
		fmt.Errorf("rollBackErr is %s", err)
	}
}
