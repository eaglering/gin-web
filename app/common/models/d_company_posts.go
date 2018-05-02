package models

type DCompanyPosts struct {
	Id        int64  `xorm:"pk autoincr BIGINT(20)"`
	CompanyId int64  `xorm:"not null comment('企业编号') BIGINT(20)"`
	Duty      string `xorm:"not null comment('岗位职责') VARCHAR(400)"`
	Needle    string `xorm:"not null comment('任职需求') VARCHAR(400)"`
	Number    int    `xorm:"not null comment('招聘数量') MEDIUMINT(8)"`
	SkillId   int    `xorm:"not null comment('技能编号') MEDIUMINT(8)"`
	SkillName string `xorm:"not null comment('技能名称') VARCHAR(20)"`
	Rest      int    `xorm:"not null comment('剩余岗位') MEDIUMINT(8)"`
	CreatedAt int    `xorm:"not null comment('发布日期') INT(10)"`
}
