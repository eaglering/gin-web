package models

type DShades struct {
	Id          int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	SkillId     int    `xorm:"not null comment('技能编号') MEDIUMINT(8)"`
	SkillName   string `xorm:"not null comment('技能名称') VARCHAR(20)"`
	OnlineCount int    `xorm:"not null comment('上线总数') INT(10)"`
	UpDown      int    `xorm:"not null comment('每日上线人数涨幅') INT(11)"`
	CreatedAt   int    `xorm:"not null comment('创建时间') INT(10)"`
}
