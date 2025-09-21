package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=#1990ak host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}
	defer db.Close()
	err = createMessagesTable()
	if err != nil {
		panic(err)
	}
	err = insertMessages()
	if err != nil {
		panic(err)
	}
	err = printMessages()
	if err != nil {
		panic(err)
	}
	name, err := promptUserName()
	if err != nil {
		panic(err)
	}
	err = printMessageByName(name)
	if err != nil {
		panic(err)
	}
}

func promptUserName() (string, error) {
	fmt.Print("Give me the user's name: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read user name: %w", err)
	}
	name = strings.TrimSpace(name)
	return name, nil
}

func printMessageByName(name string) error {
	query := `
SELECT Users.Name, Users.Email, Messages.Message
FROM Users
LEFT JOIN Messages
ON Users.ID = Messages.UserID
WHERE Users.Name LIKE $1
	`
	results, err := querySQL(query, name)
	if err != nil {
		return fmt.Errorf("failed to query messages by name: %w", err)
	}
	defer results.Close()
	m := struct {
		userName, email, message string
	}{}
	empty := true
	for results.Next() {
		err = results.Scan(&m.userName, &m.email, &m.message)
		if err != nil {
			return fmt.Errorf("failed to scan message by name: %w", err)
		}
		empty = false
		fmt.Println("the user:", m.userName, "with email:", m.email, "sent message:", m.message)
	}
	if empty {
		fmt.Println("The query returned nothing, no such user:", name)
	}
	return nil
}

func printMessages() error {
	query := "SELECT * FROM Messages"
	rsl, err := querySQL(query)
	if err != nil {
		return fmt.Errorf("failed to query messages: %w", err)
	}
	defer rsl.Close()
	var messageID, userID int
	var text string
	for rsl.Next() {
		err := rsl.Scan(&messageID, &text, &userID)
		if err != nil {
			return fmt.Errorf("failed to scan message: %w", err)
		}
		fmt.Printf("ID:%d\tText:%s\tUserID:%d\n", messageID, text, userID)
	}
	return nil
}

func createMessagesTable() error {
	_, err := db.Exec("DROP TABLE IF EXISTS Messages;")
	if err != nil {
		return fmt.Errorf("failed to drop table: %w", err)
	} else {
		fmt.Println("Table \"Messages\" dropped if existed")
	}
	TableCreate := `
CREATE TABLE Messages
(
  ID SERIAL PRIMARY KEY,
  Message VARCHAR(280) NOT NULL,
  UserID Integer NOT NULL
)
TABLESPACE pg_default;
ALTER TABLE Messages OWNER to postgres;
`
	_, err = db.Exec(TableCreate)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}
	return nil
}

func insertMessages() error {
	_, err := execSQL(`
INSERT INTO Messages (UserID, Message) 
VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8), ($9, $10)
`,
		1, "Greetings from the digital realm!",
		2, "Exploring the wonders of code today!",
		3, "Building the future, one line at a time.",
		1, "Hello world!",
		3, "Coding is fun!")
	if err != nil {
		return fmt.Errorf("failed to insert messages: %w", err)
	}
	fmt.Println("Inserted 5 messages into the Messages table")
	return nil
}

func execSQL(statement string, args ...any) (sql.Result, error) {
	stmt, err := db.Prepare(statement)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL statement: %w", err)
	}
	defer stmt.Close()
	rsl, err := stmt.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL statement: %w", err)
	}
	return rsl, nil
}

func querySQL(statement string, args ...any) (*sql.Rows, error) {
	stmt, err := db.Prepare(statement)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare SQL statement: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %w", err)
	}
	return rows, nil
}
