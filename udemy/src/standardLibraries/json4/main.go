package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// アンマーシャルのカスタム

type A struct{}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       A         `json:"A"`
}

// 構造体⇒JSONへの変換時に値をカスタムする
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr" + u.Name,
	})
	return v, err
}

// JSON⇒構造体への変換時に値をカスタムする
func (u *User) UnmarshalJSON(b []byte) error {
	// 仮のユーザ型としてUser2型を作成する
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

func main() {

	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "test@test.com"
	u.Created = time.Now()

	// Marshal JSONに変換 ⇒ 返り値：byteのスライス
	bs, err := json.Marshal(u) // この際にMarshalJSON()メソッドが暗黙的に呼び出されている。
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
