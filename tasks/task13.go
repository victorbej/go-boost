package tasks

/*
 Задача 13.	Поменять местами два числа без создания временной переменной.
*/

import "fmt"

func TaskThirteen() {
	x := 1
	y := 2
	fmt.Printf("x=%d, y=%d\n", x, y)

	x, y = y, x
	fmt.Printf("x=%d, y=%d\n", x, y)
}
