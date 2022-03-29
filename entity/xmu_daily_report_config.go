package entity

import "gorm.io/gorm"

type contentType string

const (
	STRING   = "string"
	NUMBER   = "number"
	EMAIL    = "email"
	DROPDOWN = "dropdown"
)

type XmuDailyReportConfig struct {
	gorm.Model
	Type   contentType
	Key    string
	Values string
}
