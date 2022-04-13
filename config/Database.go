package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

func Database(c Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database)
	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}

func ReadEnv() Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("ERROR LOAD FILE", err)
	}

	res := Config{}
	res.Host = os.Getenv("DB_HOST")
	res.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	res.Database = os.Getenv("DB_DATABASE")
	res.Username = os.Getenv("DB_USERNAME")
	res.Password = os.Getenv("DB_PASSWORD")

	return res
}