package Models

import "time"

type User struct {
	User_id   uint `gorm:"primary_key"` // 设置id为主键
	Username  string
	Password  string
	User_type uint
	Gender    uint
	Birthday  time.Time
}

// 表名默认为模型名称的复数，此处手动设置User的表名为`users`
func (User) TableName() string {
	return "user"
}
