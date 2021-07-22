/*
Программа принимает на вход строку source и число times. Требуется склеить source саму с собой times раз и вернуть результат:
https://stepik.org/lesson/526867/step/8?unit=519586
*/
package main

import (
	"fmt"
)

func main() {
	var source, result string
	var times int
	fmt.Scan(&source, &times)

	for x := 0; x < times; x++ {
		result += source
	}

	fmt.Println(result)
}
