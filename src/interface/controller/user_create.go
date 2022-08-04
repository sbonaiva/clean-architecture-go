package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/usecase"
	"github.com/sbonaiva/clean-architecture-go/interface/controller/dto"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type createCustomerController struct {
	usecase usecase.CreateCustomerUseCase
	logger  util.Logger
}

type CreateCustomerController interface {
	Handle(ctx *gin.Context)
}

func NewCreateCustomerController(u usecase.CreateCustomerUseCase, l util.Logger) CreateCustomerController {
	return &createCustomerController{
		usecase: u,
		logger:  l,
	}
}

func (c *createCustomerController) Handle(ctx *gin.Context) {

	var req dto.CustomerRequest

	errBind := ctx.ShouldBindJSON(&req)

	if errBind != nil {
		ctx.JSON(http.StatusInternalServerError, dto.NewDefaultErrorResponse())
	}

	customer := dto.FromCustomerRequest(req)

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

	ctx.JSON(http.StatusCreated, dto.ToCustomerResponse(customer))
}
