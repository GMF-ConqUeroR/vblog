package impl

import (
	"gitee.com/xpigpig/vblog/apps/blog"
	"gitee.com/xpigpig/vblog/conf"
	"gitee.com/xpigpig/vblog/ioc"
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

func (i *Impl) Name() string {
	return blog.AppName
}

// 把这个对象托管到Ioc
// 通过Import的方式，自动调用包的init方法，把需要托管的对象注册到ioc中
// import _ "gitee.com/go-course/go11/vblog/apps/blog/impl"
func init() {
	ioc.RegistryController(&Impl{})
}
