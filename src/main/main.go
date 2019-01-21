package main

import (
	"gojson"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	gojson.Makejson()
	//openDB()
	//dbinquire()
	/* go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		<-c
		closeDB()
	}()*/
}
