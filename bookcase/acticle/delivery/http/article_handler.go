package http

import (
	"context"
	"github.com/labstack/echo"
	"github.com/neocxf/go-exercises/bookcase/acticle"
	"github.com/neocxf/go-exercises/bookcase/models"
	"github.com/sirupsen/logrus"
	validator2 "gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
)

type ResponseErr struct {
	Message string `json:"message"`
}

type HttpArticleHandler struct {
	AService article.ArticleService
}

func NewArticleHttpHeader(e *echo.Echo, articleService article.ArticleService) {
	handler := &HttpArticleHandler{
		AService: articleService,
	}

	e.GET("/article", handler.FetchArticle)
	e.POST("/article", handler.Store)
	e.GET("/article/:id", handler.GetById)
	e.DELETE("/article/:id", handler.Delete)
}

func (a *HttpArticleHandler) FetchArticle(c echo.Context) error {
	num, _ := strconv.Atoi(c.QueryParam("num"))
	cursor := c.QueryParam("cursor")
	ctx := c.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	listAr, nextCursor, err := a.AService.Fetch(ctx, cursor, int64(num))

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseErr{Message: err.Error()})
	}

	c.Response().Header().Set("X-Cursor", nextCursor)

	return c.JSON(http.StatusOK, listAr)
}

func (a *HttpArticleHandler) GetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	art, err := a.AService.GetById(ctx, int64(id))

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseErr{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)

}

func (a *HttpArticleHandler) Store(c echo.Context) error {
	var article models.Article
	err := c.Bind(&article)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&article); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	ar, err := a.AService.Store(ctx, &article)

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseErr{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ar)
}

func (a *HttpArticleHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	_, err = a.AService.Delete(ctx, int64(id))

	if err != nil {
		return c.JSON(getStatusCode(err), ResponseErr{Message: err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func isRequestValid(ar *models.Article) (bool, error) {
	validate := validator2.New()

	err := validate.Struct(ar)

	if err != nil {
		return false, err
	}

	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)

	switch err {
	case models.CONFLICT_ERROR:
		return http.StatusConflict
	case models.NOT_FOUND_ERROR:
		return http.StatusNotFound
	case models.INTERNAL_SERVER_ERROR:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}

	return 0
}
