package constants

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	LoginUrl        string
	PutStopLimitUrl string
	Ask             int
	Bid             int
	MinPrice        int
	MaxPrice        int
	MinAmount       int
	MaxAmount       int

	WalletId  int
	Method    string
	AssetPair string

	RedisAddr     string
	RedisDB       int
	RedisPassword string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	LoginUrl = os.Getenv("LOGIN_URL")
	PutStopLimitUrl = os.Getenv("PUT_STOP_LIMIT_URL")

	Ask, _ = strconv.Atoi(os.Getenv("ASK"))
	Bid, _ = strconv.Atoi(os.Getenv("BID"))
	MinPrice, _ = strconv.Atoi(os.Getenv("MIN_PRICE"))
	MaxPrice, _ = strconv.Atoi(os.Getenv("MAX_PRICE"))
	MinAmount, _ = strconv.Atoi(os.Getenv("MIN_AMOUNT"))
	MaxAmount, _ = strconv.Atoi(os.Getenv("MAX_AMOUNT"))

	WalletId, _ = strconv.Atoi(os.Getenv("WALLET_ID"))
	Method = os.Getenv("METHOD")
	AssetPair = os.Getenv("ASSET_PAIR")

	RedisAddr = fmt.Sprintf(`%s:%s`, os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	RedisDB, _ = strconv.Atoi(os.Getenv("REDIS_DB"))
	RedisPassword = os.Getenv("REDIS_PASSWORD")
}
