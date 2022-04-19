package main

import (
	"hexagonal/domain"
	h "hexagonal/helpers"
	"log"
	"net/http"

	"hexagonal/repositories"
)

func main() {

	// repo, _ := repository.NewMongoRepository(conf.Database.URL, conf.Database.DB, conf.Database.Timeout)
	repo := repositories.NewMysqlRepository(h.ENV("DB_URL"))
	service := domain.NewProductService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/products/{code}", handler.Get)
	r.Post("/products", handler.Post)
	r.Delete("/products/{code}", handler.Delete)
	r.Get("/products", handler.GetAll)
	r.Put("/products", handler.Put)
	log.Fatal(http.ListenAndServe(conf.Server.Port, r))

}
