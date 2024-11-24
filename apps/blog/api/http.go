package api

import (
	"gitee.com/xpigpig/vblog/apps/blog"

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

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/blog/api/v1/blogs", h.CreateBlog)
}
