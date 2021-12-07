package http

import (
	"github.com/Dann-Go/book-store/internal/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

type BookHandler struct {
	BUsecase domain.BookUsecase
	valid    validator.Validate
}
type ResponseError struct {
	Message string `json:"message"`
}

func (bh *BookHandler) NewBookHandler(group *gin.RouterGroup, usecase domain.BookUsecase, validator *validator.Validate) {
	bh.BUsecase = usecase
	bh.valid = *validator
	group.GET("", bh.GetAll)
	group.GET("/:id", bh.GetById)
	group.POST("", bh.Add)
	group.PUT("/:id", bh.Update)
	group.DELETE("/:id", bh.Delete)
}

func (bh *BookHandler) Add(ctx *gin.Context) {
	json := domain.Book{}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := bh.valid.Struct(json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := bh.BUsecase.Add(&json)
	ctx.JSON(http.StatusOK, result)
}
func (bh *BookHandler) GetAll(ctx *gin.Context) {
	result, err := bh.BUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, result)
}
func (bh *BookHandler) GetById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := bh.BUsecase.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, result)
}
func (bh *BookHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	json := domain.Book{}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = bh.valid.Struct(json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := bh.BUsecase.Update(&json, id)
	ctx.JSON(http.StatusOK, result)
}
func (bh *BookHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = bh.BUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "book was deleted")
}
