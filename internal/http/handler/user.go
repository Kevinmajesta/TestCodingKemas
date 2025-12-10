package handler

import (
	"kevinmajesta/testkemas/internal/entity"
	"kevinmajesta/testkemas/internal/http/binder"
	"kevinmajesta/testkemas/internal/service"
	"kevinmajesta/testkemas/pkg/response"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return ProductHandler{productService: productService}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	input := binder.ProductCreateRequest{}

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "There is an input error"))
	}

	newProduct := &entity.Products{
		Name:  input.Name,
		Price: input.Price,
		Stock: input.Stock,
	}

	product, err := h.productService.CreateProduct(newProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully input a new product", product))
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	var input binder.ProductUpdateRequest

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "There is an input error"))
	}

	if input.ProductID == uuid.Nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Product ID cannot be empty"))
	}

	if input.Name == "" {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Product name cannot be empty"))
	}
	if input.Price <= 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Price must be greater than 0"))
	}
	if input.Stock < 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Stock must be 0 or more"))
	}

	updatedProduct := entity.UpdateProduct(
		input.ProductID,
		input.Name,
		input.Price,
		input.Stock,
	)

	result, err := h.productService.UpdateProduct(updatedProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "Successfully updated product", result))
}
