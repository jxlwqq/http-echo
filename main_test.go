package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	t.Parallel()
	r := gin.New()
	r.GET("/", TextHandler)
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", resp.Code)
	}
}

func TestHealthHandler(t *testing.T) {
	t.Parallel()
	r := gin.New()
	r.GET("/health", TextHandler)
	req, _ := http.NewRequest("GET", "/health", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", resp.Code)
	}
}

func TestNotFoundHandler(t *testing.T) {
	t.Parallel()
	r := gin.New()
	r.GET("/你好", TextHandler)
	req, _ := http.NewRequest("GET", "/你好", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("Wrong status code: %d", resp.Code)
	}
}
