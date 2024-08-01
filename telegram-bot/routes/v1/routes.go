package routes_v1

import (
	handlers_v1 "telegram-bot/handlers/v1"
	"telegram-bot/utils"

	"github.com/gin-gonic/gin"
)

type route struct {
	routeGroup *gin.RouterGroup
	ctx        *gin.Context
}

func Route(ctx *gin.Context, rGroup *gin.RouterGroup) {
	r := route{
		routeGroup: rGroup.Group("v1"),
		ctx:        ctx,
	}
	utils.ExecuteMethods(r)
}

func (r route) GiphyRoutes() {
	giphyHandler := handlers_v1.NewGiphyHandler(r.ctx)
	r.routeGroup.GET("/trending-gifs", giphyHandler.TrendingGifs)
	r.routeGroup.GET("/trending-stickers", giphyHandler.TrendingStickers)
	r.routeGroup.GET("/search-gifs", giphyHandler.SearchGifs)
	r.routeGroup.GET("/search-stickers", giphyHandler.SearchStickers)
}
