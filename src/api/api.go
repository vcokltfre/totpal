package api

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/vcokltfre/totpal/src/web"
)

func Start(bind string) error {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	games := map[string]*game{}

	e.GET("/", func(c echo.Context) error {
		return c.HTML(200, web.IndexHTML)
	})

	e.POST("/api/games", func(c echo.Context) error {
		id := uuid.NewString()
		games[id] = &game{
			players: map[string]string{},
		}

		logrus.Infof("Created game %s", id)

		return c.JSON(201, map[string]string{
			"id": id,
		})
	})

	e.GET("/api/games/:id", func(c echo.Context) error {
		id := c.Param("id")
		game, exists := games[id]
		if !exists {
			return c.NoContent(404)
		}

		playerId := c.QueryParam("player_id")
		if playerId == "" {
			return c.JSON(200, map[string]any{
				"articles_submitted": game.ArticlesSubmitted(),
				"selected_article":   game.SelectedArticle(),
				"player_id":          uuid.NewString(),
			})
		}

		return c.JSON(200, map[string]any{
			"articles_submitted": game.ArticlesSubmitted(),
			"selected_article":   game.SelectedArticle(),
		})
	})

	e.POST("/api/games/:id/articles", func(c echo.Context) error {
		id := c.Param("id")
		game, exists := games[id]
		if !exists {
			return c.NoContent(404)
		}

		playerId := c.QueryParam("player_id")
		if playerId == "" {
			return c.NoContent(400)
		}

		type requestBody struct {
			Title string `json:"title"`
		}
		body := &requestBody{}
		if err := c.Bind(body); err != nil {
			return c.NoContent(400)
		}

		game.AddArticle(playerId, body.Title)

		return c.JSON(201, nil)
	})

	e.GET("/api/games/:id/articles/pick", func(c echo.Context) error {
		id := c.Param("id")
		game, exists := games[id]
		if !exists {
			return c.NoContent(404)
		}

		article := game.PickArticle()
		if article == "" {
			return c.NoContent(400)
		}

		return c.JSON(200, map[string]string{
			"title": article,
		})
	})

	logrus.Info("Starting API on ", bind)

	return e.Start(bind)
}
