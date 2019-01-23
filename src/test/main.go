package main

import (
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

type S struct{}
type T struct {
	S //
}

func (S) sVal()  {}
func (*S) sPtr() {}
func (T) tVal()  {}
func (*T) tPtr() {}
func methodSet(a interface{}) { //
	t := reflect.TypeOf(a)
	fmt.Println(t.NumMethod())
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
	fmt.Println(111)
}
func main() {
	var t T
	methodSet(t) // T
	println("----------")
	methodSet(&t) // *T
}
