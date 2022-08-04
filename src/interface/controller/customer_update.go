package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/usecase"
	"github.com/sbonaiva/clean-architecture-go/interface/controller/dto"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type updateCustomerController struct {
	usecase usecase.UpdateCustomerUseCase
	logger  util.Logger
}

type UpdateCustomerController interface {
	Handle(ctx *gin.Context)
}

func NewUpdateCustomerController(u usecase.UpdateCustomerUseCase, l util.Logger) UpdateCustomerController {
	return &updateCustomerController{
		usecase: u,
		logger:  l,
	}
}

func (c *updateCustomerController) Handle(ctx *gin.Context) {

	id := ctx.Param("id")

	var req dto.CustomerRequest

	errBind := ctx.ShouldBindJSON(&req)

	if errBind != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewDefaultErrorResponse())
	}

	customer := dto.FromCustomerRequest(req)
	customer.ID = id

	err := c.usecase.Execute(customer)

	if err != nil {

		switch e := err.(type) {
		case *domain.CoreError:
			ctx.JSON(e.Status, dto.ToErrorResponse(e))
		default:
			ctx.JSON(http.StatusInternalServerError, dto.NewDefaultErrorResponse())
		}

		return
	}

	ctx.JSON(http.StatusOK, dto.ToCustomerResponse(customer))
}
