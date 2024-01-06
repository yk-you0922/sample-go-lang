package main

import (
	"fmt"
	"sync"
	"time"
)


var st struct { A, B, C int}

// ミューテックスを保持するパッケージ変数（参照型）
var mutex *sync.Mutex

// 引数の値をstのA、B、Cにそれぞれ代入する。
func UpdateAndPrint(n int) {
	// ロック
	mutex.Lock()
	
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	time.Sleep(time.Millisecond)
	fmt.Println(st)

	// アンロック
	mutex.Unlock()
}

func main() {
	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	// 無限ループ。UpdateAndPrint()が終わる前にmain関数自体が処理を終えるため設定している。
	for {

	}

}
