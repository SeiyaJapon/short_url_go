package application

import (
	"URL_shortener/Internal/domain"
	"errors"
	"testing"
)

// MockURLShorter implementa la interfaz URLShorter para pruebas
type MockURLShorter struct {
	shortenFunc func(url *domain.URL) (string, error)
}

func (m *MockURLShorter) Shorten(url *domain.URL) (string, error) {
	return m.shortenFunc(url)
}

func TestURLShortenerUseCase_ShortenURL_Success(t *testing.T) {
	// Configurar el mock para simular un acortamiento exitoso
	expectedShortURL := "http://short.url/abc123"
	mockShorter := &MockURLShorter{
		shortenFunc: func(url *domain.URL) (string, error) {
			return expectedShortURL, nil
		},
	}

	// Crear instancia del caso de uso con el mock
	usecase := NewURLShortenerUseCase(mockShorter)

	// Ejecutar el caso de uso
	originalURL := "https://www.example.com/very/long/path"
	result, err := usecase.ShortenURL(originalURL)

	// Verificar resultados
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo: %v", err)
	}

	if result != expectedShortURL {
		t.Errorf("URL incorrecta. Se esperaba: %s, se obtuvo: %s", expectedShortURL, result)
	}
}

func TestURLShortenerUseCase_ShortenURL_InvalidURL(t *testing.T) {
	// Configurar el mock para simular un acortamiento exitoso (no debería llegar a llamarse)
	mockShorter := &MockURLShorter{
		shortenFunc: func(url *domain.URL) (string, error) {
			t.Fatal("No debería haber llegado a llamar a Shorten con una URL inválida")
			return "", nil
		},
	}

	// Crear instancia del caso de uso con el mock
	usecase := NewURLShortenerUseCase(mockShorter)

	// URL inválida que debería fallar en la validación
	invalidURL := "invalid-url"
	result, err := usecase.ShortenURL(invalidURL)

	// Verificar resultados
	if err == nil {
		t.Error("Se esperaba un error por URL inválida, pero no se recibió ninguno")
	}

	if result != "" {
		t.Errorf("Se esperaba una cadena vacía, pero se obtuvo: %s", result)
	}
}

func TestURLShortenerUseCase_ShortenURL_RepositoryError(t *testing.T) {
	// Configurar el mock para simular un error en el repositorio
	mockError := errors.New("error en el repositorio")
	mockShorter := &MockURLShorter{
		shortenFunc: func(url *domain.URL) (string, error) {
			return "", mockError
		},
	}

	// Crear instancia del caso de uso con el mock
	usecase := NewURLShortenerUseCase(mockShorter)

	// Ejecutar el caso de uso
	originalURL := "https://www.example.com/path"
	result, err := usecase.ShortenURL(originalURL)

	// Verificar resultados
	if err == nil {
		t.Error("Se esperaba un error del repositorio, pero no se recibió ninguno")
	}

	if err != mockError {
		t.Errorf("Error incorrecto. Se esperaba: %v, se obtuvo: %v", mockError, err)
	}

	if result != "" {
		t.Errorf("Se esperaba una cadena vacía, pero se obtuvo: %s", result)
	}
}

func TestURLShortenerUseCase_NewURLShortenerUseCase(t *testing.T) {
	// Crear una instancia de mock
	mockShorter := &MockURLShorter{}

	// Crear instancia del caso de uso
	usecase := NewURLShortenerUseCase(mockShorter)

	// Verificar que la instancia se creó correctamente
	if usecase == nil {
		t.Error("NewURLShortenerUseCase devolvió nil")
	}

	if usecase.urlShorter != mockShorter {
		t.Error("El campo urlShorter no se inicializó correctamente")
	}
}
