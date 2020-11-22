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

func ConnectionToPostgre() {
	connStr := "user=postgres password=postgres dbname=freeit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into users (first_name, last_name, age) values ('Владимир', 'Лепешко', $1)", 
        32)
    if err != nil{
        panic(err)
    }
	
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
