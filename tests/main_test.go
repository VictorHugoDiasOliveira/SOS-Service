package tests

import (
	"authservice/config"
	"authservice/handlers"
	"authservice/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for POST /user/add
func TestLogin(t *testing.T) {
	router := config.SetupRouter()
	router = handlers.Login(router)

	w := httptest.NewRecorder()

	// Create an example user for testing
	exampleUser := models.User{
		ID:       1,
		Email:    "victorhugo@gmail.com",
		Password: "12345",
	}

	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Compare the response body with the json data of exampleUser
	assert.Equal(t, string(userJson), w.Body.String())
}
