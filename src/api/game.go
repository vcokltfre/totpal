package api

import "math/rand/v2"

type game struct {
	players  map[string]string
	selected string
}

func (g *game) AddArticle(playerId, title string) {
	g.players[playerId] = title
}

func (g *game) ArticlesSubmitted() int {
	return len(g.players)
}

func (g *game) PickArticle() string {
	articles := []string{}
	for _, title := range g.players {
		articles = append(articles, title)
	}

	if len(articles) == 0 {
		return ""
	}

	article := articles[rand.IntN(len(articles))]
	g.selected = article
	return article
}

func (g *game) SelectedArticle() string {
	return g.selected
}
