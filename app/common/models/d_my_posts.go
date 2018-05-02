package models

type DMyPosts struct {
	Id          int64  `xorm:"pk autoincr BIGINT(20)"`
	UserId      int64  `xorm:"not null comment('用户编号') unique(user_id) BIGINT(20)"`
	CompanyId   int64  `xorm:"not null comment('企业编号') unique(user_id) BIGINT(20)"`
	CompanyName string `xorm:"not null comment('企业名称') VARCHAR(32)"`
	PostId      int64  `xorm:"not null comment('岗位编号') BIGINT(20)"`
	SkillId     int    `xorm:"not null comment('技能编号') MEDIUMINT(8)"`
	SkillName   string `xorm:"not null comment('技能名称') VARCHAR(20)"`
	CreatedAt   int    `xorm:"not null comment('申请时间') INT(10)"`
	StartAt     int    `xorm:"not null comment('入职时间') INT(10)"`
	EndAt       int    `xorm:"not null comment('离职时间') INT(10)"`
}
