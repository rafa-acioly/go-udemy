package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	res, err := stmt.Exec("rafael", 1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.RowsAffected())
}
