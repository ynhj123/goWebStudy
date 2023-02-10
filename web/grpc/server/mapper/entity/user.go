package entity

import "time"

type User struct {
	Id           int64 `gorm:"primarykey"`
	Accesskey    string
	Username     string
	Password     string
	Icon         string
	Phone        string
	Email        string
	Gender       string
	Birthday     time.Time
	Nation       string
	Introduction string
	CountryCode  string
	DeviceId     string
	OsVersion    string
	DeviceName   string
	FromType     string
	FromId       string
	Badge        string
	Coin         string
	Type         string
	IsActive     string
	CreatedTime  time.Time
	UpdateTime   time.Time
	Occupation   string
}

func (User) TableName() string {
	return "user"
}
