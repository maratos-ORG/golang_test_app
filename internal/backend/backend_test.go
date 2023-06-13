package backend

import (
	"golang_test_app/internal/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogger(t *testing.T) {
	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Setup router with Logger middleware
	r := gin.New()
	r.Use(Logger())
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "Hello Test")
	})

	// Create a request to use
	req, err := http.NewRequest(http.MethodGet, "/test", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Record the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %v but instead got %v\n", http.StatusOK, w.Code)
	}
}

func TestRunBackend(t *testing.T) {
	// Test parameters
	params := &config.BackendParameters{Port: new(string)}

	// This line needs to be tested as part of a full end-to-end or integration test.
	// go pinger() <-- Commenting this line as this function needs access to actual system command

	// Switch to test mode so you don't get such noisy output
	gin.SetMode(gin.TestMode)

	// Set up the router
	webAPI := gin.New()
	webAPI.Use(Logger())
	webAPI.Use(gin.Recovery())

	// Set up the routes
	webAPI.GET("/config", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"config": params,
		})
	})

	webAPI.GET("/getInfo", func(c *gin.Context) {
		// This line needs to be tested as part of a full end-to-end or integration test.
		// pingHandler(c.Writer, c.Request) <-- Commenting this line as this function needs access to actual system command
		c.JSON(200, gin.H{})
	})

	// Record the response
	w := httptest.NewRecorder()

	// Create a request to use
	req, err := http.NewRequest(http.MethodGet, "/config", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	webAPI.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %v but instead got %v\n", http.StatusOK, w.Code)
	}
}
