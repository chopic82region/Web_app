package db

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func Conect() (*sql.DB, error) {

	viper.SetConfigName("cfg")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error read config:", err)
	}

	port := viper.GetString("db.port")
	// Если порт в формате Docker "5432:5432", берем только первую часть
	if strings.Contains(port, ":") {
		port = strings.Split(port, ":")[0]
	}
	host := viper.GetString("db.host")
	if host == "" {
		host = "localhost"
	}
	dbname := viper.GetString("db.dbname")
	password := viper.GetString("db.password")
	sslmode := viper.GetString("db.sslmode")
	username := viper.GetString("db.username")

	connectToString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", host, port, username, dbname, password, sslmode)
	
	// Логируем строку подключения без пароля для отладки
	log.Printf("Connecting to database: host=%s port=%s user=%s dbname=%s sslmode=%s", host, port, username, dbname, sslmode)

	db, err := sql.Open("postgres", connectToString)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return db, nil
}
