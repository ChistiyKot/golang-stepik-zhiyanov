/*
Напишите программу, которая определяет название языка по его коду. Правила:
https://stepik.org/lesson/526867/step/11?unit=519586
*/
package main

import (
	"fmt"
)

func main() {
	var code, lang string
	fmt.Scan(&code)

	if code == "en" {
		lang = "English"
	} else if code == "fr" {
		lang = "French"
	} else if code == "ru" || code == "rus" {
		lang = "Russian"
	} else {
		lang = "Unknown"
	}

	fmt.Println(lang)
}
