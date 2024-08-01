package dto

type GiphyResponse struct {
	Data []struct {
		Type             string `json:"type"`
		Id               string `json:"id"`
		Slug             string `json:"slug"`
		URL              string `json:"url"`
		BitlyGifURL      string `json:"bitly_gif_url"`
		BitlyURL         string `json:"bitly_url"`
		EmbedURL         string `json:"embed_url"`
		Username         string `json:"username"`
		Source           string `json:"source"`
		Rating           string `json:"rating"`
		ContentURL       string `json:"content_url"`
		SourceTld        string `json:"source_tld"`
		SourcePostURL    string `json:"source_post_url"`
		IsSticker        int    `json:"is_sticker"`
		ImportDatetime   string `json:"import_datetime"`
		TrendingDatetime string `json:"trending_datetime"`
		User             struct {
			AvatarURL   string `json:"avatar_url"`
			BannerURL   string `json:"banner_url"`
			ProfileURL  string `json:"profile_url"`
			Username    string `json:"username"`
			DisplayName string `json:"display_name"`
			IsVerified  bool   `json:"is_verified"`
		} `json:"user"`
		Images struct {
			Original struct {
				Height   string `json:"height"`
				Width    string `json:"width"`
				Size     string `json:"size"`
				URL      string `json:"url"`
				Mp4Size  string `json:"mp4_size"`
				Mp4      string `json:"mp4"`
				WebpSize string `json:"webp_size"`
				Webp     string `json:"webp"`
				Frames   string `json:"frames"`
				Hash     string `json:"hash"`
			} `json:"original"`
			Downsized struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized"`
			DownsizedLarge struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_large"`
			DownsizedMedium struct {
				Height string `json:"height"`
				Width  string `json:"width"`
				Size   string `json:"size"`
				URL    string `json:"url"`
			} `json:"downsized_medium"`
			DownsizedSmall struct {
				Height  string `json:"height"`
				Width   string `json:"width"`
				Mp4Size string `json:"mp4_size"`
				Mp4     string `json:"mp4"`
			} `json:"downsized_small"`
		} `json:"images"`
		Title string `json:"title"`
	} `json:"data"`
	Pagination struct {
		TotalCount int `json:"total_count"`
		Count      int `json:"count"`
		Offset     int `json:"offset"`
	} `json:"pagination"`
	Meta struct {
		Status     int    `json:"status"`
		Msg        string `json:"msg"`
		ResponseID string `json:"response_id"`
	} `json:"meta"`
}
