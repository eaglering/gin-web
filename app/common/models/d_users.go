package models

type DUsers struct {
	Id           int64  `xorm:"pk autoincr BIGINT(20)"`
	Mobile       string `xorm:"comment('手机号') unique VARCHAR(11)"`
	Nickname     string `xorm:"comment('昵称') VARCHAR(20)"`
	Sex          int    `xorm:"not null default 1 comment('性别') TINYINT(1)"`
	Country      string `xorm:"comment('国家') VARCHAR(20)"`
	Province     string `xorm:"comment('省份') VARCHAR(20)"`
	City         string `xorm:"comment('城市') VARCHAR(20)"`
	Language     string `xorm:"comment('语言') VARCHAR(20)"`
	Headimageurl string `xorm:"comment('头像') VARCHAR(400)"`
	Realname     string `xorm:"comment('姓名') VARCHAR(20)"`
	Idcard       string `xorm:"comment('身份证') VARCHAR(18)"`
	CreatedAt    int    `xorm:"not null comment('创建时间') INT(10)"`
}
