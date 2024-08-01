package services

import (
	"math/rand"
	"telegram-bot/servicecalls"
	"telegram-bot/servicecalls/interfaces"
	"telegram-bot/utils/constants"

	"github.com/gin-gonic/gin"
)

type GiphyService struct {
	gsc interfaces.IGiphyServiceCall
}

func NewGiphyService(ctx *gin.Context) GiphyService {
	return GiphyService{
		gsc: servicecalls.NewGiphyServiceCall(ctx),
	}
}

func (g GiphyService) GetTrendingGifs(ctx *gin.Context) (string, error) {
	gifs, err := g.gsc.GetTrendingGifs(ctx, constants.GIPHY_LIMIT, constants.GIPHY_OFFSET, constants.GIPHY_RATING,
		constants.GIPHY_RANDOM_ID, constants.GIPHY_BUNDLE)

	if err != nil {
		return "", err
	}

	randomGifIndex := rand.Intn(len(gifs.Data))
	return gifs.Data[randomGifIndex].Images.Original.URL, nil
}

func (g GiphyService) GetTrendingStickers(ctx *gin.Context) (string, error) {
	stickers, err := g.gsc.GetTrendingStickers(ctx, constants.GIPHY_LIMIT, constants.GIPHY_OFFSET, constants.GIPHY_RATING,
		constants.GIPHY_RANDOM_ID, constants.GIPHY_BUNDLE)

	if err != nil {
		return "", err
	}

	randomGifIndex := rand.Intn(len(stickers.Data))
	return stickers.Data[randomGifIndex].Images.Original.URL, nil
}

func (g GiphyService) SearchGifs(ctx *gin.Context, search, language string) ([]string, error) {

	gifUrls := make([]string, 0)

	gifs, err := g.gsc.SearchGifs(ctx, constants.GIPHY_LIMIT, constants.GIPHY_OFFSET, search, language, constants.GIPHY_RATING,
		constants.GIPHY_RANDOM_ID, constants.GIPHY_BUNDLE)

	if err != nil {
		return gifUrls, err
	}

	for _, gif := range gifs.Data {
		gifUrls = append(gifUrls, gif.Images.Original.URL)
	}

	return gifUrls, nil
}

func (g GiphyService) SearchStickers(ctx *gin.Context, search, language string) ([]string, error) {
	stickerUrls := make([]string, 0)

	stickers, err := g.gsc.SearchGifs(ctx, constants.GIPHY_LIMIT, constants.GIPHY_OFFSET, search, language, constants.GIPHY_RATING,
		constants.GIPHY_RANDOM_ID, constants.GIPHY_BUNDLE)

	if err != nil {
		return stickerUrls, err
	}

	for _, gif := range stickers.Data {
		stickerUrls = append(stickerUrls, gif.Images.Original.URL)
	}

	return stickerUrls, nil
}
