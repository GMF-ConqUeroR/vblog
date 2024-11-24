package impl

import (
	"gitee.com/xpigpig/vblog/apps/blog"
	"gitee.com/xpigpig/vblog/conf"
	"gorm.io/gorm"
)

var _ blog.Service = &Impl{}

type Impl struct {
	db *gorm.DB
}

func (i *Impl) Init() error {
	db := conf.Values().MySQL.ORM()
	//db.AutoMigrate(&blog.Blog{})
	i.db = db.Debug()
	return nil
}
