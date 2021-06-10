package models

type FmtSmsSendLog struct {
	LogId      int64     `gorm:"column:msg_id"`      //  id
	TaskId     string    `gorm:"column:task_id"`     // 任务ID
	TemplateId string    `gorm:"column:template_id"` // 模板ID
	Phone      string    `gorm:"column:phone"`       // 手机号
	Status     int       `gorm:"column:status"`      // 回调状态 0 待通知 1 成功 2 失败
	CreateYw   string    `gorm:"column:create_yw"`   // 创建人
	CreateTime int64     `gorm:"column:create_time"` // 创建时间
	CreateDate LocalDate `gorm:"column:create_date"` // 创建日期
}
