package usecases

import (
	"github.com/sirupsen/logrus"
)

// type Celsius float32
// type Fahrenheit float32
type HistoryHandler struct{
	Samples []Temper
	Logger	*logrus.Logger
}

func CheckAndSave(h *HistoryHandler, t Celsius, n int) {
	for i:= 0; i<n; i++ {
		f:= toFahrenheit(t) + Fahrenheit(i)
		h.Samples = append(h.Samples, f)
		h.Logger.Info("Предыдущие изменения и текущее значение - ", h.Samples, f)
	}
}

type Temper struct {
	Celsius float32
	Fahrenheit float32
}

func toFahrenheit(ter *Temper) float32 {
	var temp Temper
	return ((temp.Fahrenheit*(9/5) + 32))	
}
