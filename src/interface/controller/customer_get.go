package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/usecase"
	"github.com/sbonaiva/clean-architecture-go/interface/controller/dto"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type getCustomerController struct {
	usecase usecase.GetCustomerUseCase
	logger  util.Logger
}

type GetCustomerController interface {
	Handle(ctx *gin.Context)
}

func NewGetCustomerController(u usecase.GetCustomerUseCase, l util.Logger) GetCustomerController {
	return &getCustomerController{
		usecase: u,
		logger:  l,
	}
}

func (c *getCustomerController) Handle(ctx *gin.Context) {

	id := ctx.Param("id")

	customer, err := c.usecase.Execute(id)

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
