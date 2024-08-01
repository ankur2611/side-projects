package servicecalls

import (
	"encoding/json"
	"fmt"
	"telegram-bot/dto"
	"telegram-bot/servicecalls/interfaces"
	"telegram-bot/utils/endpoints"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type GiphyServiceCall struct {
	sc     interfaces.IServiceCalls
	url    string
	apiKey string
}

func NewGiphyServiceCall(ctx *gin.Context) GiphyServiceCall {
	return GiphyServiceCall{
		sc:     NewServiceCall(ctx),
		url:    viper.GetString("GIPHY_URL"),
		apiKey: viper.GetString("GIPHY_API_KEY"),
	}
}

func (g GiphyServiceCall) GetTrendingGifs(ctx *gin.Context, limit, offset int,
	rating, randomID, bundle string) (dto.GiphyResponse, error) {

	var response dto.GiphyResponse
	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TRENDING_GIFS_URL,
		g.apiKey, limit, offset, rating, randomID, bundle))

	byteResponse, err := g.sc.Get(ctx, endpoint, nil)

	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}

	return response, nil
}

func (g GiphyServiceCall) GetTrendingStickers(ctx *gin.Context, limit, offset int,
	rating, randomID, bundle string) (dto.GiphyResponse, error) {

	var response dto.GiphyResponse
	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.TRENDING_STICKERS_URL,
		g.apiKey, limit, offset, rating, randomID, bundle))

	byteResponse, err := g.sc.Get(ctx, endpoint, nil)

	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}

	return response, nil
}

func (g GiphyServiceCall) SearchGifs(ctx *gin.Context, limit, offset int,
	search, language, rating, randomID, bundle string) (dto.GiphyResponse, error) {

	var response dto.GiphyResponse
	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.SEARCH_GIFS_URL,
		g.apiKey, search, language, limit, offset, rating, randomID, bundle))

	byteResponse, err := g.sc.Get(ctx, endpoint, nil)

	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}

	return response, nil
}

func (g GiphyServiceCall) SearchStickers(ctx *gin.Context, limit, offset int,
	search, language, rating, randomID, bundle string) (dto.GiphyResponse, error) {

	var response dto.GiphyResponse
	endpoint := fmt.Sprintf("%s%s", g.url, fmt.Sprintf(endpoints.SEARCH_STICKERS_URL,
		g.apiKey, search, language, limit, offset, rating, randomID, bundle))

	byteResponse, err := g.sc.Get(ctx, endpoint, nil)

	if err != nil {
		return response, err
	}

	if err := json.Unmarshal(byteResponse, &response); err != nil {
		return response, err
	}

	return response, nil
}
