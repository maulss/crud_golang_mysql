package crud_golang_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root/root@tcp(127.0.0.1:3306)/crud_golang_mysql")
	if err != nil {
		panic(err)
	}

	defer db.Close()
}
