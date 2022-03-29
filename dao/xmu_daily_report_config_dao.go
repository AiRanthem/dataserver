package dao

import "data-pusher/entity"

type XmuDailyReportConfigDao struct {
	*dbAccess
}

func (x *XmuDailyReportConfigDao) Create(e *entity.XmuDailyReportConfig) {
	x.db.Create(e)
}

func (x *XmuDailyReportConfigDao) All() []*entity.XmuDailyReportConfig {
	var e []*entity.XmuDailyReportConfig
	x.db.Find(&e)
	return e
}

func NewXmuDailyReportConfigDao(user string, pass string, ip string, port int, dbName string) *XmuDailyReportConfigDao {
	access := MySQL(user, pass, ip, port, dbName, &entity.XmuDailyReportConfig{})
	return &XmuDailyReportConfigDao{access}
}
