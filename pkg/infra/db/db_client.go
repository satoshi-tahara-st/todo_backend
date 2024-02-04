package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/xo/dburl"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var errNoSuchDialector = errors.New("no such dialector")

type DBConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	DBName   string

	MaxOpenConns           int
	MaxIdleConns           int
	ConnMaxLifetimeSeconds int
}

func NewDBConfig(driver string, user string, password string, host string, dbName string, maxOpenConn int, maxIdleConn int, lifeTimeSec int) DBConfig {
	return DBConfig{
		Driver:                 driver,
		Username:               user,
		Password:               password,
		Host:                   host,
		DBName:                 dbName,
		MaxOpenConns:           maxOpenConn,
		MaxIdleConns:           maxIdleConn,
		ConnMaxLifetimeSeconds: lifeTimeSec,
	}
}

func NewDB(conf DBConfig) (*gorm.DB, error) {
	dialector, err := findDialector(conf)
	if err != nil {
		return nil, err
	}

	gormDb, err := gorm.Open(dialector)
	if err != nil {
		return nil, fmt.Errorf("failed to open db:\n %w", err)
	}

	sqlDb, err := gormDb.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB:\n %w", err)
	}

	sqlDb.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDb.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDb.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifetimeSeconds) * time.Second)

	return gormDb, nil
}

// findDialector 接続先DBに合わせたgorm.Dialectorを検出します
func findDialector(conf DBConfig) (gorm.Dialector, error) {
	du, err := dburl.Parse(fmt.Sprintf("%s://%s:%s@%s/%s", conf.Driver, conf.Username, conf.Password, conf.Host, conf.DBName))
	if err != nil {
		return nil, fmt.Errorf("failed to parse db url:\n %w", err)
	}

	switch conf.Driver {
	case "mysql":
		return mysql.Open(du.DSN), nil
	}

	return nil, fmt.Errorf("invalid driver name (%s):\n %w", conf.Driver, errNoSuchDialector)
}
