package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Pega as variáveis de ambiente para a conexão
	host := viper.GetString("DB_HOST")
	port := viper.GetString("DB_PORT")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbname := viper.GetString("DB_NAME")

	// Tenta se conectar ao banco de dados principal
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database '%s', attempting to create it...", dbname)

		// Se a conexão falhar, tenta conectar ao banco de dados padrão 'postgres'
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
			host, port, user, password)

		var tempDB *gorm.DB
		tempDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect to the postgres server: %v", err)
		}

		// Cria o banco de dados principal
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", dbname)
		if err := tempDB.Exec(createDBQuery).Error; err != nil {
			log.Fatalf("Failed to create database: %v", err)
		}

		log.Println("Database created successfully. Reconnecting to the new database...")

		// Fecha a conexão temporária
		sqlDB, _ := tempDB.DB()
		sqlDB.Close()

		// Tenta se conectar novamente ao banco de dados principal
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		// Retenta a conexão por um tempo para garantir que o BD esteja pronto
		for i := 0; i < 5; i++ {
			DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				log.Println("Successfully reconnected to the new database.")
				break
			}
			log.Printf("Reconnection attempt %d failed: %v. Retrying in 2 seconds...", i+1, err)
			time.Sleep(2 * time.Second)
		}

		if err != nil {
			log.Fatalf("Failed to reconnect to the database after creation: %v", err)
		}

	}
	log.Println("Database connection established.")
}

func AutoMigrate(models ...interface{}) {
	if DB == nil {
		log.Fatalf("Database connection is nil")
	}
	log.Println("Running GORM AutoMigrate...")
	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("Migrations completed successfully.")
}

func IsDatabaseNotExistError(err error) bool {
	return os.IsExist(err) || err.Error() == "FATAL: database does not exist"
}
