package config

import (
	"os"
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	wd, _ := os.Getwd()
	err := godotenv.Load(".env")
	fmt.Println(wd)
	if err != nil {
		panic("failed to get .env" + err.Error())
	}
}