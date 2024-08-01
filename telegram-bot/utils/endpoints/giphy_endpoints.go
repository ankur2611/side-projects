package endpoints

// Giphy API endpoints
const (
	TRENDING_GIFS_URL = "/v1/gifs/trending?api_key=%s&limit=%d&offset=%d&rating=%s&random_id=%s&bundle=%s"
	SEARCH_GIFS_URL   = "/v1/gifs/search?api_key=%s&q=%s&lang=%s&limit=%d&offset=%d&rating=%s&random_id=%s&bundle=%s"

	TRENDING_STICKERS_URL = "/v1/stickers/trending?api_key=%s&limit=%d&offset=%d&rating=%s&random_id=%s&bundle=%s"
	SEARCH_STICKERS_URL   = "/v1/stickers/search?api_key=%s&q=%s&lang=%s&limit=%d&offset=%d&rating=%s&random_id=%s&bundle=%s"
)
