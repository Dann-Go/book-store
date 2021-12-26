package http

import (
	"github.com/Dann-Go/book-store/internal/domain"
	"github.com/Dann-Go/book-store/internal/domain/responses"
	"github.com/Dann-Go/book-store/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookHandler struct {
	BUsecase domain.BookUsecase
}
type ResponseError struct {
	Message string `json:"message"`
}

func (bh *BookHandler) NewBookHandler(group *gin.RouterGroup, usecase domain.BookUsecase) {
	bh.BUsecase = usecase
	group.GET("", bh.GetAll)
	group.GET("/:id", bh.GetById)
	group.POST("", bh.Add)
	group.PUT("/:id", bh.Update)
	group.DELETE("/:id", bh.Delete)
}

func (bh *BookHandler) Add(ctx *gin.Context) {
	json := domain.Book{}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	_ = bh.BUsecase.Add(&json)
	ctx.JSON(http.StatusOK, responses.NewServerGoodResponse("Books was added"))
}
func (bh *BookHandler) GetAll(ctx *gin.Context) {
	result, err := bh.BUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result)
}
func (bh *BookHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	result, err := bh.BUsecase.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	middleware.BOOKS_RESERVED.WithLabelValues(ctx.Param("id")).Inc()
	ctx.JSON(http.StatusOK, result)
}
func (bh *BookHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	json := domain.Book{}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	_ = bh.BUsecase.Update(&json, id)
	ctx.JSON(http.StatusOK, responses.NewServerGoodResponse("Books was updated"))
}
func (bh *BookHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	err = bh.BUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, responses.NewServerGoodResponse("Books was deleted"))
}
