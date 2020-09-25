package usecases

import (
	"strconv"
	"fmt"
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

func toFahrenheit(t interface{}) {
	// var temp Fahrenheit
	// temp = Fahrenheit(t*(9/5) + 32)

	c, ok := t.(Celsius)

	if ok {
		temp:= Celsius(t*((9/5) + 32)
		fmt.Printf("Работаем с температурой в Цельсиях", temp)
	}

	if i, ok := t.(int); ok {
		temp:= t*(9/5) + 32
		fmt.Printf("Работаем с температурой в интах", temp)
	} 
	if k,ok := t.(string); ok {
		conv, _ = strconv.Atoi(k)
		temp:= conv*(9/5)+32
		fmt.Printf("Работаем со строкой", temp)
	}

	return temp
}
