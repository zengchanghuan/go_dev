package models

// User 和数据表结构相同（首字母大写），文件命名建议和数据库表命名相同
type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// TableName 表示配置操作数据库的表名称
func (User) TableName() string {
	return "user"
}
