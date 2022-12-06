package seeders

import (
	"fmt"
	"github.com/Gopherlinzy/gohub/database/factories"
	"github.com/Gopherlinzy/gohub/pkg/console"
	"github.com/Gopherlinzy/gohub/pkg/database"
	"github.com/Gopherlinzy/gohub/pkg/logger"
	"github.com/Gopherlinzy/gohub/pkg/seed"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

func init() {

	// 添加 Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		// 创建 10 个用户
		users, rules := factories.MakeUsers(10)

		// 批量创建用户
		result := db.Table("users").Create(&users)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}
		// 用户创建成功则生成 casbin 对应角色
		database.Gohub_DB.Model(&gormadapter.CasbinRule{}).Create(rules)
		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
