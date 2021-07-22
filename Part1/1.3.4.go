/*
Напишите программу, которая укорачивает строку до указанной длины и добавляет в конце многоточие:
https://stepik.org/lesson/526868/step/4?unit=519587
*/
package main

import (
	"fmt"
)

func main() {
	var text, res string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	if width < len(text) {
		res = text[:width] + "..."
	} else {
		res = text
	}

	fmt.Println(res)
}
