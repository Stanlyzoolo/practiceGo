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

func CheckAndSave(h *HistoryHandler, t interface{}, n int) {
	for i:= 0; i<n; i++ {
		f:= toFahrenheit(t) + Fahrenheit(i)
		h.Samples = append(h.Samples, f)
		h.Logger.Info("Предыдущие изменения и текущее значение - ", h.Samples, f)
	}
}

func toFahrenheit(t interface{}) Fahrenheit {
	var temp Fahrenheit
	if 	c, ok := t.(Celsius); ok {
		temp := Fahrenheit(c*((9/5) + 32))
		fmt.Printf("Работаем с температурой в Цельсиях", temp)
		return temp
	}

	if c, ok := t.(int); ok {
		temp := Fahrenheit(c*(9/5) + 32)
		fmt.Printf("Работаем с температурой в интах", temp)
		return temp
	} 

	if c, ok := t.(string); ok {
		conv, _ := strconv.Atoi(c)
		temp := Fahrenheit(conv*(9/5)+32)
		fmt.Printf("Работаем со строкой", temp)
		return temp
	}
	return temp
}
