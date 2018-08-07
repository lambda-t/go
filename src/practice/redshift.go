package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
)


func MakeRedshfitConnection(username, password, host, port, dbName string) (*sql.DB, error) {

	url := fmt.Sprintf("sslmode=require user=%v password=%v host=%v port=%v dbname=%v",
		username,
		password,
		host,
		port,
		dbName)

	var err error
	var db *sql.DB
	if db, err = sql.Open("postgres", url); err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("redshift ping error : (%v)", err)
	}
	return db, nil
}

func main() {
	res,err := MakeRedshfitConnection("root","ShiftRed123.",
		"ganalytic.cvqb8lx7aald.us-west-2.redshift.amazonaws.com","5439","ganalytic")
	if err != nil {
		log.Fatal(err)
	}

	rows err= res.Query("select * from explorer limit 10")
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}

