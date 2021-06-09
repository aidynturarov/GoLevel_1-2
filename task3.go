package main

import "fmt"

func main() {
	var a uint16

	fmt.Println("Для вычисление разряда введите трехзначное число")
	if _, err := fmt.Scanln(&a); err != nil || a >= 1000 || a < 100 {
		fmt.Println("Ошибка, число не трехзначное или строка")
	} else {
		var i uint8
		for a > 0 {
			i++
			switch i {
			case 1:
				fmt.Println("Единицы", a%10)
			case 2:
				fmt.Println("Десятки", a%10)
			case 3:
				fmt.Println("Сотни", a%10)
			default:
				break
			}
			a /= 10
		}
	}
}
