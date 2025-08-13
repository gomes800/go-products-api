package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gomes800/go-products-api/internal/adapter/db"
	"github.com/gomes800/go-products-api/internal/adapter/http"
	"github.com/gomes800/go-products-api/internal/domain"
	"github.com/gomes800/go-products-api/internal/usecase"
)

func main() {

	database := db.SetupDB()

	database.AutoMigrate(&domain.Product{})

	repo := db.NewProductRepository(database)
	uc := usecase.NewProductUseCase(repo)

	r := gin.Default()
	http.NewProductHandler(r, uc)

	r.Run(":8080")
}
