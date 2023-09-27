package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type ProgramConfig struct {
	DB_Username   string
	DB_Password   string
	DB_Port       string
	DB_Host       string
	DB_Name       string
	SECRET        string
	SECRETREFRESH string
}

func InitConfig() *ProgramConfig {
	var res = new(ProgramConfig)
	res = loadConfig()

	if res == nil {
		logrus.Fatal("Config : Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *ProgramConfig {
	var res = new(ProgramConfig)

	err := godotenv.Load(".env")

	if err != nil {
		logrus.Error("Config : Cannot load config file,", err.Error())
		return nil
	}

	if val, found := os.LookupEnv("DB_USERNAME"); found {
		res.DB_Username = val
	}

	if val, found := os.LookupEnv("DB_PASSWORD"); found {
		res.DB_Password = val
	}
	if val, found := os.LookupEnv("DB_PORT"); found {
		res.DB_Port = string(val)
	}

	if val, found := os.LookupEnv("DB_HOST"); found {
		res.DB_Host = val
	}

	if val, found := os.LookupEnv("DB_NAME"); found {
		res.DB_Name = val
	}

	if val, found := os.LookupEnv("SECRET"); found {
		res.SECRET = val
	}

	if val, found := os.LookupEnv("SECRETREFRESH"); found {
		res.SECRETREFRESH = val
	}

	return res

}
