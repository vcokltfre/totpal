package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/vcokltfre/totpal/src/api"
)

func init() {
	godotenv.Load()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.Info("Initialisation complete")
}

func main() {
	if err := api.Start(os.Getenv("API_BIND")); err != nil {
		logrus.Fatal(err)
	}
}
