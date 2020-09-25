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

func toFahrenheit(t interface{}) {
	if 	c, ok := t.(Celsius); ok {
		temp:= c*((9/5) + 32)
		fmt.Printf("Работаем с температурой в Цельсиях", Fahrenheit(temp))
	}

	if i, ok := t.(int); ok {
		temp:= i*(9/5) + 32
		fmt.Printf("Работаем с температурой в интах", Fahrenheit(temp))
	} 
	if k,ok := t.(string); ok {
		conv, _ := strconv.Atoi(k)
		temp:= conv*(9/5)+32
		fmt.Printf("Работаем со строкой", Fahrenheit(temp))
	}
}
