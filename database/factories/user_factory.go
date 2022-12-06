// Package factories 存放工厂方法
package factories

import (
	"github.com/Gopherlinzy/gohub/app/models/user"
	"github.com/Gopherlinzy/gohub/pkg/helpers"
	"github.com/bxcodec/faker/v4"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func MakeUsers(times int) ([]user.User, []gormadapter.CasbinRule) {

	var objs []user.User
	var rules []gormadapter.CasbinRule

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:         faker.Username(),
			Email:        faker.Email(),
			Phone:        helpers.RandomNumber(11),
			Password:     "123456",
			City:         helpers.RandomString(5),
			Introduction: faker.Sentence(),
			Status: func(x int) bool {
				if x == 1 {
					return true
				}
				return false
			}(helpers.RandomInt(2)),
			RoleName: "user",
		}
		rule := gormadapter.CasbinRule{
			Ptype: "g", V0: model.Name, V1: "user",
		}
		rules = append(rules, rule)
		objs = append(objs, model)
	}

	return objs, rules
}
