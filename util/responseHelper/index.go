package responseHelper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ok(c *gin.Context, res any) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "SUCCESS",
		"data":    res,
	})
}

func NotFound(c *gin.Context, res any) {
	c.JSONP(http.StatusNotFound, gin.H{
		"status":  "404",
		"message": "NOT_FOUND",
		"error":   res,
	})
}

func InternalServerError(c *gin.Context, res any) {
	c.JSONP(http.StatusInternalServerError, gin.H{
		"status":  "500",
		"message": "INTERNAL_SERVER_ERROR",
		"error":   res,
	})
}

func BadRequest(c *gin.Context, res any) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"status":  "400",
		"message": "BAD_REQUEST",
		"error":   res,
	})
}
