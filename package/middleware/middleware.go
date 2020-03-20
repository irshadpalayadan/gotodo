package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GlobalMiddleware() gin.HandlerFunc {

	logrus.Info("added global middleware")

	return func(c *gin.Context) {
		logrus.Info("middleware called before the api call")
		c.Next()
	}
}
