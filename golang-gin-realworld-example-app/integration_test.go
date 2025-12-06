package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Add routes: /api/users, /api/articles, etc.
	RegisterRoutes(r)
	return r
}

// ------------------- Authentication Tests -------------------
func TestUserRegistrationFlow(t *testing.T) {
	router := setupRouter()
	body := `{"user":{"username":"testuser","email":"test@example.com","password":"password"}}`
	req, _ := http.NewRequest("POST", "/api/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUserLoginFlow(t *testing.T) {
	router := setupRouter()
	body := `{"user":{"email":"test@example.com","password":"password"}}`
	req, _ := http.NewRequest("POST", "/api/users/login", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetCurrentUser(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/api/user", nil)
	req.Header.Set("Authorization", "Token valid_jwt_token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

// ------------------- Article CRUD Tests -------------------
func TestCreateArticleAuthenticated(t *testing.T) {
	router := setupRouter()
	body := `{"article":{"title":"New Article","description":"Desc","body":"Body"}}`
	req, _ := http.NewRequest("POST", "/api/articles", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", "Token valid_jwt_token")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateArticleUnauthenticated(t *testing.T) {
	router := setupRouter()
	body := `{"article":{"title":"New Article","description":"Desc","body":"Body"}}`
	req, _ := http.NewRequest("POST", "/api/articles", bytes.NewBuffer([]byte(body)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 401, w.Code)
}

func TestListArticles(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/api/articles", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetSingleArticle(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/api/articles/test-article", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateArticleAuthorized(t *testing.T) {
	router := setupRouter()
	body := `{"article":{"title":"Updated Title"}}`
	req, _ := http.NewRequest("PUT", "/api/articles/test-article", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", "Token valid_jwt_token")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateArticleUnauthorized(t *testing.T) {
	router := setupRouter()
	body := `{"article":{"title":"Updated Title"}}`
	req, _ := http.NewRequest("PUT", "/api/articles/test-article", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", "Token invalid_jwt")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 403, w.Code)
}

func TestDeleteArticleAuthorized(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/api/articles/test-article", nil)
	req.Header.Set("Authorization", "Token valid_jwt_token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteArticleUnauthorized(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/api/articles/test-article", nil)
	req.Header.Set("Authorization", "Token invalid_jwt")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 403, w.Code)
}

// ------------------- Article Interaction Tests -------------------
func TestFavoriteArticle(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("POST", "/api/articles/test-article/favorite", nil)
	req.Header.Set("Authorization", "Token valid_jwt_token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUnfavoriteArticle(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/api/articles/test-article/favorite", nil)
	req.Header.Set("Authorization", "Token valid_jwt_token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateComment(t *testing.T) {
	router := setupRouter()
	body := `{"comment":{"body":"Great article"}}`
	req, _ := http.NewRequest("POST", "/api/articles/test-article/comments", bytes.NewBuffer([]byte(body)))
	req.Header.Set("Authorization", "Token valid_jwt_token")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestListComments(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("GET", "/api/articles/test-article/comments", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteComment(t *testing.T) {
	router := setupRouter()
	req, _ := http.NewRequest("DELETE", "/api/articles/test-article/comments/1", nil)
	req.Header.Set("Authorization", "Token valid_jwt_token")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}
