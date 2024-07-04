package controllers

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
    router := gin.Default()
    router.GET("/healthcheck", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "UP"})
    })

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/healthcheck", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Equal(t, "{\"status\":\"UP\"}", w.Body.String())
}
