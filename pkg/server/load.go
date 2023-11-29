package server

import (
	"log"
	"net/http"

	"github.com/Ankit-2205/put-call-ratio/pkg/optionchain"
	"github.com/Ankit-2205/put-call-ratio/pkg/optionchain/handler"
	"github.com/Ankit-2205/put-call-ratio/pkg/repository"
	"github.com/gofiber/fiber/v2"
)

type container struct {
	handlerContainer handlerContainer
}

type handlerContainer struct {
	handler handler.Handler
}

func loadContainer() (container, error) {

	httpClient := http.Client{}
	return container{
		handlerContainer: handlerContainer{
			handler: handler.NewHandler(optionchain.NewOptionChainService(repository.NewOptionChainRepo(httpClient))),
		},
	}, nil
}

func loadRoutes(app *fiber.App) {

	container, err := loadContainer()
	if err != nil {
		log.Fatalf("failed to load containers: %v", err)
	}

	v1Routes := app.Group("/api/v1")

	v1Routes.Get("/option-chains", container.handlerContainer.handler.GetOptionChain)
}
