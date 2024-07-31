// Code generated by Wire. DO NOT EDIT.
//go:build !wireinject
// +build !wireinject

package di

import (
	_ "github.com/lib/pq"
	"product-service/internal/api"
	"product-service/internal/api/handler"
	"product-service/internal/config"
	"product-service/internal/db"
	"product-service/internal/repository"
	"product-service/internal/service"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	url, sqlxDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	err = db.Migrate(url)
	if err != nil {
		return nil, err
	}
	productRepository := repository.NewProductRepository(sqlxDB)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	server := http.NewServer(productHandler)
	return server, nil
}
