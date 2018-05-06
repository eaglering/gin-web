package models

type Users struct {
	Id        int    `xorm:"not null pk autoincr INT(10)"`
	Name      string `xorm:"not null VARCHAR(32)"`
	CreatedAt int    `xorm:"not null INT(10) created"`
}
