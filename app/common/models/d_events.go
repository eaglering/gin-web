package models

type DEvents struct {
	Id          int    `xorm:"not null pk autoincr INT(10)"`
	Title       string `xorm:"not null VARCHAR(32)"`
	Description string `xorm:"not null VARCHAR(400)"`
	Background  string `xorm:"not null VARCHAR(255)"`
	Expand1     string `xorm:"VARCHAR(255)"`
	Expand2     string `xorm:"VARCHAR(255)"`
	StartAt     int    `xorm:"not null INT(10)"`
	EndAt       int    `xorm:"not null INT(10)"`
	Active      int    `xorm:"not null default 0 TINYINT(1)"`
	CreatedAt   int    `xorm:"not null INT(10)"`
}
