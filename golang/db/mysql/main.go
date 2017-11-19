package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

/*
+--------------+---------+------+-----+---------+-------+
| Field        | Type    | Null | Key |   UNQ   | Extra |
+--------------+---------+------+-----+---------+-------+
| id       	   | int(11) | NO   | PRI | YES     |       |
| username	   | string  | NO   |     | YES     |       |
| password	   | string  | NO   |     | NO      |       |
+--------------+---------+------+-----+---------+-------+
*/
func main() {
	db, err := sql.Open("mysql", "username:password@/labor")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO user VALUES(2, 'davinz', 'password' )")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec()
	if err != nil {
		panic(err.Error())
	}

}
