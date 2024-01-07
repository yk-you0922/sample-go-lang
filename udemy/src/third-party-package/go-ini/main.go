package main

import (
	"fmt"

	"gopkg.in/go-ini/ini.v1"
)

// iniの構造体
type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

// グローバル変数として定義
var Config ConfigList

func init() {
	// 読み込み対象ファイルの読み込み
	cfg, _ := ini.Load("config.ini")

	// 各セクションからキーを元に取得し、グローバル変数にセット
	// MustInt/MustString:iniファイルに設定値が存在しない場合の初期値を指定
	// Stringはiniファイルに設定値が存在しない場合、空文字を返却する設定（文字列型の初期値として）
	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustInt(8080),
		DbName:    cfg.Section("db").Key("name").MustString("example.com"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("Port = %v\n", Config.Port)
	fmt.Printf("Dbname = %v\n", Config.DbName)
	fmt.Printf("SQLDriver = %v\n", Config.SQLDriver)
}
