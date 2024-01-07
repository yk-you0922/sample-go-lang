package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func main() {
	Db, _ := sql.Open("sqlite3", "../example.sql")

	defer Db.Close()

	// データの追加
	cmd := "INSERT INTO persons (name, age) VALUES (?, ?)" // (?, ?)とすることでSQLインジェクション対策となる。
	_, err := Db.Exec(cmd, "Taro", 20)
	if err != nil {
		log.Fatalln(err)
	}
}
