package tests

// import (
// 	"authservice/config"
// 	"authservice/controllers"
// 	"authservice/models"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strconv"
// 	"strings"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// // Test for POST /user/add
// func TestGetUserById(t *testing.T) {
// 	router := config.SetupRouter()
// 	router = controllers.GetUserById(router)

// 	w := httptest.NewRecorder()

// 	// Create an example user for testing
// 	exampleUser := models.User{
// 		ID:       1,
// 		Email:    "victorhugo@gmail.com",
// 		Password: "12345",
// 	}

// 	userJson, _ := json.Marshal(exampleUser)
// 	req, _ := http.NewRequest("GET", "/users/"+strconv.Itoa(exampleUser.ID), strings.NewReader(string(userJson)))
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	assert.Equal(t, string(userJson), w.Body.String())
// }