package usedb

import (
	"database/sql"
	"fmt"
	"log"
)

type dbconfig struct {
	user        string
	host        string
	password    string
	port        string
	name        string
	maxpoolsize int
}

var db *sql.DB

func closeDB() {
	db.Close()
}

func dbinquire() {
	rows, err := db.Query("select t1int,t2char from test")
	if err != nil {
		log.Fatal(err)
	}
	var (
		t1 int
		t2 string
	)
	//Next是获取每一行
	for rows.Next() {
		//Scan是获取一行中列
		err = rows.Scan(&t1, &t2)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(t1, t2)
	}
}

func dbinsert() {
	stmt, err := db.Prepare("insert into test(t1int,t2char,t3json) values (?,?,NULL)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(7, "hello world")
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() {
	dbc := new(dbconfig)
	dbc.user = "root"
	dbc.host = "localhost"
	dbc.password = "asd123"
	dbc.port = "3306"
	dbc.name = "test"

	dbstr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		dbc.user,
		dbc.password,
		dbc.host,
		dbc.port,
		dbc.name)
	var err error
	db, err = sql.Open("mysql", dbstr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("open success")
	}

}
