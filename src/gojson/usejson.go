package gojson

import (
	"encoding/json"
	"fmt"
	"log"
)

type test1 struct {
	A string
	B string
	C string
}

type test2 struct {
	Name string
	Old  int
	Like string
	D    test1
}

var data test2
var user []byte

func Makejson() {
	t := test1{"hello", ",", "world"}
	data = test2{"jack", 13, "fish", t}
	var err error
	user, err = json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
}

func Expjson() {
	var expuser test2
	err := json.Unmarshal(user, &expuser)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(expuser)
}
