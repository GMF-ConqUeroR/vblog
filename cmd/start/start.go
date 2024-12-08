package start

import (
	"gitee.com/xpigpig/vblog/conf"
	"gitee.com/xpigpig/vblog/ioc"
	"gitee.com/xpigpig/vblog/logger"
	"gitee.com/xpigpig/vblog/protocol"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	// 加载所有的业务实现和接口实现, init
	// 统一到apps包下面, 然后只导入一次,
	// apps: 代表当前已经开发完成的业务功能实例和接口实例
	// 把所有的业务模块全部加载
	_ "gitee.com/xpigpig/vblog/apps"
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

		// 手动加载blog模块
		// blogService := &impl.Impl{}
		// err := blogService.Init()
		// cobra.CheckErr(err)
		// api.NewHandler(blogService).Registry(r)

		// 通过ioc自动加载各个模块
		logger.Log().Debug().Msgf("加载模块%v成功", ioc.ListContrillers())
		logger.Log().Debug().Msgf("加载接口%v成功", ioc.ListApi())
		err := ioc.InitAllControllers()
		cobra.CheckErr(err)
		err = ioc.InitAllApi(r)
		cobra.CheckErr(err)

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
