package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abdullinmm/todoapp/internal/auth"
	"github.com/abdullinmm/todoapp/internal/handlers"
)

// TestAuthMiddleware tests the AuthMiddleware function.
func TestAuthMiddleware(t *testing.T) {
	jwtSecret := "supersecretkey123"
	nextCalled := false

	// Wrap in Middleware Wolf handler
	handler := handlers.AuthMiddleware(jwtSecret, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCalled = true
		w.WriteHeader(http.StatusOK)
	}))

	// 1. Without token
	req1 := httptest.NewRequest(http.MethodGet, "/protected", nil)
	rr1 := httptest.NewRecorder()
	handler.ServeHTTP(rr1, req1)
	if rr1.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", rr1.Code)
	}
	if nextCalled {
		t.Errorf("next handler should not be called without token")
	}

	// 2. Garbage token
	nextCalled = false
	req2 := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req2.Header.Set("Authorization", "Bearer invalid.token.here")
	rr2 := httptest.NewRecorder()
	handler.ServeHTTP(rr2, req2)
	if rr2.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", rr2.Code)
	}
	if nextCalled {
		t.Errorf("next handler should not be called with invalid token")
	}

	// 3. Valid token
	nextCalled = false
	token, err := auth.GenerateJWT(123, jwtSecret)
	if err != nil {
		t.Fatalf("failed to generate jwt: %v", err)
	}
	req3 := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req3.Header.Set("Authorization", "Bearer "+token)
	rr3 := httptest.NewRecorder()
	handler.ServeHTTP(rr3, req3)
	if rr3.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr3.Code)
	}
	if !nextCalled {
		t.Errorf("next handler was not called with valid token")
	}
}
