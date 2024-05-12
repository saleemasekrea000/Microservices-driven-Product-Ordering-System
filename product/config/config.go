package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configration struct {
	DbUrl string
	Port  int
}

func New() *Configration {
	godotenv.Load("../.env")
	port := os.Getenv("port")
	portInt, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	host := os.Getenv("DB_HOST")
	dbport := os.Getenv("DB_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	println(host, dbport, user, password, dbname)
	if dbname == "" || user == "" || password == "" || host == "" || dbport == "" {
		panic("DB_HOST, DB_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB env not set")
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, dbport, user, password, dbname)

	return &Configration{
		DbUrl: connStr,
		Port:  portInt,
	}
}
