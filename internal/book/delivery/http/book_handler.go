package http

import (
	"BookStore/internal/domain"
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

func (bh *BookHandler) InitRoutes(bUsecase domain.BookUsecase) *gin.Engine {
	bh.BUsecase = bUsecase
	router := gin.New()
	router.GET("/books", bh.GetAll)
	router.GET("/books/:id", bh.GetById)
	router.POST("/books", bh.Add)
	router.PUT("/books/:id", bh.Update)
	router.DELETE("/books/:id", bh.Delete)
	return router
}

func (bh *BookHandler) Add(ctx *gin.Context) {
	result := bh.BUsecase.Add(&domain.Book{
		ID: 4,
		Title: "Мы",
		Authors: []string{"Евгений Замятин"},
		Year:    "1924"})

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
	result := bh.BUsecase.Update(&domain.Book{
		ID: 4,
		Title: "Мы",
		Authors: []string{"Евгений Замятин"},
		Year:    "1924"}, 1)

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
