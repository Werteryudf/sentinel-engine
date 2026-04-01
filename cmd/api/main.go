package main

import (
	"github.com/Werteryudf/sentinel-engine/internal/infrastructure/postgres"
	"fmt"
)  

func main() {
    db, err := postgres.NewConnection()

	if err != nil {
        fmt.Println("Не подключилось")
		panic(err)

    }

	fmt.Println("Подключилось")
    defer db.Close()
    
}