package service

import (
	"data-pusher/config"
	"data-pusher/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type XmuDailyReportConfigHandler struct {
	xmuDailyReportConfigDao *dao.XmuDailyReportConfigDao
}

func (h *XmuDailyReportConfigHandler) Handle(ctx *gin.Context) {
	all := h.xmuDailyReportConfigDao.All()
	var data []gin.H
	for _, entity := range all {
		data = append(data, gin.H{
			"type":   entity.Type,
			"key":    entity.Key,
			"values": entity.Values,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error": 0,
		"data":  data,
	})
}

func NewXmuDailyReportConfigHandler(c config.Config) *XmuDailyReportConfigHandler {
	return &XmuDailyReportConfigHandler{
		xmuDailyReportConfigDao: dao.NewXmuDailyReportConfigDao(
			c.MustString("mysql-user"),
			c.MustString("mysql-passwd"),
			c.MustString("mysql-ip"),
			c.MustInt("mysql-port"),
			c.MustString("mysql-db")),
	}
}
