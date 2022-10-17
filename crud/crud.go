package crud

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_USER_PASSWD")
	dbName := os.Getenv("DB_DATABASE_NAME")
	dbServer := os.Getenv("DB_HOST_URL")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+"tcp("+dbServer+")/"+dbName)
	if err != nil {
		fmt.Println("Connection error")
	}

	return db, nil
}

func DBSelect(queryString string) (*sql.Rows, error) {
	db, err := dbConn()

	if err != nil {
		fmt.Println("Database connection error")
	}

	selectDB, err := db.Query(queryString)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return selectDB, err
}

func DBInsert(queryString string) *sql.Rows {
	db, err := dbConn()

	if err != nil {
		fmt.Println("Database connection error")
	}

	insertDB, err := db.Query(queryString)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return insertDB
}

func DBDelete(queryString string) (int64, error) {
	db, err := dbConn()

	if err != nil {
		fmt.Println("Database connection error")
	}

	deleteDB, err := db.Exec(queryString)

	if err != nil {
		panic(err.Error())
	}

	result, _ := deleteDB.RowsAffected()

	defer db.Close()

	return result, err
}
