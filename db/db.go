package db

import (
	"database/sql"
	"fmt"
	"github.com/desferreira/alurago/config"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)


func CreateConnection() *sql.DB{
	config.ReadConfig()
	user := viper.GetString("database.user")
	dbname := viper.GetString("database.dbname")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	sslmode := viper.GetString("database.sslmode")


	sqlString := fmt.Sprintf("user=%v dbname=%v password=%v host=%v sslmode=%v", user,
		dbname, password, host, sslmode)

	db, err := sql.Open("postgres", sqlString)
	if err != nil {
		panic(err)
	}
	return db

}

func CloseConnection(db *sql.DB) {
	db.Close()
}