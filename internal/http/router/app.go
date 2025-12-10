package router

import (
	"net/http"

	"kevinmajesta/testkemas/internal/http/handler"
	"kevinmajesta/testkemas/pkg/route"
)

func PublicRoutes(productHandler handler.ProductHandler) []*route.Route {
	return []*route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/products",
			Handler: productHandler.CreateProduct,
		},
		{
			Method:  http.MethodPut,
			Path:    "/products/:product_id",
			Handler: productHandler.UpdateProduct,
		},
	}
}
