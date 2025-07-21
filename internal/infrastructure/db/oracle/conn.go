package oracle

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/GustavoZeglan/Cine/internal/config"
	oracleGorm "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func Connect() *gorm.DB {
	once.Do(func() {
		dbConn := config.AppConfig.OracleDB
		dsn := fmt.Sprintf("oracle://%s:%s@%s:%v/%s", dbConn.User, dbConn.Password, dbConn.Host, dbConn.Port, dbConn.Name)
		db, err := gorm.Open(oracleGorm.New(oracleGorm.Config{
			DSN: dsn,
		}))
		if err != nil {
			log.Fatal("Error:", err)
		}

		if err != nil {
			log.Fatal("Failed to connect to database:", err)
			os.Exit(1)
		}

		DB = db
	})

	return DB
}

func GetConnection() *gorm.DB {
	return DB
}
