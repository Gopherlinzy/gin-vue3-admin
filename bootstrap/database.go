package bootstrap

import (
	"errors"
	"fmt"
	"github.com/Gopherlinzy/gohub/pkg/configYaml"
	"github.com/Gopherlinzy/gohub/pkg/database"
	"github.com/Gopherlinzy/gohub/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	var dbConfig gorm.Dialector
	switch configYaml.Gohub_Config.App.DbType {
	case "mysql":
		// 构建 DSN
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			configYaml.Gohub_Config.MySQL.Username,
			configYaml.Gohub_Config.MySQL.Password,
			configYaml.Gohub_Config.MySQL.Host,
			configYaml.Gohub_Config.MySQL.Port,
			configYaml.Gohub_Config.MySQL.Database,
			configYaml.Gohub_Config.MySQL.Charset,
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		panic(errors.New("database connection not supported"))
	}
	// 连接数据库，并设置 GORM 的日志模式
	database.Connection(dbConfig, logger.NewGormLogger())

	// 设置最大连接数
	database.SqlDB.SetMaxOpenConns(configYaml.Gohub_Config.MySQL.MaxOpenConnections)
	// 设置最大空闲连接数
	database.SqlDB.SetMaxIdleConns(configYaml.Gohub_Config.MySQL.MaxIdleConections)
	// 设置每个链接的过期时间
	database.SqlDB.SetConnMaxLifetime(time.Duration(configYaml.Gohub_Config.MySQL.MaxLifeSeconds) * time.Second)
}
