package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

type Query struct {
	dbConnection *sql.DB
	whereSQL     []string
}

func (query *Query) where(where string) *Query {
	query.whereSQL = append(query.whereSQL, where)

	return query
}

func (query *Query) get() (*sql.Rows, error) {
	return query.execute()
}

func (query *Query) execute() (*sql.Rows, error) {
	sql := "select * from table where" + strings.Join(query.whereSQL, "")

	return query.dbConnection.Query(sql)
}

func LearnDatabase() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}

	query := Query{dbConnection: db}

	query.where("columnA = valueA").where("ColumnB = valueB").get()

	defer db.Close()
}
