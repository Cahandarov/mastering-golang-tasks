package db_sql_pkg

import "fmt"

func executeQuery() {
	rows, err := DB.Query("select * from students")

	if err != nil {
		fmt.Println("Error executing query", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int

		err = rows.Scan(&id, &name, &age)

		fmt.Printf("ID:%d, Name:%s, Sge:%d\n", id, name, age)

	}
}

func Task4() {
	fmt.Println("Task 4  ****************")
	InitializeDB()
	executeQuery()
	CloseDB()
}
