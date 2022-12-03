# gohub

# 所有路由

请求方法 |	API 地址	 |	说明
---- |	----  |	-----
POST |	/api/v1/auth/login/using-phone |	短信 + 手机号登录
POST |	/api/v1/auth/login/using-password	| 手机号、用户名、邮箱 + 密码
POST | /api/v1/auth/login/refresh-token	| 刷下 Token
POST |	/api/v1/auth/password-reset/using-email	|邮件密码重置
POST	|/api/v1/auth/password-reset/using-phone	|短信验证码密码重置
POST	|/api/v1/auth/signup/using-phone	|使用手机号注册
POST	|/api/v1/auth/signup/using-email	|使用邮箱注册
POST	|/api/v1/auth/signup/phone/exist	|手机号是否已注册
POST	|/api/v1/auth/signup/email/exist	|email 是否已支持
POST	|/api/v1/auth/verify-codes/phone	|发送短信验证码
POST	|/api/v1/auth/verify-codes/email	|发送邮件验证码
POST	|/api/v1/auth/verify-codes/captcha	|获取图片验证码
GET	|/api/v1/user	|获取当前用户
GET	|/api/v1/users	|用户列表
PUT	|/api/v1/users	|修改个人资料
PUT	|/api/v1/users/email	|修改邮箱
PUT	|/api/v1/users/phone	|修改手机号
PUT	|/api/v1/users/password	|修改密码
PUT	|/api/v1/users/avatar	|上传头像
GET	|/api/v1/categories	|分类列表
POST |/api/v1/categories	|创建分类
PUT	|/api/v1/categories/:id	|更新分类
DELETE|	/api/v1/categories/:id	|删除分类
GET	|/api/v1/topics	|话题列表
POST	|/api/v1/topics	|创建话题
PUT	|/api/v1/topics/:id	|更新话题
DELETE	|/api/v1/topics/:id	|删除话题
GET	|/api/v1/topics/:id	|获取话题
GET|	/api/v1/links	|友情链接列表

# 第三方依赖
- gin —— 路由、路由组、中间件
- zap —— 高性能日志方案
- gorm —— ORM 数据操作
- cobra —— 命令行结构
- viper —— 配置信息
- cast —— 类型转换
- redis —— Redis 操作
- jwt —— JWT 操作
- base64Captcha —— 图片验证码
- govalidator —— 请求验证器
- limiter —— 限流器
- email —— SMTP 邮件发送
- aliyun-communicate —— 发送阿里云短信
- ansi —— 终端高亮输出
- strcase —— 字符串大小写操作
- pluralize —— 英文字符单数复数处理
- faker —— 假数据填充
- imaging —— 图片裁切

# 自定义的包
- app —— 应用对象
- auth —— 用户授权
- cache —— 缓存
- captcha —— 图片验证码
- config —— 配置信息
- console —— 终端
- database —— 数据库操作
- file —— 文件处理
- hash —— 哈希
- helpers —— 辅助方法
- jwt —— JWT 认证
- limiter —— API 限流
- logger —— 日志记录
- mail —— 邮件发送
- migrate —— 数据库迁移
- paginator —— 分页器
- redis —— Redis 数据库操作
- response —— 响应处理
- seed —— 数据填充
- sms —— 发送短信
- str —— 字符串处理
- verifycode —— 数字验证码

# 代码行数

```go
$ gocloc .
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
------------------------------------------------------------------------------- 
Go                             122           1133            969           4672 
TOML                             1              5             24             28 
------------------------------------------------------------------------------- 
TOTAL                          123           1138            993           4700 
------------------------------------------------------------------------------- 
```

# 所有命令
```go
$ go run main.go -h
Default will run "serve" command, you can use "-h" flag to see all subcommands

Usage:
  Gohub [command]

Available Commands:
  cache       Cache management
  completion  Generate the autocompletion script for the specified shell        
  help        Help about any command
  key         Generate App Key, will print the generated Key
  make        Generate file and code
  migrate     Run database migration
  play        Likes the Go Playground, but running at our application context   
  seed        Insert fake data to the database
  server      Start web server

Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testin
g file
  -h, --help         help for Gohub

Use "Gohub [command] --help" for more information about a command.
```

make 命令：
```go
$ go run main.go make -h      
Generate file and code

Usage:
  Gohub make [command]

Available Commands:
  apicontroller Create api controller，exmaple: make apicontroller v1/user      
  cmd           Create a command, should be snake_case, exmaple: make cmd buckup
_database
  factory       Create model's factory file, example: make factory user
  migration     Create a migration file, example: make migration add_users_table
  model         Crate model file, example: make model user
  policy        Create policy file, example: make policy user
  request       Create request file, example make request user
  seeder        Create seeder file, example:  make seeder user

Flags:
  -h, --help   help for make

Global Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testin
g file

Use "Gohub make [command] --help" for more information about a command.
```

migrate 命令：
```go
$ go run main.go migrate -h   
Run database migration

Usage:
  Gohub migrate [command]

Available Commands:
  down        Reverse the up command
  fresh       Drop all tables and re-run all migrations
  refresh     Reset and re-run all migrations
  reset       Rollback all database migrations
  up          Run unmigrated migrations

Flags:
  -h, --help   help for migrate

Global Flags:
  -e, --env string   load .env file, example: --env=testing will use .env.testin
g file

Use "Gohub migrate [command] --help" for more information about a command. 
```


