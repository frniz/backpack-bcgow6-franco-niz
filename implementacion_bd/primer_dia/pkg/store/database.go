package store

import (
	"database/sql"
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDatabase() (engine *gin.Engine, db *sql.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		return nil, nil, errors.New("error: Loading .env")
	}

	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	db, err = sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		return nil, nil, errors.New("error configurando la DB")
	}

	engine = gin.Default()

	return engine, db, nil
}
