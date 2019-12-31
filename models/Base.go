package models

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

var db *gorm.DB

func init() {
	e := godotenv.Load() //* load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, username, dbName, password) //* build db connection string

	var err error
	db, err = gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db.Debug().AutoMigrate(&Account{}, &SavingsPlan{}, &Savings{}) //* db migration
}

//* Base contains common columns for all tables
type Base struct {
	ID        string     `gorm:"type:uuid;primary_key;column:id"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

//* BeforeCreate will set up a UUID rather than numeric ID
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4().String())
	return nil
}

//* return a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
