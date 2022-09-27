package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	SERVER_PORT int
	DB_DRIVER   string
	DB_HOST     string
	DB_USERNAME string
	DB_PORT     int
	DB_PASSWORD string
	DB_NAME     string
}

var lock = &sync.Mutex{}
var config *AppConfig

func GetConfig() *AppConfig {

	lock.Lock()
	defer lock.Unlock()

	if config == nil {
		config = initConfig()
	}

	return config

}

func initConfig() *AppConfig {

	var defaultConfig AppConfig

	serverPortConv, errConv1 := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if errConv1 != nil {
		log.Fatal("error parse DB PORT")
		return nil
	}
	defaultConfig.SERVER_PORT = serverPortConv
	defaultConfig.DB_DRIVER = os.Getenv("DB_DRIVER")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	port, errConv := strconv.Atoi(os.Getenv("DB_PORT"))
	if errConv != nil {
		log.Fatal("error parse DB port")
		return nil
	}
	defaultConfig.DB_PORT = port

	return &defaultConfig

}
