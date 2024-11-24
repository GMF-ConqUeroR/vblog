package api

import (
	"net/http"

	"gitee.com/xpigpig/vblog/apps/blog"
	"gitee.com/xpigpig/vblog/common"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateBlog(c *gin.Context) {
	in := blog.NewCreateBlogRequest()
	err := c.BindJSON(in)
	if err != nil {
		// 异常处理
		common.RespFail(c, err)
		return
	}

	ins, err := h.svc.CreateBlog(c.Request.Context(), in)
	if err != nil {
		common.RespFail(c, err)
		return
	}
	c.JSON(http.StatusOK, ins)
}
