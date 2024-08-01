package interfaces

import "github.com/gin-gonic/gin"

type IGiphy interface {
	GetTrendingGifs(ctx *gin.Context) (string, error)
	GetTrendingStickers(ctx *gin.Context) (string, error)
	SearchGifs(ctx *gin.Context, search, language string) ([]string, error)
	SearchStickers(ctx *gin.Context, search, language string) ([]string, error)
}
