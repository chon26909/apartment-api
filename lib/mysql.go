package lib

import (
	"context"
	"fmt"
	"time"

	"github.com/chon26909/apartment-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func NewMySqlConnection(config *config.Config) *gorm.DB {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Db.Username,
		config.Db.Password,
		config.Db.Host,
		config.Db.Port,
		config.Db.Database,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: false,
	})

	if err != nil {
		panic(err)
	}

	err = db.Migrator().AutoMigrate()

	if err != nil {
		fmt.Println("err auto migrate", err)
	}

	return db
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	stmt, _ := fc()
	fmt.Printf("\nSQL => %v\n", stmt)
}
