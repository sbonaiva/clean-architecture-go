package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbonaiva/clean-architecture-go/core/domain"
	"github.com/sbonaiva/clean-architecture-go/core/usecase"
	"github.com/sbonaiva/clean-architecture-go/interface/controller/dto"
	"github.com/sbonaiva/clean-architecture-go/util"
)

type deleteCustomerController struct {
	usecase usecase.DeleteCustomerUseCase
	logger  util.Logger
}

type DeleteCustomerController interface {
	Handle(ctx *gin.Context)
}

func NewDeleteCustomerController(u usecase.DeleteCustomerUseCase, l util.Logger) DeleteCustomerController {
	return &deleteCustomerController{
		usecase: u,
		logger:  l,
	}
}

func (c *deleteCustomerController) Handle(ctx *gin.Context) {

	id := ctx.Param("id")

	err := c.usecase.Execute(id)

	if err != nil {

		switch e := err.(type) {
		case *domain.CoreError:
			ctx.JSON(e.Status, dto.ToErrorResponse(e))
		default:
			ctx.JSON(http.StatusInternalServerError, dto.NewDefaultErrorResponse())
		}

		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
