package application

import (
	"URL_shortener/Internal/domain"
	"errors"
	"testing"
)

// MockURLRedirecter implementa la interfaz URLRedirecter para pruebas
type MockURLRedirecter struct {
	redirectFunc func(url *domain.URL) (string, error)
}

func (m *MockURLRedirecter) Redirect(url *domain.URL) (string, error) {
	return m.redirectFunc(url)
}

func TestRedirectHandlerUseCase_RedirectURL_Success(t *testing.T) {
	// Configurar el mock para simular un redireccionamiento exitoso
	mockRedirecter := &MockURLRedirecter{
		redirectFunc: func(url *domain.URL) (string, error) {
			return "https://www.example.com/original-page", nil
		},
	}

	// Crear instancia del caso de uso con el mock
	usecase := NewRedirectHandlerUseCase(mockRedirecter)

	// Ejecutar el caso de uso
	result, err := usecase.RedirectURL("abc123")

	// Verificar resultados
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo: %v", err)
	}

	expectedURL := "https://www.example.com/original-page"
	if result != expectedURL {
		t.Errorf("URL incorrecta. Se esperaba: %s, se obtuvo: %s", expectedURL, result)
	}
}

func TestRedirectHandlerUseCase_RedirectURL_Error(t *testing.T) {
	// Configurar el mock para simular un error
	mockError := errors.New("URL no encontrada")
	mockRedirecter := &MockURLRedirecter{
		redirectFunc: func(url *domain.URL) (string, error) {
			return "", mockError
		},
	}

	// Crear instancia del caso de uso con el mock
	usecase := NewRedirectHandlerUseCase(mockRedirecter)

	// Ejecutar el caso de uso
	result, err := usecase.RedirectURL("invalid-short-url")

	// Verificar resultados
	if err == nil {
		t.Error("Se esperaba un error, pero no se recibió ninguno")
	}

	if err != mockError {
		t.Errorf("Error incorrecto. Se esperaba: %v, se obtuvo: %v", mockError, err)
	}

	if result != "" {
		t.Errorf("Se esperaba una cadena vacía, pero se obtuvo: %s", result)
	}
}

func TestRedirectHandlerUseCase_NewRedirectHandlerUseCase(t *testing.T) {
	// Crear una instancia de mock
	mockRedirecter := &MockURLRedirecter{}

	// Crear instancia del caso de uso
	usecase := NewRedirectHandlerUseCase(mockRedirecter)

	// Verificar que la instancia se creó correctamente
	if usecase == nil {
		t.Error("NewRedirectHandlerUseCase devolvió nil")
	}

	if usecase.urlRedirecter != mockRedirecter {
		t.Error("El campo urlRedirecter no se inicializó correctamente")
	}
}
