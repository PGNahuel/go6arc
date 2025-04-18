package main

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/pgnahuel/go6arc/internal/adapters/inbound/http"
	"github.com/pgnahuel/go6arc/internal/adapters/outbound"
	"github.com/pgnahuel/go6arc/internal/services"
)

func main() {
	app := chi.NewRouter()

	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	err := SetupBasicHandlers(app)
	if err != nil {
		panic(fmt.Sprintf("fail initializing application %s", err.Error()))
	}
}

func SetupBasicHandlers(app *chi.Mux) error {
	basicReader := outbound.NewBasicReader()
	basicCreator := outbound.NewBasicCreator()
	basicEliminator := outbound.NewBasicEliminator()
	basicRemover := outbound.NewBasicRemover()
	basicService := services.NewBasicService(basicReader, basicCreator, basicEliminator, basicRemover)
	basicHandler := http.NewBasicHandler(basicService)

	app.Get("/", basicHandler.List)
	app.Get("/{id}", basicHandler.Read)
	app.Post("/", basicHandler.Create)
	app.Delete("/{id}", basicHandler.Delete)
	app.Put("/{id}", basicHandler.Update)

	return nil
}
