package blog

import "context"

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*Blog, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
}
