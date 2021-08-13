package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/powsianik/thinking-in-code/internal/config"
	"testing"
)

func TestRoutes(t *testing.T){
	var app config.AppConfig

	mux := routes(&app)
	switch mux.(type) {
	case *chi.Mux:
	default:
		t.Error("type is not http.Handler")
	}
}