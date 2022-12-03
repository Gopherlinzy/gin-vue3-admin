package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Gopherlinzy/gohub/pkg/configYaml"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"sync"
)

// Gohub_DB 全局 DB 对象
var (
	Gohub_DB     *gorm.DB
	SqlDB        *sql.DB
	Gohub_DBList map[string]*gorm.DB

	lock sync.RWMutex
)

// Connection 连接数据库
func Connection(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	// 使用 gorm.Open 连接数据
	var err error
	Gohub_DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	// 处理错误
	if err != nil {
		fmt.Println(err.Error())
	}

	SqlDB, err = Gohub_DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return Gohub_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := Gohub_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}

func CurrentDatabase() (dbname string) {
	dbname = Gohub_DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	var err error
	switch configYaml.Gohub_Config.App.DbType {
	case "mysql":
		err = deleteMySQLTables()
	case "sqlite":
		err = deleteAllSqliteTables()
	default:
		panic(errors.New("database connection not supported"))
	}
	return err
}

func deleteMySQLTables() error {
	dbname := CurrentDatabase()
	var tables []string

	err := Gohub_DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}
	// 暂时关闭外键检测
	Gohub_DB.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		if err := Gohub_DB.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	return nil
}

func deleteAllSqliteTables() error {
	tables := []string{}

	// 读取所有数据表
	err := Gohub_DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'").Error
	if err != nil {
		return err
	}

	// 删除所有表
	for _, table := range tables {
		err := Gohub_DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

// TableName 获取表名称
func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: Gohub_DB}
	stmt.Parse(obj)
	return stmt.Schema.Table
}
