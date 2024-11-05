package http

import (
	"URL_shortener/Internal/application"
	"net/http"
)

type RedirectHandler struct {
	redirectUseCase application.RedirectHandlerUseCase
}

func NewRedirectHandler(redirectUseCase application.RedirectHandlerUseCase) *RedirectHandler {
	return &RedirectHandler{redirectUseCase: redirectUseCase}
}

func (h *RedirectHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Path[1:]
	originalURL, err := h.redirectUseCase.RedirectURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}
