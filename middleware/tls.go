package middleware

import (
	"data-pusher/enums"
	"data-pusher/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/secure"
	"net/http"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New()
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			log.WithError(err).Errorln("tls error")
			c.JSON(http.StatusOK, utils.Response(enums.StatusTlsError, err))
			return
		}

		c.Next()
	}
}
