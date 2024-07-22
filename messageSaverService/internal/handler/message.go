package handler

import (
	"encoding/json"
	"message-saver/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getMsgStatistic(c *gin.Context) {
	totalMsg, err := h.services.GetTotalMessages()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	processedMsg, err := h.services.GetProcessedMessages()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"total-messages":     totalMsg,
		"processed-messages": processedMsg,
	})
}

func (h *Handler) saveMessage(c *gin.Context) {

	var message model.MessageInput
	if err := json.NewDecoder(c.Request.Body).Decode(&message); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "json decoding error"})
		return
	}

	if message.Text == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "no input data"})
		return
	}

	msg, err := h.services.Message.SaveMessage(message)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, msg)
}
