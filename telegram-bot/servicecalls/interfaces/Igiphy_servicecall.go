package interfaces

import (
	"telegram-bot/dto"

	"github.com/gin-gonic/gin"
)

type IGiphyServiceCall interface {
	GetTrendingGifs(ctx *gin.Context, limit, offset int, rating, randomID, bundle string) (dto.GiphyResponse, error)
	GetTrendingStickers(ctx *gin.Context, limit, offset int, rating, randomID, bundle string) (dto.GiphyResponse, error)
	SearchGifs(ctx *gin.Context, limit, offset int, search, language, rating, randomID, bundle string) (dto.GiphyResponse, error)
	SearchStickers(ctx *gin.Context, limit, offset int, search, language, rating, randomID, bundle string) (dto.GiphyResponse, error)
}
