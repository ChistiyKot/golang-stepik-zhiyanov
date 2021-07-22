/*
Напишите программу, которая считает количество минут во временном отрезке.
https://stepik.org/lesson/526867/step/3?unit=519586
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)

	fmt.Println(s, "=", d.Minutes(), "min")
}
