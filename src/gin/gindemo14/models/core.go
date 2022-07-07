package models

//https://gorm.io/zh_CN/docs/connecting_to_the_database.html
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

//与数据库建立连接，Gorm 连接数据库
func init() {
	/*
			nslookup  localhost
			上面命令可以查看当前主机ip
		root 用户名
		optix.123
		gin dbName 不是数据库表 http://127.0.0.1/
	*/
	//dsn := "root:optix.123@tcp(192.168.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := "root:optix.123@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
