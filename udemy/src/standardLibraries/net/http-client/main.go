package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	/**
	* Getメソッド
	*/
	res, _ := http.Get("https://example.com")
	
	fmt.Println(res.StatusCode)
	
	fmt.Println(res.Proto)
	
	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Content-Type"])
	
	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body) // bodyはバイトのスライス
	fmt.Println(string(body))

	/**
	* Postメソッド
	*/
	vs := url.Values{}

	vs.Add("id", "1")
	vs.Add("message", "メッセージ")

	res, err := http.PostForm("https://example.com", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

}
