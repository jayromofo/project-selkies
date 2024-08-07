package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInfo struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes() { return }

func InitializeDatabase() *DBInfo {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("HOST")
	user := os.Getenv("DBUSER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	port := "5432"

	return &DBInfo{
		host,
		user,
		password,
		dbname,
		port,
	}
}

func (r *Repository) Connect(info DBInfo) error {

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		info.Host, info.User, info.Password, info.DBName, info.Port,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("failed to connect to database")
	}

	fmt.Println("You connected to ", info.DBName)

	r.DB = db
	return nil

}
