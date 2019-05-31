package http

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/LIYINGZHEN/ginexample/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(jwt jwt.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID, err := c.Cookie("sessionID")
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		claims, err := jwt.ValidateToken(sessionID)
		if err != nil {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("admin", strconv.FormatBool(claims.Admin))
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func Logger(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		path := c.Request.URL.Path
		method := c.Request.Method
		ip := c.ClientIP()

		c.Next()

		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		l.Printf("%-13s | %-12s | %s %s", latency, ip, method, path)
	}
}
