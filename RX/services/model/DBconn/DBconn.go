package DBConn

import (
  "database/sql"
  "fmt"
  _ "github.com/lib/pq"
  SQLRead "RX/services/lib/SQLRead"
)

var (
    host     = "127.0.0.1"
    port     = 5432
    user     = "Mohan"
    password = "admin123"
    dbname   = "RX"
)

func DBConn() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
  	db, err := sql.Open("postgres", psqlInfo)
  	return db, err
}

func ReadResultSet(result *sql.Rows) (map[string][]string){
    return SQLRead.CreateResultSet(result)
}