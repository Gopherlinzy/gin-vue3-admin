package main

import (
	"fmt"
	"github.com/Gopherlinzy/gin-vue3-admin/app/cmd"
	"github.com/Gopherlinzy/gin-vue3-admin/app/cmd/make"
	"github.com/Gopherlinzy/gin-vue3-admin/bootstrap"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/console"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   "Gohub",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			configYaml.InitConfig(cmd.Yaml)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
			bootstrap.SetupCache()

			// 初始化数据库数据
			bootstrap.IntizationData()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,

		make.CmdMake,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--yaml
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

}
