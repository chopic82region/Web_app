package main

import (
	"log"

	"todo-app/database/db"
	"todo-app/internal"

	"github.com/spf13/viper"
)

func main() {
	handler := internal.NewHandler()
	server := internal.NewServer(*handler)

	viper.SetConfigName("cfg")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error read config:", err)
	}

	servPort := viper.GetString("ports")

	log.Println("Connecting to database...")
	postgres, err := db.Conect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection established")

	log.Println("Running migrations...")
	if err := db.Migrate(postgres); err != nil {
		log.Fatal("error of migration:", err)
	}
	log.Println("Migrations completed successfully")

	// Проверяем, что таблица создана
	var tableExists bool
	err = postgres.QueryRow("SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'tasks')").Scan(&tableExists)
	if err != nil {
		log.Printf("Warning: could not check if table exists: %v", err)
	} else if tableExists {
		log.Println("✓ Table 'tasks' exists in database")
	} else {
		log.Println("✗ Table 'tasks' NOT found in database!")
	}

	if err := server.Run(servPort); err != nil {
		log.Fatal("error of starting server")
		return
	}

}
