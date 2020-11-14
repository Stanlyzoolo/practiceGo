package main

import (
	"os"
	"time"

	"github.com/Stanlyzoolo/practiceGo/pkg/channels"
	"github.com/Stanlyzoolo/practiceGo/pkg/datatypes"
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

	str := "Work or not, tell me please, master?"
	tempHistoryHandler.Info("перед чтением", str)

	ch := make(chan string)
	go channels.ReadTheChannel(ch, tempHistoryHandler)
	time.Sleep(5 * time.Second)
	ch <- str

	tempHistoryHandler.Info("после чтения", str)
}
