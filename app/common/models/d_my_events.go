package models

type DMyEvents struct {
	Id        int64 `xorm:"pk autoincr BIGINT(20)"`
	UserId    int64 `xorm:"not null comment('用户编号') BIGINT(20)"`
	EventId   int   `xorm:"not null comment('事件编号') INT(10)"`
	CreatedAt int   `xorm:"not null comment('创建时间') INT(10)"`
}
