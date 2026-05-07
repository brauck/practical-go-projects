package crud

import (
	"database/sql"
	"errors"
	"fmt"
)

type Teacher struct {
	id        int
	firstname string
	lastname  string
}

func CreateTeacher(firstname string, lastname string, db *sql.DB) (int, error) {
	insertedId := 0
	err := db.QueryRow("INSERT INTO public.teacher (create_time, firstname, lastname) VALUES (NOW(),$1, $2) RETURNING id;", firstname, lastname).Scan(&insertedId)
	if err != nil {
		return 0, err
	}
	if insertedId == 0 {
		return 0, errors.New("something went wrong id inserted is equal to zero")
	}
	fmt.Println("create teacher ok")
	return insertedId, nil
}

func ReadOneteacher(id int, db *sql.DB) (*Teacher, error) {
	teacher := Teacher{}
	err := db.QueryRow("SELECT id, firstname, lastname FROM teacher WHERE id = $1 ", id).Scan(&teacher.id, &teacher.firstname, &teacher.lastname)
	if err != nil {
		return &teacher, err
	}
	return &teacher, nil
}

func ReadAllTeachers(db *sql.DB) (*[]Teacher, error) {
	rows, err := db.Query("SELECT id, firstname, lastname FROM teacher ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	teachers := make([]Teacher, 0)
	for rows.Next() {
		teacher := Teacher{}
		if err := rows.Scan(&teacher.id, &teacher.firstname, &teacher.lastname); err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return &teachers, nil
}

func UpdateTeacher(db *sql.DB) {
	res, err := db.Exec("UPDATE teacher SET firstname = $1 WHERE id = $2", "Daniel", 1)
	if err != nil {
		fmt.Println(err)
		// query was not a success; something went wrong
	}

	affected, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}

	if affected != 1 {
		fmt.Printf("Something went wrong %d rows were affected expected 1\n", affected)
	} else {
		fmt.Println("Update is a success")
	}
}