# 自动化生成CRUD接口
## 1. 创建模型
```shell
$ go run main.go make model category
```

修改下 category_model.go 文件里的模型定义

app/models/category/category_model.go

```go
.
.
.
type Category struct {
    models.BaseModel

    Name        string `json:"name,omitempty"`
    Description string `json:"description,omitempty"`

    models.CommonTimestampsField
}
.
.
.
```

## 2. 创建迁移（数据表）

```shell
$ go run main.go make migration add_categories_table category
```

这里有两个参数一个是生成表，一个是表名

去database/migrations目录下打开生成的 migration 文件，定制表结构：
```go
.
.
.
func init() {

    type Category struct {
        models.BaseModel

        Name        string `gorm:"type:varchar(255);not null;index"`
        Description string `gorm:"type:varchar(255);default:null"`

        models.CommonTimestampsField
    }

    up := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.AutoMigrate(&Category{})
    }

    down := func(migrator gorm.Migrator, DB *sql.DB) {
        migrator.DropTable(&Category{})
    }
.
.
.
```

## 3. 执行迁移 生成数据表
```shell
$ go run main.go migrate up
```
![img.png](public/uploads/image/img.png)
## 4. 生成数据验证 request 文件
```shell
$ go run main.go make request category
```
修改请求数据结构，以及验证规则和错误：

app/requests/category_request.go
```go
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	Name        string `valid:"name" json:"name"`
	Description string `valid:"description" json:"description,omitempty"`
}

func CategorySave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
		"description": []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 8 个字",
			"not_exists:名称已存在",
		},
		"description": []string{
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}
```

## 5. 自定义验证规则
我们可以使用自定义验证规则来完善数据验证功能, 这里我们底层使用的验证器 govalidator 虽然支持 min 和 max 来设置字符串长度规则，但是不适用于中文字符串。

所以上面我们使用了 min_cn 和 max_cn 的自定义规则，现在来创建这两个规则：

app/requests/validators/custom_rules.go

```go
.
.
.
// 此方法会在初始化时执行，注册自定义表单验证规则
func init() {
    .
    .
    .
    // max_cn:8 中文长度设定不超过 8
    govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
        valLength := utf8.RuneCountInString(value.(string))
        l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
        if valLength > l {
            // 如果有自定义错误消息的话，使用自定义消息
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("长度不能超过 %d 个字", l)
        }
        return nil
    })

    // min_cn:2 中文长度设定不小于 2
    govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
        valLength := utf8.RuneCountInString(value.(string))
        l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
        if valLength < l {
            // 如果有自定义错误消息的话，使用自定义消息
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("长度需大于 %d 个字", l)
        }
        return nil
    })
}
```

## 6. 生成控制器 controller 文件
```shell
$ go run main.go make apicontroller v1/category
```
app/http/controllers/api/v1 目录下生成的 categories_controller.go 里有很多内容。


## 7. 创建工厂, 生成假数据

首先我们来填充一些数据，方便测试分页。

先来创建分类工厂：
```shell
$ go run main.go make factory category
[database/factories/category_factory.go] created.
```
修改内容如下；

database/factories/category_factory.go

```go
.
.
.
func MakeCategories(count int) []category.Category {

    var objs []category.Category

    // 设置唯一性，如 Category 模型的某个字段需要唯一，即可取消注释
    faker.SetGenerateUniqueValues(true)

    for i := 0; i < count; i++ {
        categoryModel := category.Category{
            Name:        faker.Username(),
            Description: faker.Sentence(),
        }
        objs = append(objs, categoryModel)
    }

    return objs
}
```

> 因为分类名称要保持唯一，所以取消了上面的 faker.SetGenerateUniqueValues(true) 的注释。

## 8. 生成 Seed 文件

```go
$ go run main.go make seeder category
[database/seeders/categories_seeder.go] created.
```

文件不用修改

## 9. 填充数据

我们只需要填充 SeedCategoriesTable 即可：
```go
$ go run main.go seed SeedCategoriesTable
Table [categories] 10 rows seeded
```


## 10. 修改控制器文件里的Index方法 (如果数据需要分页)
app/http/controllers/api/v1/categories_controller.go
```go
.
.
.

func (ctrl *CategoriesController) Index(c *gin.Context) {
    request := requests.PaginationRequest{}
    if ok := requests.Validate(c, &request, requests.Pagination); !ok {
        return
    }

	// category 修改为 你的表名
    data, pager := category.Paginate(c, 10)
    response.JSON(c, gin.H{
        "data":  data,
        "pager": pager,
    })
}
```


## 11. 注册路由

在 router/api.go 文件里注册你的刚刚生成的路由

```go
.
.
.
cgc := new(controllers.CategoriesController)
cgcGroup := v1.Group("/categories")
{
cgcGroup.GET("", cgc.Index)
cgcGroup.POST("", middlewares.AuthJWT(), cgc.Store)
cgcGroup.PUT("/:id", middlewares.AuthJWT(), cgc.Update)
cgcGroup.DELETE("/:id", middlewares.AuthJWT(), cgc.Delete)
}
```

需要 Token 验证, 注册或者登录获得 token

![select.png](public/uploads/image/select.png)
![create.png](public/uploads/image/create.png)
![update.png](public/uploads/image/update.png)
![delete.png](public/uploads/image/delete.png)
