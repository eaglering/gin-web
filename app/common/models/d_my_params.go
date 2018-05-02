package models

type DMyParams struct {
	Id         int64 `xorm:"pk comment('用户编号') BIGINT(20)"`
	Money      int   `xorm:"not null comment('金钱（单位分）') INT(10)"`
	Healthy    int   `xorm:"not null comment('健康值') INT(10)"`
	Charm      int   `xorm:"not null comment('魅力值') INT(10)"`
	KeepOnline int   `xorm:"not null comment('保持在线的时间') INT(10)"`
	AtPost     int   `xorm:"not null default 0 comment('离职状态') TINYINT(1)"`
	CreatedAt  int   `xorm:"not null comment('开始工作的时间') INT(10)"`
}
