package main

import (
	"github.com/Stanlyzoolo/practiceGo/pkg/usecases"
	"github.com/Stanlyzoolo/practiceGo/pkg/datatypes"
	"github.com/Stanlyzoolo/practiceGo/pkg/storage"
	"os"
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
		Logger: log,
		Samples: []usecases.Fahrenheit{},
	}

	dataHistoryHandler := &datatypes.createQueueObj{
		Logger: log,
		Samples: []datatypes.stack[],
	}

	usecases.CheckAndSave(tempHistoryHandler, usecases.Celsius(32), 5)

	storage.ConnectionToPostgre()
	storage.ConnectToRedis()
}
