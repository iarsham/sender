package app

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func GetEnvs() {
	if err := godotenv.Load(); err != nil {
		panic("error in loading env variables")
	}
}

func GetLogger() *logrus.Logger {
	Log = logrus.New()
	Log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return Log
}

func FileManager() {
	if _, err := os.Open(os.Getenv("FILE")); err != nil {
		_, err := os.Create(os.Getenv("FILE"))
		if err != nil {
			Log.Fatal(err)
		}
	}
}
