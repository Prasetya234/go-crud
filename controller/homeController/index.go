package homeController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Server lab running at "+time.Now().String())
}
