package http

import (
	"URL_shortener/Internal/application"
	"encoding/json"
	"net/http"
)

type URLHandler struct {
	urlShortenerUseCase application.URLShortenerUseCase
}

func NewShortenHandler(useCase application.URLShortenerUseCase) *URLHandler {
	return &URLHandler{
		urlShortenerUseCase: useCase,
	}
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req URLShortenerRequests
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := h.urlShortenerUseCase.ShortenURL(req.OriginalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := URLShortenerResponse{
		OriginalURL: req.OriginalURL,
		ShortURL:    shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
