package main

import (
	"fmt"
	"math"
)

func diameterCircle(s float64) float64 {
	var d float64
	d = 2 * math.Sqrt(s/math.Pi)
	return math.Round(d*100) / 100
}

func circumference(s float64) float64 {
	var p float64
	p = 2 * math.Sqrt(s/math.Pi) * math.Pi
	return math.Round(p*100) / 100
}

func main() {
	var s float64

	fmt.Print("Введите площадь круга S = ")
	_, err := fmt.Scanln(&s)
	if s < 0 || err != nil {
		fmt.Println("Ошибка 'S' - отрицательное число или строка")
	} else {
		fmt.Println("Диаметр окружности D =", diameterCircle(s))
		fmt.Println("Длина окружности P =", circumference(s))
	}
}
