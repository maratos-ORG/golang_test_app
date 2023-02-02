package backend

import (
	"golang_test_app/internal/config"
	log "golang_test_app/internal/logging"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	// Set the log format
	return func(c *gin.Context) {
		// Starting time
		startTime := time.Now().UnixMilli()
		// Processing request
		c.Next()
		// End Time
		endTime := time.Now().UnixMilli()
		// execution time
		latencyTime := endTime - startTime
		// Request method
		reqMethod := c.Request.Method
		// Request route
		reqURI := c.Request.RequestURI
		// status code
		statusCode := c.Writer.Status()
		// Request IP
		clientIP := c.ClientIP()

		log.WithFields(map[string]interface{}{
			"caller":     clientIP,
			"status":     statusCode,
			"latency ms": latencyTime,
			"method":     reqMethod,
			"url":        reqURI}).Info("API call")
	}
}

func RunBackend(params *config.BackendParameters) {

	gin.SetMode(gin.ReleaseMode)
	webAPI := gin.New()
	webAPI.Use(Logger())
	webAPI.Use(gin.Recovery())

	webAPI.GET("/config", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"config": params,
		})
	})
	err := webAPI.Run(":" + *params.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal("unable to run webapi: %s", err)
	}
}
