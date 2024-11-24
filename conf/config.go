package conf

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	orm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	MySQL *MySQL `json:"mysql" toml:"mysql"`
	Http  *Http  `json:"http" toml:"http"`
}

type Http struct {
	Host string `json:"host" toml:"host" env:"HTTP_HOST"`
	Port int    `json:"port" toml:"port" env:"HTTP_PORT"`
}

func NewDefualtHttp() *Http {
	return &Http{
		Host: "127.0.0.1",
		Port: 18080,
	}
}
func (h *Http) Address() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

type MySQL struct {
	Host     string `json:"host" toml:"host" env:"MYSQL_HOST"`
	Port     int    `json:"port" toml:"port" env:"MYSQL_PORT"`
	DB       string `json:"db"   toml:"db" env:"MYSQL_DB"`
	Username string `json:"username" toml:"username" env:"MYSQL_USERNAME"`
	Password string `json:"password" toml:"password" env:"MYSQL_PASSWORD"`

	// 连接池相关配置
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`

	// 面临并发安全
	lock sync.Mutex
	db   *gorm.DB
}

// 返回一个带有默认值的 mysql 对象
func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:     "127.0.0.1",
		Port:     3306,
		DB:       "blog",
		Username: "blog",
		Password: "blog123",
	}
}

func DefaultConfig() *Config {
	return &Config{
		MySQL: NewDefaultMySQL(),
		Http:  NewDefualtHttp(),
	}
}

// Stringer is implemented by any value that has a String method,
// which defines the “native” format for that value.
// The String method is used to print values passed as an operand
// to any format that accepts a string or to an unformatted printer
// such as Print.
//
//	type Stringer interface {
//		String() string
//	}
func (c *Config) String() string {
	res, _ := json.MarshalIndent(c, "", "	")
	return string(res)
}

// 连接池对象
func (m *MySQL) GetConnPool() (*sql.DB, error) {
	// multiStatements 让db 可以执行多个语句 select; insert;
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&multiStatements=true",
		m.Username, m.Password, m.Host, m.Port, m.DB)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("数据库连接失败：%s", err.Error())
	}

	// 配置连接池
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleTime != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping db error:%s", err.Error())
	}

	return db, nil
}

// 这里的*gorm.DB 是一个单列实例
func (m *MySQL) ORM() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.db == nil {
		pdb, err := m.GetConnPool()
		if err != nil {
			panic(err)
		}

		// 使用连接池 初始化 ORM DB 对象
		m.db, err = gorm.Open(orm_mysql.New(orm_mysql.Config{
			Conn: pdb,
		}), &gorm.Config{
			// 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
			PrepareStmt: true,
			// 对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行。
			// 但这会降低性能，如果没有这方面的要求，您可以在初始化时禁用它，这将获得大约 30%+ 性能提升
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic(err)
		}
	}

	return m.db
}
