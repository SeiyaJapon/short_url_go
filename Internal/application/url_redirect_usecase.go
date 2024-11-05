package application

import (
	"URL_shortener/Internal/domain"
	"URL_shortener/Internal/domain/interfaces"
)

type RedirectHandlerUseCase struct {
	urlRedirecter interfaces.URLRedirecter
}

func NewRedirectHandlerUseCase(urlRedirecter interfaces.URLRedirecter) *RedirectHandlerUseCase {
	return &RedirectHandlerUseCase{urlRedirecter: urlRedirecter}
}

func (u *RedirectHandlerUseCase) RedirectURL(shortURL string) (string, error) {
	myUrl, _ := domain.NewURL(shortURL)
	originalURL, err := u.urlRedirecter.Redirect(myUrl)
	if err != nil {
		return "", err
	}

	return originalURL, nil
}
