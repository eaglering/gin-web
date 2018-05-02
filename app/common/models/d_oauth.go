package models

type DOauth struct {
	Id        int64  `xorm:"pk autoincr BIGINT(20)"`
	Platform  int    `xorm:"not null comment('来源平台') unique(platform) TINYINT(3)"`
	Openid    string `xorm:"not null comment('唯一ID') unique(platform) VARCHAR(32)"`
	UserId    int    `xorm:"not null comment('用户ID') INT(10)"`
	CreatedAt int    `xorm:"not null comment('创建时间') INT(10)"`
}
