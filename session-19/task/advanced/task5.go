package advanced

import (
	"fmt"
	"session-19/task/db_sql_pkg"
)

func preparedStatement() {
	stmt, err := db_sql_pkg.DB.Prepare("insert into students (name, age) values ($1, $2)")

	if err != nil {
		fmt.Println("Error preparing statement", err)
	}

	res, err := stmt.Exec("Emil", 17)
	if err != nil {
		fmt.Println("Error executing statement", err)
	}
	effectedRow, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Error getting rows affected", err)
	}
	if effectedRow > 0 {
		fmt.Println("Insert successful")
	}

}

func Task5() {
	fmt.Println("Task 5  ****************")
	db_sql_pkg.InitializeDB()
	preparedStatement()
	db_sql_pkg.CloseDB()
}
