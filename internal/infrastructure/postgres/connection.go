package postgres

import (
	"database/sql"
	_"github.com/lib/pq"
	"fmt"
)  

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres_user"
    password = "postgres_password"
    dbname   = "postgres_db"
)

func NewConnection()(*sql.DB, error) {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil, err
    }
    
    err = db.Ping()
    if err != nil {
        return nil, err
    }

	return db, nil 
}