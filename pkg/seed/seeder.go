// Package seed 处理数据库填充相关逻辑
package seed

import (
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/console"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"gorm.io/gorm"
)

// 存放所有 Seeder
var seeders []Seeder

// 按顺序执行的 Seeder 数组
// 支持一些必须按顺序执行的 seeder，例如 topic 创建的时必须依赖于 user,
// 所以 TopicSeeder 应该在 UserSeeder 后执行
var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

// Seeder 对应每一个 database/seeders 目录下的 Seeder 文件
type Seeder struct {
	Func SeederFunc
	Name string
}

// Add 注册到 seeders 数组中
func Add(name string, fc SeederFunc) {
	seeders = append(seeders, Seeder{
		Func: fc,
		Name: name,
	})
}

// SetRunOrder 设置『按顺序执行的 Seeder 数组』
func SetRunOrder(name []string) {
	orderedSeederNames = name
}

// RunSeeder 运行单个 Seeder
func RunSeeder(name string) {
	for _, sdr := range seeders {
		if sdr.Name == name {
			sdr.Func(database.Gohub_DB)
			break
		}
	}
}

// RunAll 运行所有 Seeder
func RunAll() {
	// 先运行 ordered 的
	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Running Ordered Seeder: " + sdr.Name)
			executed[name] = name
			sdr.Func(database.Gohub_DB)
		}
	}

	// 再运行剩下的
	for _, sdr := range seeders {
		// 过滤已运行
		if _, ok := executed[sdr.Name]; !ok {
			console.Warning("Running Seeder: " + sdr.Name)
			sdr.Func(database.Gohub_DB)
		}
	}
}

// GetSeeder 通过名称来获取 Seeder 对象
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if sdr.Name == name {
			return sdr
		}
	}
	return Seeder{}
}
