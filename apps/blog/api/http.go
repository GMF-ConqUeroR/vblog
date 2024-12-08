package api

import (
	"gitee.com/xpigpig/vblog/apps/blog"
	"gitee.com/xpigpig/vblog/ioc"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc blog.Service
}

func NewHandler(svc blog.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) Init() error {
	h.svc = ioc.GetController(blog.AppName).(blog.Service)
	return nil
}

func (h *Handler) Name() string {
	return blog.AppName
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/blog/api/v1/blogs", h.CreateBlog)
}

func init() {
	ioc.RegisryApi(&Handler{})
}
