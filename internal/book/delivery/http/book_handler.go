package http

import (
	"github.com/Dann-Go/book-store/internal/domain"
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
func (bh *BookHandler)NewBookHandler(group *gin.RouterGroup, usecase domain.BookUsecase) {
	bh.BUsecase = usecase
	group.GET("", bh.GetAll)
	group.GET("/:id", bh.GetById)
	group.POST("", bh.Add)
	group.PUT("/:id", bh.Update)
	group.DELETE("/:id", bh.Delete)
}

func (bh *BookHandler) Add(ctx *gin.Context) {
	json := domain.Book{}
	if err := ctx.ShouldBindJSON(&json);err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result := bh.BUsecase.Add(&json)

	ctx.JSON(http.StatusOK,result)
}
func (bh *BookHandler) GetAll(ctx *gin.Context) {
	result, err := bh.BUsecase.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, result)
}
func (bh *BookHandler) GetById(ctx *gin.Context) {
	id, _:= strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	result, err := bh.BUsecase.GetById(id)
	if err != nil {

	}
	ctx.JSON(http.StatusOK, result)
}
func (bh *BookHandler) Update(ctx *gin.Context) {
	id, _:= strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	json := domain.Book{}
	if err := ctx.ShouldBindJSON(&json);err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result := bh.BUsecase.Update(&json,id)
	ctx.JSON(http.StatusOK,result)
}
func (bh *BookHandler) Delete(ctx *gin.Context) {
	id, _:= strconv.Atoi(ctx.Request.URL.Query().Get("id"))
	err := bh.BUsecase.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, "book was deleted")
}
