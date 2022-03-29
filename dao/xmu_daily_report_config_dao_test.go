package dao

import (
	"data-pusher/entity"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestGet(t *testing.T) {
	dao := NewXmuDailyReportConfigDao("root", "toor", "127.0.0.1", 3306, "test_pusher")
	dao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.STRING,
		Key:    "username",
		Values: "",
	})
	all := dao.All()
	for _, config := range all {
		log.WithField("data", config).Infoln("ok")
	}
}
