package main

import (
	"app/internal/auth"
	"app/internal/repository"
	"app/internal/repository/cache"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

func newTestApplication(t *testing.T, cfg config) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	mockRepository := repository.NewMockRepository()
	mockCacheRepository := cache.NewMockRepository()
	testAuth := &auth.TestAuthenticator{}

	return &application{
		logger:          logger,
		repository:      mockRepository,
		cacheRepository: mockCacheRepository,
		authenticator:   testAuth,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d", expected, actual)
	}
}
