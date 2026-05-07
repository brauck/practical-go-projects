// data-storage/postgresql/connection/main.go
package main

import (
	"database/sql"
	"fmt"

	"data-storage/postgresql/connection/crud"

	_ "github.com/lib/pq"
)

func main() {
	//isTableEcists := false

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=school password=root sslmode=disable")
	if err != nil {
		fmt.Println("db open", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	/* f, err := os.Open("sql_scripts/create_table.sql")
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := db.Exec(string(b))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("create table", res) */

	//crud.CreateTeacher("name5", "surname5", db)

	teacher, err := crud.ReadOneteacher(2, db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("get one teacher: ", *teacher)

	teachers, err := crud.ReadAllTeachers(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*teachers)

	crud.UpdateTeacher(db)
	teachers, err = crud.ReadAllTeachers(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*teachers)
}
