package impl

import (
	"context"
	"net/http"

	"gitee.com/xpigpig/vblog/apps/blog"
	"gitee.com/xpigpig/vblog/common"
)

func (i *Impl) CreateBlog(ctx context.Context, in *blog.CreateBlogRequest) (*blog.Blog, error) {
	if err := in.Validate(); err != nil {
		return nil, common.NewApiException(http.StatusBadRequest, http.StatusBadRequest, err.Error())
	}
	blogReq := blog.NewBlog(in)
	err := i.db.WithContext(ctx).Save(blogReq).Error
	if err != nil {
		return nil, common.NewApiException(http.StatusInternalServerError, http.StatusInternalServerError, err.Error())
	}

	return blogReq, nil
}

func (i *Impl) DeleteBlog(context.Context, *blog.DeleteBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

func (i *Impl) UpdateBlog(context.Context, *blog.UpdateBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

func (i *Impl) DescribeBlog(context.Context, *blog.DescribeBlogRequest) (*blog.Blog, error) {
	return nil, nil
}

func (i *Impl) QueryBlog(context.Context, *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	return nil, nil
}
