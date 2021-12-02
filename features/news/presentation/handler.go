package presentation

import (
	"net/http"
	"antonio/features/news"
	"github.com/labstack/echo/v4"
)

type NewsHandler struct {
	newsService news.Service
}

func NewNewsHandler(ns news.Service) *NewsHandler {
	return &NewsHandler{ns}
}

func (ns *NewsHandler) GetNewsHandler(e echo.Context) error {
	keyword := e.QueryParam("keyword")
	data, err := ns.newsService.GetNews(keyword)
	if err != nil {
		return e.JSON(http.StatusBadRequest, nil)
	}
	return e.JSON(http.StatusOK, data)
}
