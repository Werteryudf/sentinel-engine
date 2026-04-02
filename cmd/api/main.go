package main

import (
	"github.com/Werteryudf/sentinel-engine/internal/infrastructure/postgres"
	"fmt"
	"github.com/Werteryudf/sentinel-engine/internal/repository"
	"net/http"
    handler "github.com/Werteryudf/sentinel-engine/internal/transport/http"
	"github.com/go-chi/chi/v5"
)
 

func main() {
    db, err := postgres.NewConnection()

	if err != nil {
        fmt.Println("Не подключилось")
		panic(err)

    }

	fmt.Println("Подключилось")
    defer db.Close()

	assetRepo := repository.NewAssetRepository(db)
	assetHandler := handler.NewAssetHandler(assetRepo)  
	
	r := chi.NewRouter()
	r.Get("/assets", assetHandler.GetAll)
	http.ListenAndServe(":8080", r)
}