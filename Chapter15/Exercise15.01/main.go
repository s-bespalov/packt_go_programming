package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var property string
	db, err := sql.Open("postgres", "user=postgres password=#1990ak host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}
	defer db.Close()
	TableCreate := `
CREATE TABLE Number
(
	Number integer NOT NULL,
	Property text COLLATE pg_catalog."default" NOT NULL
)
WITH (
	OIDS = FALSE
)
TABLESPACE pg_default;
ALTER TABLE Number OWNER to postgres;
`
	_, err = db.Exec(TableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table \"Number\" was created")
	}
	insert, err := db.Prepare("insert into Number values ($1, $2)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			property = "Even"
		} else {
			property = "Odd"
		}
		_, err := insert.Exec(i, property)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("The number %d is %s\n", i, property)
		}
	}
	fmt.Println("The numbers are ready")
}
