package product_http

import (
	"github.com/labstack/echo/v4"
	"kom-product-service/domain"
	"net/http"
)

// ProductHandler represent the httphandler for product
type ProductHandler struct {
	ProductUsecase domain.ProductUsecase
}

// New will initialize the product resources endpoint
func New(e *echo.Echo, uc domain.ProductUsecase) {
	handler := &ProductHandler{
		ProductUsecase: uc,
	}
	e.GET("/products", handler.Fetch)
	e.POST("/products", handler.StoreMany)
	e.PUT("/product", handler.Update)
	e.DELETE("/product/:id", handler.Delete)
}

// Fetch will fetch the product based on given params
func (p *ProductHandler) Fetch(c echo.Context) error {
	var products, err = p.ProductUsecase.ListProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, products)
}

// StoreMany will store the books by given request body
func (p *ProductHandler) StoreMany(c echo.Context) error {
	var products []domain.Product
	err := c.Bind(&products)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	result, err := p.ProductUsecase.InsertProduct(products)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, result)
}

// Update will update product by given param
func (p *ProductHandler) Update(c echo.Context) error {
	var product domain.Product
	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	_, err = p.ProductUsecase.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}

// Delete will delete product by given param
func (p *ProductHandler) Delete(c echo.Context) error {
	_, err := p.ProductUsecase.DeleteProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.NoContent(http.StatusOK)
}
