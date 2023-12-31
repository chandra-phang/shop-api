package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shop-api/api/controller"
	v1request "shop-api/dto/request/v1"
	"shop-api/infrastructure/log"
	"shop-api/services"

	"github.com/labstack/echo/v4"
)

type productController struct {
	svc services.IProductService
}

// creates a new instance of this controller with reference to ProductService
func InitProductController() *productController {
	//  initializes its "svc" field with a service instance returned by "application.GetProductService()".
	return &productController{
		svc: services.GetProductService(),
	}
}

func (c *productController) CreateProduct(ctx echo.Context) error {
	reqBody, _ := ioutil.ReadAll(ctx.Request().Body)
	dto := v1request.CreateProductDTO{}

	if err := json.Unmarshal(reqBody, &dto); err != nil {
		log.Errorf(ctx, err, "[ProductController][CreateProduct] Failed to unmarshal request body %v into dto", reqBody)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	if err := json.Unmarshal(reqBody, &dto); err != nil {
		log.Errorf(ctx, err, "[ProductController][CreateProduct] Failed to unmarshal request body %v into dto", reqBody)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err := dto.Validate(ctx)
	if err != nil {
		log.Errorf(ctx, err, "[ProductController][CreateProduct] Validation failed for request dto %v ", dto)
		return controller.WriteError(ctx, http.StatusBadRequest, err)
	}

	err = c.svc.CreateProduct(ctx, dto)
	if err != nil {
		log.Errorf(ctx, err, "[ProductController][CreateProduct] Failed to create product for request dto %v ", dto)
		return controller.WriteError(ctx, http.StatusInternalServerError, err)
	}

	return controller.WriteSuccess(ctx, http.StatusCreated, nil)
}
