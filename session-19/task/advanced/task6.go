package advanced

import (
	"fmt"
	"session-19/task/db_sql_pkg"
)

func transactions() {
	tx, err := db_sql_pkg.DB.Begin()
	if err != nil {
		fmt.Println("Error starting transaction", err)
	}
	query := `update students set age=$1 where name=$2`
	_, err = tx.Exec(query, 21, "Ali")

	if err != nil {
		tx.Rollback()
		fmt.Println("Error executing query", err)
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("Error committing transaction", err)
		return
	}

	fmt.Println("Transaction successful")

}
func Task6() {
	fmt.Println("Task 6  ****************")
	db_sql_pkg.InitializeDB()
	transactions()
	db_sql_pkg.CloseDB()
}
