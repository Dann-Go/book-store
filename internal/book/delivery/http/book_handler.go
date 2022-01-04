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

// Add
// @Summary      Add
// @Description  Add new book
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        input body domain.Book  true  "Book info"
// @Success      200  {object}  responses.ServerGoodResponse
// @Failure      400  {object}  responses.ServerBadRequestError
// @Failure      500  {object}  responses.ServerInternalError
// @Router       /api/books [post]
func (bh *BookHandler) Add(ctx *gin.Context) {
	json := domain.Book{}
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.NewServerBadRequestError(err.Error()))
		return
	}
	_ = bh.BUsecase.Add(&json)
	ctx.JSON(http.StatusOK, responses.NewServerGoodResponse("Books was added"))
}

// GetAll
// @Summary      GetAll
// @Description  show all books
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        title   path      string  false  "Book title"
// @Success      200  {object}  []domain.Book
// @Failure      500  {object}  responses.ServerInternalError
// @Router       /api/books [get]
func (bh *BookHandler) GetAll(ctx *gin.Context) {
	var result []domain.Book
	var err error
	if len(ctx.Request.URL.Query().Get("title")) > 0 {
		result, err = bh.BUsecase.GetByTitle(ctx.Request.URL.Query().Get("title"))
	} else {
		result, err = bh.BUsecase.GetAll()
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.NewServerInternalError(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// GetById
// @Summary      GetById
// @Description  Find books by id
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "Book Id"
// @Success      200  {object}  domain.Book
// @Failure      400  {object}  responses.ServerBadRequestError
// @Failure      500  {object}  responses.ServerInternalError
// @Router       /api/books/{id} [get]
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

// Update
// @Summary      Update
// @Description  Update books by id
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "Book Id"
// @Param        input body domain.Book  true  "Book info"
// @Success      200  {object}  responses.ServerGoodResponse
// @Failure      400  {object}  responses.ServerBadRequestError
// @Failure      500  {object}  responses.ServerInternalError
// @Router       /api/books/{id} [put]
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

// Delete
// @Summary      Delete
// @Description  Delete books by id
// @Tags         lists
// @Accept       json
// @Produce      json
// @Param        id  path      int  true  "Book Id"
// @Success      200  {object}  responses.ServerGoodResponse
// @Failure      400  {object}  responses.ServerBadRequestError
// @Failure      500  {object}  responses.ServerInternalError
// @Router       /api/books/{id} [delete]
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
