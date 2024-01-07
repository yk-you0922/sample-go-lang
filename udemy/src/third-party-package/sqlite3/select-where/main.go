package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	Name string
	Age  int
}

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "../example.sql")

	defer Db.Close()

	// 特定のデータを取得
	cmd := "SELECT * FROM persons WHERE age = ?"

	row := Db.QueryRow(cmd, 20)
	var p Person
	err := row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)
}
