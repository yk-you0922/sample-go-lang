package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A"`
}

func main() {

	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@test.com"
	u.Created = time.Now()

	// Marshal JSONに変換 ⇒ 返り値：byteのスライス
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	// -----------------------------------

	fmt.Printf("%T\n", bs) // bsの型を調べる⇒バイトのスライス

	u2 := new(User)

	// Unmarshal JSONをデータに変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u2)
}