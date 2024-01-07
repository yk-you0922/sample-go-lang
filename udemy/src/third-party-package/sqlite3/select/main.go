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
	cmd := "SELECT * FROM persons"

	rows, _ := Db.Query(cmd, 20)
	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Println(err)
		}
		pp = append(pp, p)
	}

	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
