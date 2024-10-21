package http

import (
	"URL_shortener/Internal/application"
	"encoding/json"
	"net/http"
)

type Handler struct {
	urlShortenerUseCase application.URLShortenerUseCase
}

func NewHandler(useCase application.URLShortenerUseCase) *Handler {
	return &Handler{
		urlShortenerUseCase: useCase,
	}
}

func (h *Handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
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
