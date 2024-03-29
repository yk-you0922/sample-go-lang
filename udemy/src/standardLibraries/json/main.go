package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A,omitempty"`
}

func main() {

	u := new(User)
	u.ID = 0
	u.Name = ""
	u.Email = "test@test.com"
	u.Created = time.Now()

	// Marshal JSONに変換 ⇒ 返り値：byteのスライス
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
