package main

import (
	"data-pusher/config"
	"data-pusher/dao"
	"data-pusher/entity"
	"testing"
)

func TestInitDB(t *testing.T) {
	c := config.MustGetConfig("xmu-daily-report")
	configDao := dao.NewXmuDailyReportConfigDao(
		c.MustString("mysql-user"),
		c.MustString("mysql-passwd"),
		c.MustString("mysql-ip"),
		c.MustInt("mysql-port"),
		c.MustString("mysql-db"))
	configDao.DB().Where("1=1").Delete(entity.XmuDailyReportConfig{})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.STRING,
		Name:   "学号",
		Key:    "username",
		Values: "",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.STRING,
		Name:   "密码",
		Key:    "password",
		Values: "",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.STRING,
		Name:   "VPN密码",
		Key:    "password_vpn",
		Values: "",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.EMAIL,
		Name:   "电子邮件",
		Key:    "email",
		Values: "",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.DROPDOWN,
		Name:   "区域",
		Key:    "district",
		Values: "思明区,翔安区",
	})
}
