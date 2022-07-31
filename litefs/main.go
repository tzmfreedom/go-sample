package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	return
	db, err := sql.Open("sqlite3", os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE foo (id integer);")
	if err != nil {
		log.Printf("%q\n", err)
		return
	}
}
