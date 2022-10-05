package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func getDbConnectionString(c *Conf) *string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Db.User, c.Db.Password, c.Db.Host, c.Db.Schema)
	return &dsn
}

// GetDbClient 获取数据库连接对象
func GetDbClient() *gorm.DB {
	once.Do(func() {
		filePath := getConfFilePath()
		dsn := getDbConnectionString(getConfiguration(filePath))

		_db, err := gorm.Open(mysql.Open(*dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
		if err != nil {
			panic(err)
		}

		db = _db
	})

	return db
}
