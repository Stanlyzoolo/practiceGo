package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Users struct {
	first_name string
	last_name  string
	age        int
}


func ConnectionToDatabase() {
	var err error
	connStr := "user=postgres password=postgres dbname=freeit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	WrightInDatabase(db, "Stanislav", "Lepeshko", 24)
	ReadFromDatabase(db)
}

func WrightInDatabase(db *sql.DB, first_name, last_name string, age int) {
	_, err := db.Exec("insert into users (first_name, last_name, age) values ($1, $2, $3)",
		first_name, last_name, age)
	if err != nil {
		panic(err)
	}
}

func ReadFromDatabase(db *sql.DB) {
	rows, err := db.Query("select * from users")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	users := []Users{}

	for rows.Next() {
		p := Users{}
		err := rows.Scan(&p.first_name, &p.last_name, &p.age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}

	for _, p := range users {
		fmt.Println("Результат запроса: ", p.first_name, p.last_name, p.age)
	}
}
