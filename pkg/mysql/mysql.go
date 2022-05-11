package mysql

import (
	"fmt"
	"strconv"

	"github.com/yanruogu/gin_demo/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	DbHandler *gorm.DB
	cfg       *config.Database
}

func New(cfg *config.Database) *DB {
	return &DB{
		cfg: cfg,
	}
}

func (d *DB) Init() error {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.cfg.User, d.cfg.Password, d.cfg.Host, strconv.Itoa(d.cfg.Port), d.cfg.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	d.DbHandler = db
	return nil
}
func (d *DB) Close() {
	s, _ := d.DbHandler.DB()
	s.Close()
}
