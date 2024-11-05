package interfaces

import "URL_shortener/Internal/domain"

type URLRedirecter interface {
	Redirect(url *domain.URL) (string, error)
}
