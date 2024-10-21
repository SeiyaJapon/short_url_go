package http

type URLShortenerResponse struct {
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}
