package handlers_v1

import (
	"telegram-bot/services"
	"telegram-bot/services/interfaces"

	"github.com/gin-gonic/gin"
)

type GiphyHandler struct {
	gs interfaces.IGiphyService
}

func NewGiphyHandler(ctx *gin.Context) GiphyHandler {
	return GiphyHandler{
		gs: services.NewGiphyService(ctx),
	}
}

func (g GiphyHandler) TrendingGifs(ctx *gin.Context) {

	gif, err := g.gs.GetTrendingGifs(ctx)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"gif": gif})
}

func (g GiphyHandler) TrendingStickers(ctx *gin.Context) {

	sticker, err := g.gs.GetTrendingStickers(ctx)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"sticker": sticker})
}

func (g GiphyHandler) SearchGifs(ctx *gin.Context) {

	search := ctx.Query("search")
	language := ctx.Query("language")

	gifs, err := g.gs.SearchGifs(ctx, search, language)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"gifs": gifs})
}

func (g GiphyHandler) SearchStickers(ctx *gin.Context) {

	search := ctx.Query("search")
	language := ctx.Query("language")

	stickers, err := g.gs.SearchStickers(ctx, search, language)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"stickers": stickers})
}
