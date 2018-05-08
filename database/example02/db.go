package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/?charset=utf8")
	checkErr(err)

	defer db.Close()

	db.Exec("drop database if exists tmpdb")
	db.Exec("create database tmpdb")
	db.Exec("create table tmpdb.users(id int not null, name varchar(20), age int, PRIMARY KEY (id))")
	db.Exec("insert into tmpdb.users values(?, ?, ?)", 1, "Tom", 20)
	db.Exec("insert into tmpdb.users values(?, ?, ?)", 2, "Tom", 10)
	db.Exec("insert into tmpdb.users values(?, ?, ?)", 3, "Jerry", 20)
	db.Exec("insert into tmpdb.users values(?, ?, ?)", 3, "Gopher", 30)

	// Query
	fmt.Println("--Query--")
	age := 20
	rows1, err1 := db.Query("select name from tmpdb.users where age = ?", age) // Query all matched rows

	checkErr(err1)
	defer rows1.Close()

	for rows1.Next() {
		var name string
		if err := rows1.Scan(&name); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s is %d\n", name, age)
	}

	// QueryRow
	fmt.Println("--QueryRow--")
	name := "Tom"
	row := db.QueryRow("SELECT age FROM tmpdb.users WHERE name = ?", name) // Query 1 row
	err2 := row.Scan(&age)
	checkErr(err2)
	fmt.Printf("%s is %d\n", name, age)

	// Prepare statements
	fmt.Println("--Prepare Statements--")
	stmt, err3 := db.Prepare("SELECT name FROM tmpdb.users WHERE age = ?")

	checkErr(err3)
	defer stmt.Close()

	rows4, err4 := stmt.Query(age)
	checkErr(err4)
	defer rows4.Close()

	for rows4.Next() {
		var name string
		if err := rows4.Scan(&name); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s is %d\n", name, age)
	}

	// Trasaction
	fmt.Println("--Transaction--")
	tx, err5 := db.Begin()
	checkErr(err5)
	tx.Exec("update tmpdb.users set name='Micky' where age = ?", 30)
	tx.Exec("delete from tmpdb.users where age = ?", 10)
	tx.Commit()

	db.Exec("drop table tmpdb.users")
	db.Exec("drop database tmpdb")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
