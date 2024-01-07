# サードパーティパッケージ

## database/sqlite3
使用パッケージ
- database/sql
- github.com/mattn/go-sqlite3

参考コード
- [テーブル作成](../../../udemy/src/third-party-package/sqlite3/create-table/main.go)
- [Insert](../../../udemy/src/third-party-package/sqlite3/insert/main.go)
- [Select](../../../udemy/src/third-party-package/sqlite3/select/main.go)
- [Select-Where](../../../udemy/src/third-party-package/sqlite3/select-where/main.go)
- [Update](../../../udemy/src/third-party-package/sqlite3/update/main.go)
- [Delete](../../../udemy/src/third-party-package/sqlite3/delete/main.go)

## database/psql
1/7 一旦飛ばす

## go-ini
`go get gopkg.in/go-ini/ini.v1` コマンドでインストール

iniファイルの読み込みを行うパッケージ

configの設定をiniファイルから読み取り、プログラム上で利用できるようにする。

config.iniの用意
```ini
[web]
port = 8080

[db]
name = webapp.sql
driver = sqlite3
```

アプリ側のサンプルコード
```go
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
```

## uuid
`go get "github.com/google/uuid"` でインストール

ソフトウェア上でオブジェクトを一意に識別するためのIDを生成するパッケージ

```go
package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuidObj, _ := uuid.NewUUID()
	fmt.Println("  ", uuidObj.String())

	uuidObj2, _ := uuid.NewRandom()
	fmt.Println("  ", uuidObj2.String())
}
```

出力結果
```
   cc25124d-ad28-11ee-9190-00155d3c9a38
   c24e3204-a345-420f-9f59-8f1b98fa0f79
```