package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/usecase"
	"github.com/sbonaiva/clean-architecture-go/interface/controller/dto"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type listCustomerController struct {
	usecase usecase.ListCustomerUseCase
	logger  util.Logger
}

type ListCustomerController interface {
	Handle(ctx *gin.Context)
}

func NewListCustomerController(u usecase.ListCustomerUseCase, l util.Logger) ListCustomerController {
	return &listCustomerController{
		usecase: u,
		logger:  l,
	}
}

func (c *listCustomerController) Handle(ctx *gin.Context) {

	customers, err := c.usecase.Execute()

	if err != nil {

		switch e := err.(type) {
		case *domain.CoreError:
			ctx.JSON(e.Status, dto.ToErrorResponse(e))
		default:
			ctx.JSON(http.StatusInternalServerError, dto.NewDefaultErrorResponse())
		}

		return
	}

	ctx.JSON(http.StatusOK, dto.ToCustomerResponseSlice(customers))
}
