package blog

import (
	"time"

	"gitee.com/xpigpig/vblog/common"
)

type Meta struct {
	ID        int64 `json:"id"`
	CreateAt  int64 `json:"create_at"`
	UpdateAt  int64 `json:"update_at"`
	PublishAt int64 `json:"publish_at"`
}

func NewMeta() *Meta {
	return &Meta{
		CreateAt: time.Now().Unix(),
	}
}

type CreateBlogRequest struct {
	Title   string            `json:"title" gorm:"column:title" validate:"required"`
	Author  string            `json:"author" gorm:"column:author" validate:"required"`
	Content string            `json:"content" validate:"required"`
	Tags    map[string]string `json:"tags" gorm:"serializer:json"`
	Status  STATUS            `json:"status"`
}

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags:   map[string]string{},
		Status: STATUS_DRAFT,
	}
}

// 检查用户参数是否合法
func (cbr *CreateBlogRequest) Validate() error {
	return common.Validate(cbr)
}

type Blog struct {
	*Meta
	*CreateBlogRequest
}

func NewBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		Meta:              NewMeta(),
		CreateBlogRequest: req,
	}
}

//	type Tabler interface {
//		TableName() string
//	}
//
// 自定义定义gorm 存入数据时表的名称
func (b *Blog) TableName() string {
	return "blog"
}

type BlogSet struct {
	Items []*Blog
}

type QueryBlogRequest struct {
}

type DescribeBlogRequest struct {
}

type UpdateBlogRequest struct {
	*CreateBlogRequest
}

type DeleteBlogRequest struct {
}
