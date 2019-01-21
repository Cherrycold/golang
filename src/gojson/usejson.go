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
	//d test1
}

func Makejson() {
	data := test2{"jack", 13, "fish"}
	user, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(user)
	fmt.Println(user)
}
