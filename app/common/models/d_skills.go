package models

type DSkills struct {
	Id         int    `xorm:"not null pk autoincr MEDIUMINT(8)"`
	Name       string `xorm:"not null comment('编程语言') VARCHAR(20)"`
	Tags       string `xorm:"not null comment('标签数组') TEXT"`
	Code       string `xorm:"not null comment('代码') TEXT"`
	CreatedAt  int    `xorm:"not null comment('添加时间') INT(10)"`
	SharePrice int64  `xorm:"not null default 0 comment('股价') BIGINT(20)"`
}
