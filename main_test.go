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
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.DROPDOWN,
		Name:   "是否在校",
		Key:    "in",
		Values: "在校,不在校",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.DROPDOWN,
		Name:   "校区（在校时生效）",
		Key:    "campus",
		Values: "思明校区,翔安校区",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.DROPDOWN,
		Name:   "是否住在校内（在校时生效）",
		Key:    "stay",
		Values: "住校内,不在校内",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.DROPDOWN,
		Name:   "是否住在校内学生宿舍（住校内时生效）",
		Key:    "livedorm",
		Values: "住校内学生宿舍,住校内非学生宿舍",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.STRING,
		Name:   "楼栋（住校内学生宿舍时生效，请保证填写与健康打卡页面上的一致）",
		Key:    "building",
		Values: "",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.STRING,
		Name:   "房间号",
		Key:    "room",
		Values: "",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.DROPDOWN,
		Name:   "是否家在厦门（不住校内时生效）",
		Key:    "livexm",
		Values: "是,否",
	})
	configDao.Create(&entity.XmuDailyReportConfig{
		Type:   entity.STRING,
		Name:   "校外住址（不住校内时生效）",
		Key:    "address",
		Values: "",
	})
}
