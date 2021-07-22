/*
Напишите программу, которая принимает на вход фразу и составляет аббревиатуру по первым буквам слов:
https://stepik.org/lesson/526868/step/8?unit=519587
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var abbr string

	phrase := readString()
	words := strings.Fields(phrase)

	for _, word := range words {
		firstLetter := []rune(word)[0]

		if unicode.IsLetter(firstLetter) {
			abbr += string(unicode.ToUpper(firstLetter))
		}
	}

	fmt.Println(abbr)
}

func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
