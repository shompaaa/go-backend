package product

import (
	"ecommerce/repo"
	"ecommerce/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middlewares
	porductRepo repo.ProductRepo
}

func NewHandler(middlewares *middlewares.Middlewares,
	productRepo repo.ProductRepo,
	) *Handler {
	return &Handler{
		middlewares: middlewares,
		porductRepo: productRepo,
	}
}
