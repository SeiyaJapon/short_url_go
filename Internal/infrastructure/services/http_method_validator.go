package services

import "net/http"

type HTTPMethodValidator struct{}

func (v *HTTPMethodValidator) ValidateMethod(r *http.Request, method string) error {
	if r.Method != method {
		return http.ErrNotSupported
	}
	return nil
}
