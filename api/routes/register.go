package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Register(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/:type/:channel", webhook)
	}
	r.GET("/ping", ping)
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func webhook(c *gin.Context) {
	var r interface{}
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logrus.Infof("%+v", r)
	c.JSON(200, r)
}
