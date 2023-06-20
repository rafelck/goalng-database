package golang_database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestEmpty(t *testing.T) {

}

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "rafelck:password@tcp(localhost:3306)/go_restful_api")

	if err != nil {
		panic(err)
	}

	defer db.Close()
}
