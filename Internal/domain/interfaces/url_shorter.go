package interfaces

import "URL_shortener/Internal/domain"

type URLShorter interface {
	Shorten(url *domain.URL) (string, error)
}
