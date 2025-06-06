package app

import (
	"net/http"

	"github.com/glebnaz/cache-webinar/internal/cache"
	"github.com/glebnaz/cache-webinar/internal/model"
	"github.com/labstack/echo/v4"
)

type App struct {
	cache cache.Cache
}

type CreatePostRequest struct {
	Post model.Post `json:"post"`
}

type GetFeedRequest struct {
	Id string `json:"id"`
}

func (a *App) NewCreatePost() echo.HandlerFunc {
	subs := []string{"1"}
	return func(c echo.Context) error {
		var resp CreatePostRequest
		err := c.Bind(&resp)
		if err != nil {
			c.Error(err)
			return err
		}

		err = a.cache.WriteToSubs(c.Request().Context(), resp.Post, subs)
		if err != nil {
			c.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, struct {
			Status string
			Subs   []string
		}{Status: "ok", Subs: subs})
	}
}

func (a *App) NewGetFeed() echo.HandlerFunc {
	return func(c echo.Context) error {
		var resp GetFeedRequest
		err := c.Bind(&resp)
		if err != nil {
			c.Error(err)
			return err
		}

		feed, err := a.cache.ReadFeed(c.Request().Context(), resp.Id)
		if err != nil {
			c.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, feed)
	}
}

func NewApp(cache cache.Cache) App {
	return App{cache: cache}
}
