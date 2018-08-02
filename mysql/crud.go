package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB
var insertId int64

func main() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4")
	if err != nil {
		fmt.Errorf("connectErr is %s", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Errorf("connectErr is %s", err)
	}

	insert()

	findOne()
	findAll()

	update()

	del()

}

func insert() {
	stmt, err := db.Prepare("insert into together_travel_log(uid,ttid,info,ctime) value (?,?,?,?)")
	if err != nil {
		fmt.Errorf("insertErr is %s", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(123, 123, "test go mysql", "2018-01-01 12:00:00")
	if err != nil {
		fmt.Errorf("insertErr is %s", err)
	}
	insertId, err = result.LastInsertId()
	if err != nil {
		fmt.Errorf("insertErr is %s", err)
	}
	fmt.Println(insertId)
}

func findOne() {
	stmt, err := db.Prepare("select id,uid,ttid,info,ctime from together_travel_log where id = ?")
	if err != nil {
		fmt.Errorf("selectErr is %s", err)
	}
	defer stmt.Close()

	var id, uid, ttid int64
	var info, ctime string

	err = stmt.QueryRow(insertId).Scan(&id, &uid, &ttid, &info, &ctime) // WHERE id = insertId
	if err != nil {
		fmt.Errorf("selectErr is %s", err)
	}

	fmt.Printf("id is %d, uid is %d, ttid is %d, info is %s, ctime is %s \n", id, uid, ttid, info, ctime)
}

func findAll() {
	rows, err := db.Query("select id,uid,ttid,info,ctime from together_travel_log")
	if err != nil {
		fmt.Errorf("selectErr is %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, uid, ttid int64
		var info, ctime string

		err = rows.Scan(&id, &uid, &ttid, &info, &ctime) // WHERE id = insertId
		if err != nil {
			fmt.Errorf("selectErr is %s", err)
		}
		fmt.Printf("id is %d, uid is %d, ttid is %d, info is %s, ctime is %s \n", id, uid, ttid, info, ctime)
	}
}

func update() {
	stmt, err := db.Prepare("update  together_travel_log set ttid = ? where id = ?")
	if err != nil {
		fmt.Errorf("updateErr is %s", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(3123123, insertId)

	if err != nil {
		fmt.Errorf("updateErr is %s", err)
	}
	rosAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Errorf("updateErr is %s", err)
	}
	fmt.Println(rosAffected)
}

func del() {
	stmt, err := db.Prepare("delete from together_travel_log where id = ?")
	if err != nil {
		fmt.Errorf("delErr is %s", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(insertId)
	if err != nil {
		fmt.Errorf("delErr is %s", err)
	}
	rosAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Errorf("delErr is %s", err)
	}
	fmt.Println(rosAffected)
}
