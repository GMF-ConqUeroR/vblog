package start

import (
	"gitee.com/xpigpig/vblog/apps/blog/api"
	"gitee.com/xpigpig/vblog/apps/blog/impl"
	"gitee.com/xpigpig/vblog/conf"
	"gitee.com/xpigpig/vblog/protocol"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "start",
	Short: "启动后端项目",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		switch configType {
		case "env":
			_, err := conf.LoadConfigFromEnv()
			cobra.CheckErr(err)
		default:
			_, err := conf.LoadConfigFromToml(configFile)
			cobra.CheckErr(err)
		}

		// 加载业务逻辑
		r := gin.Default()

		// 加载blog模块
		blogService := &impl.Impl{}
		err := blogService.Init()
		cobra.CheckErr(err)
		api.NewHandler(blogService).Registry(r)

		// 启动http服务
		httpServer := protocol.NewHttp(r)
		cobra.CheckErr(httpServer.Start())
	},
}

var (
	configType string
	configFile string
)

func init() {
	Cmd.Flags().StringVarP(&configType, "config-type", "t", "file", "加载配置文件的方式")
	Cmd.Flags().StringVarP(&configFile, "file-path", "f", "etc/config.toml", "配置文件路径")
}
