package config

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbConnectionLogger *logrus.Logger

func SetDBConnectionLogger(l *logrus.Logger) {
	dbConnectionLogger = l
}

var DB *gorm.DB

func ConnectDB(config *Config) (*gorm.DB, error) {
	var err error

	//creating database connections
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.HOST, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.SchemaName + ".", // schema name
			SingularTable: true,
		},
	})

	if err != nil {
		if dbConnectionLogger != nil {
			dbConnectionLogger.WithFields(logrus.Fields{}).Error("Failed to initialize database")
		}
		return nil, errors.New("database connection error")
	}
	fmt.Println("Table prefix set to", config.SchemaName)
	return db, nil
}

func PooledConnectDB(config *Config) error {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		config.HOST, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.SchemaName + ".", // schema name
			SingularTable: true,
		},
	})

	if err != nil {
		if dbConnectionLogger != nil {
			dbConnectionLogger.WithFields(logrus.Fields{}).Error("Failed to initialize database")
		}
		return errors.New("database connection error")
	}

	pqlDB, err := db.DB()
	if err != nil {
		return err
	}

	pqlDB.SetMaxIdleConns(config.MaxIdleConns)
	pqlDB.SetMaxOpenConns(config.MaxOpenConns)
	pqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	DB = db

	fmt.Printf("Connected to database: %s\n", config.DBName)
	return nil
}
