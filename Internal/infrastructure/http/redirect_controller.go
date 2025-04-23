package http

import (
	"URL_shortener/Internal/application"
	"URL_shortener/Internal/infrastructure/services"
	"net/http"
)

type RedirectController struct {
	redirectUseCase     application.RedirectHandlerUseCase
	methodValidator     *services.HTTPMethodValidator
	queryParamProcessor *services.QueryParamProcessor
}

func NewRedirectController(
	redirectUseCase application.RedirectHandlerUseCase,
	methodValidator *services.HTTPMethodValidator,
	queryParamProcessor *services.QueryParamProcessor,
) *RedirectController {
	return &RedirectController{
		redirectUseCase:     redirectUseCase,
		methodValidator:     methodValidator,
		queryParamProcessor: queryParamProcessor,
	}
}

func (c *RedirectController) RedirectURL(w http.ResponseWriter, r *http.Request) {
	if err := c.methodValidator.ValidateMethod(r, http.MethodGet); err != nil {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := c.queryParamProcessor.ExtractShortURL(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	originalURL, err := c.redirectUseCase.RedirectURL(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusSeeOther)
}
