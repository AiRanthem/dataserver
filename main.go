package main

import (
	"data-pusher/config"
	"data-pusher/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	xmuDailyReportConfigHandler := service.NewXmuDailyReportConfigHandler(config.MustGetConfig("xmu-daily-report"))

	r := gin.Default()
	r.GET("/api/v1/dataserver/xmu-daily-report-config", xmuDailyReportConfigHandler.Handle)

	log.Infoln("application running")
	err := r.Run(":8080")
	if err != nil {
		log.WithError(err).Errorln("fail to run application")
	}
}
