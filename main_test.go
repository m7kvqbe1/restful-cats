package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetCats(t *testing.T) {
	app := fiber.New()
	app.Get("/cats", getCats)

	req := httptest.NewRequest("GET", "/cats", nil)
	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusOK, resp.StatusCode, "they should be equal")

	var catsResponse []*Cat
	json.NewDecoder(resp.Body).Decode(&catsResponse)

	assert.Equal(t, cats, catsResponse, "they should be equal")
}

func TestCreateCat(t *testing.T) {
	app := fiber.New()
	app.Post("/cats", createCat)

	cat := &Cat{ID: "3", Name: "Fluffy", Age: 2}
	catJSON, _ := json.Marshal(cat)
	req := httptest.NewRequest("POST", "/cats", bytes.NewBuffer(catJSON))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, http.StatusCreated, resp.StatusCode, "they should be equal")

	var catResponse Cat
	json.NewDecoder(resp.Body).Decode(&catResponse)

	assert.Equal(t, cat, &catResponse, "they should be equal")
}
