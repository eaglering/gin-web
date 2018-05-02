package models

type DCompanis struct {
	Id          int64  `xorm:"pk autoincr BIGINT(20)"`
	Name        string `xorm:"not null comment('企业名称') VARCHAR(32)"`
	Description string `xorm:"not null comment('企业描述') VARCHAR(255)"`
	Uid         int64  `xorm:"not null comment('法人编号') BIGINT(20)"`
	CreatedAt   int    `xorm:"not null comment('成立时间') INT(10)"`
}
