package handler

import (
	"message-saver/internal/service"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(log *logrus.Logger) *gin.Engine {

	gin.SetMode(os.Getenv("GIN_MODE"))
	router := gin.New()

	router.Use(Logger(log), gin.Recovery())

	router.GET("/messages/statistic", h.getMsgStatistic)
	router.POST("/messages", h.saveMessage)

	return router
}

func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		log.WithFields(logrus.Fields{
			"status_code":  c.Writer.Status(),
			"latency_time": latencyTime,
			"client_ip":    c.ClientIP(),
			"method":       c.Request.Method,
			"path":         c.Request.RequestURI,
		}).Info("Request completed")
	}
}
