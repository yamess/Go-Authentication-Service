package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var Host string
var PgDbUrl string
var Version string
var DefaultUser uint = 1
var Algorithm string
var SecretKey string
var HashingCost int

const ExpiredTime = 10

func InitEnv() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Error while loading the environment file")
	}
	Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	PgDbUrl = os.Getenv("DB_URL")
	Version = os.Getenv("VERSION")
	Algorithm = os.Getenv("ALGORITHM")
	SecretKey = os.Getenv("SECRET_KEY")
	HashingCost, _ = strconv.Atoi(os.Getenv("HASH_COST"))
}
