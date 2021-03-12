package persistence

import (
	"database/sql"
	"fmt"
	"jeeo/api-boilerplate/configs"
	"log"

	_ "github.com/lib/pq"
)

//Database is the database type
type Database struct {
	Conn *sql.DB
}

// Init intializes a database connection based on provided config
func (DB *Database) Init(config configs.Config) {
	var err error
	connString := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Db.Host, config.Db.Port, config.Db.User, config.Db.Pwd, config.Db.Name)
	DB.Conn, err = sql.Open("postgres", connString)
	if err != nil {
		log.Println("Error on connect to Database: ", err.Error())
	}
}

// ProvideDatabase provides a database type instance
func ProvideDatabase(config configs.Config) *Database {
	db := &Database{}
	db.Init(config)

	return db
}
