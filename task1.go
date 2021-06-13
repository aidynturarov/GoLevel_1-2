package main

import "fmt"

func rectangleArea(x, y float64) float64 {
	return x * y
}

func main() {
	var a, b float64

	fmt.Println("Для вычисления площади прямоугольника")
	fmt.Print("Введите длину стороны a = ")
	_, errA := fmt.Scanln(&a)
	if a < 0 || errA != nil {
		fmt.Println("Ошибка 'a' - отрицательное число или строка")
	} else {
		fmt.Print("Введите длину стороны b = ")
		_, errB := fmt.Scanln(&b)
		if b < 0 || errB != nil {
			fmt.Println("Ошибка 'b' - отрицательное число или строка")
		} else {
			fmt.Println("Площадь прямоугольника S =", rectangleArea(a, b))
		}
	}
}
