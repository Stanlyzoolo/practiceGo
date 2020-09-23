package usecases

import (
	"github.com/sirupsen/logrus"
)

type Celsius float32
type Fahrenheit float32
type HistoryHandler struct{
	Samples []Fahrenheit
	Logger	*logrus.Logger
}

func CheckAndSave(h *HistoryHandler, t Celsius, n int) {
	for i:= 0; i<n; i++ {
		f:= toFahrenheit(t) + Fahrenheit(i)
		h.Samples = append(h.Samples, f)
		h.Logger.Info("Предыдущие изменения и текущее значение - ", h.Samples, f)
	}
}

func toFahrenheit(t Celsius) Fahrenheit {
	var temp Fahrenheit	
	temp = Fahrenheit((t*9/5) + 32)

	return temp
}