package main

import (
	"database/sql"
	"os"

	"github.com/Stanlyzoolo/practiceGo/pkg/datatypes"
	"github.com/Stanlyzoolo/practiceGo/pkg/storage"
	"github.com/Stanlyzoolo/practiceGo/pkg/usecases"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type EnvConfig struct {
	LogFile string `envconfig:"LOG_FILE"`
}

func main() {
	var eConf EnvConfig
	envconfig.Process("", &eConf)

	file, _ := os.OpenFile(eConf.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	var log = logrus.New()
	log.Out = file

	tempHistoryHandler := &usecases.HistoryHandler{
		Logger:  log,
		Samples: []usecases.Fahrenheit{},
	}

	usecases.CheckAndSave(tempHistoryHandler, usecases.Celsius(32), 5)

	datatypes.CreateQueueObj(tempHistoryHandler)

	word := "Lenovo"

	datatypes.ReverseWordOne(word, tempHistoryHandler)
	datatypes.ReverseWordTwo(word, tempHistoryHandler)

	sent := "Lenovo is good"
	datatypes.ReverseWordsinSentence(sent, tempHistoryHandler)

	connStr := "user=postgres password=postgres dbname=freeit sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	storage.WrightInDatabase(db, "Stanislav", "Lepeshko", 24)
	storage.ReadFromDatabase(db)
}
