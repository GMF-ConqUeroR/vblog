package protocol

import (
	"context"
	"net/http"
	"time"

	"gitee.com/xpigpig/vblog/conf"
	"github.com/gin-gonic/gin"
)

type Http struct {
	server *http.Server
}

func NewHttp(r *gin.Engine) *Http {
	return &Http{
		server: &http.Server{
			Handler:      r,
			Addr:         conf.Values().Http.Address(),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}
}

func (h *Http) Start() error {
	return h.server.ListenAndServe()
}

func (h *Http) Stop(ctx context.Context) {
	// 服务的优雅关闭, 先关闭监听,新的请求就进不来, 等待老的请求 处理完成
	// 自己介绍来自操作系统的信号量 来决定 你的服务是否需要关闭
	h.server.Shutdown(ctx)
}
