package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpSuccessResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type HttpFailResponse struct {
	Status  string `json:"status"`
	Message string `json:"msg"`
}

func CustomHttpErrorHandler(c *gin.Context) {
	c.Next()

	for _, ginErr := range c.Errors {
		c.JSON(http.StatusInternalServerError, HttpFailResponse{
			Status:  "fail",
			Message: ginErr.Error(),
		})
	}
}
