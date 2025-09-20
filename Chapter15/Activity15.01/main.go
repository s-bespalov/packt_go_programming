package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=#1990ak host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}
	defer db.Close()
	_, err = db.Exec("DROP TABLE IF EXISTS Users;")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table \"Users\" dropped if existed")
	}
	TableCreate := `
CREATE TABLE Users
(
  ID SERIAL PRIMARY KEY,
  Name VARCHAR(255) NOT NULL,
  Email VARCHAR(255) UNIQUE NOT NULL
)
TABLESPACE pg_default;
ALTER TABLE Users OWNER to postgres;
	`
	_, err = db.Exec(TableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table \"Users\" was created")
	}
	stmt, err := db.Prepare("INSERT INTO Users (Name, Email) VALUES ($1, $2)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	users := []struct {
		name, email string
	}{
		{"Alice", "alice@example.com"},
		{"Bob", "bob@example.com"},
		{"Charlie", "charlie@example.com"},
	}
	for _, user := range users {
		_, err = stmt.Exec(user.name, user.email)
		if err != nil {
			panic(err)
		}
		fmt.Printf("The user with name: %s and email: %s was successfully added\n", user.name, user.email)
	}
	stmt2, err := db.Prepare("UPDATE Users SET Email = $2 WHERE ID = $1")
	if err != nil {
		panic(err)
	}
	defer stmt2.Close()
	results, err := stmt2.Exec(1, "user@packt.com")
	if err != nil {
		panic(err)
	}
	if c, err := results.RowsAffected(); err == nil {
		if c == 1 {
			fmt.Println("User email was successfully updated")
		} else {
			fmt.Println("Strange, affected rows:", c)
		}
	} else {
		panic(err)
	}
	stmt3, err := db.Prepare("DELETE FROM Users WHERE ID = $1")
	if err != nil {
		panic(err)
	}
	defer stmt3.Close()
	_, err = stmt3.Exec(2)
	if err != nil {
		panic(err)
	}
	fmt.Println("User with ID = 2 was successfully deleted")
	stmt4, err := db.Prepare("SELECT * FROM Users")
	if err != nil {
		panic(err)
	}
	defer stmt4.Close()
	usersResult, err := stmt4.Query()
	if err != nil {
		panic(err)
	}
	var name, email string
	var id int
	fmt.Println("Final users table content:")
	for usersResult.Next() {
		err = usersResult.Scan(&id, &name, &email)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\t%s\t%s\n", id, name, email)
	}
}
