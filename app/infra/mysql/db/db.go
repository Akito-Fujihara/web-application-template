package db

import (
	"fmt"
	"log/slog"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type MysqlConfig struct {
	Database	 string
	Host		 string
	ReadonlyHost string
	Port		 int
	Username	 string
	ReadonlyUsername string
	Password	 string
	ReadonlyPassword string
}

func NewMysqlConn(config *MysqlConfig) (*gorm.DB, func(), error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)
	readDns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		config.ReadonlyUsername,
		config.ReadonlyPassword,
		config.ReadonlyHost,
		config.Port,
		config.Database,
	) 
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	if err := db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{mysql.Open(dns)},
		Replicas: []gorm.Dialector{mysql.Open(readDns)},
		Policy: dbresolver.RandomPolicy{},
	})) ; err != nil {
		return nil, nil, fmt.Errorf("failed to register dbresolver: %v", err)
	}

	// コネクションプールの設定をしても良い

	cleanup := func() {
		sqlDB, _ := db.DB() // err は close error として無視
		if err := sqlDB.Close(); err != nil {
			slog.Error(fmt.Sprintf("db connection close error: %v", err))
			return
		}
		slog.Info("db connection closed")
	}

	return db, cleanup, nil
}
