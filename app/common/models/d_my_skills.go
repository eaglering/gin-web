package models

type DMySkills struct {
	Id        int64  `xorm:"pk autoincr BIGINT(20)"`
	UserId    int64  `xorm:"not null comment('用户编号') unique(user_id) BIGINT(20)"`
	SkillId   int    `xorm:"not null comment('技能ID') unique(user_id) MEDIUMINT(8)"`
	SkillName string `xorm:"not null comment('技能名称') VARCHAR(20)"`
	CreatedAt int    `xorm:"not null comment('技能掌握时间') INT(10)"`
}
