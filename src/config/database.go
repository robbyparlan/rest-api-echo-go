package config

import (
	"fmt"
	"time"
	"log"
	"os"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var DB *sqlx.DB

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	argsLen := len(os.Args)
	if argsLen > 1 {
		if os.Args[1] == "pg" {
			viper.SetConfigName("config-pg") // name of config file (without extension)
		}
	}
	viper.AddConfigPath(".") // optionally look for config in the working directory

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(fmt.Errorf("Fatal error reading config file json: %s", err))
	}

	connString := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=%s",
		viper.Get("database.psql.host"),
		viper.Get("database.psql.username"),
		viper.Get("database.psql.password"),
		viper.GetInt("database.psql.port"),
		viper.Get("database.psql.database"),
		viper.Get("database.psql.ssl_mode"),
	)

	DB, err = sqlx.Open("postgres", connString)
	if err != nil {
		log.Println("Error creating connection pool: " + err.Error())
	}
	DB.SetMaxOpenConns(viper.GetInt("database.psql.max_open_connection")) // Sane default
	DB.SetMaxIdleConns(viper.GetInt("database.psql.max_idle_connection"))

	DB.SetConnMaxLifetime(time.Minute * 10)

	
	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
	
	log.Printf("Database Connected!\n")
}