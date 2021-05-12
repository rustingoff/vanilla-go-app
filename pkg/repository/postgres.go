package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetDataBase() *gorm.DB {
	DB, err := NewPostgresDB(Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "crud_system",
		DBName:   "crud_system",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("Can't connect to database: %s", err.Error())
		return nil
	}

	return DB
}
