package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	openDB()
	dbinquire()
	/* go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)
		<-c
		closeDB()
	}() */
}
