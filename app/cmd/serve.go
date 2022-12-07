package cmd

import (
	"fmt"
	"github.com/Gopherlinzy/gin-vue3-admin/bootstrap"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// CmdServe represents the available web sub-command.
var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	//fmt.Println(global.GohubConfig.App)
	// 初始化 gin 实例
	router := gin.New()

	// 初始化 DB
	bootstrap.SetupDB()

	// 初始化 Redis
	bootstrap.SetupRedis()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + configYaml.Gohub_Config.App.Port)
	if err != nil {
		// 错误处理, 端口被占用了或者其他错误
		fmt.Println(err.Error())
	}
}
