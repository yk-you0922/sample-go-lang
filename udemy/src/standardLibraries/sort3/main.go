package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// ここでカスタムする。
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

func main() {
	m := map[string]int{"S": 1, "J": 4, "A": 3, "N": 3}

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	// Sort ⇒ Len(), Swap(), Less()の3つの関数を実装することで利用可能となる
	// Less()にカスタムソートを実装する。
	// ※実装しなければコンパイルエラー
	sort.Sort(lt)
	fmt.Println(lt)

	// Reverse
	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)
}
