package builder

import (
	"kevinmajesta/testkemas/internal/http/handler"
	"kevinmajesta/testkemas/internal/http/router"
	"kevinmajesta/testkemas/internal/repository"
	"kevinmajesta/testkemas/internal/service"
	"kevinmajesta/testkemas/pkg/route"

	"gorm.io/gorm"
)

func BuildPublicRoutes(db *gorm.DB) []*route.Route {
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	return router.PublicRoutes(productHandler)
}
